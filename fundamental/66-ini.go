package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// ini文件解析

// MySQL配置结构体
type MySQLConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// Redis配置结构体
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

// Config配置结构体
type Config struct {
	MySQLConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIniConfig(fileName string, data interface{}) error {
	// 0. 参数的校验
	// 0.1 传递的data参数必须是指针类型(因为需要在函数中对其赋值)
	t := reflect.TypeOf(data)
	//fmt.Println(t.Kind())
	if t.Kind() != reflect.Ptr {
		err := errors.New("data parameter should be a pointer") // 新创建一个错误
		return err
	}
	// 0.2 传递的data参数必须时结构体类型(因为配置文件中有键值对需要赋值给结构体的字段)
	if t.Elem().Kind() != reflect.Struct {
		err := errors.New("data parameter should be a struct pointer") // 新创建一个错误
		return err
	}
	// 1. 读文件得到字节类型的数据
	bytesRead, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("read file failed, error is %v.\n", err)
		return err
	}
	linesSlice := strings.Split(string(bytesRead), "\n")
	fmt.Println("Debug=====1======>", linesSlice)
	var structName string
	// 2. 按行读取数据
	for index, line := range linesSlice {
		// 去掉首尾的空格
		line = strings.TrimSpace(line)
		// 如果是空行，则跳过
		if len(line) == 0 {
			continue
		}

		// 2.1 如果是注释跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2 如果时”[“开头说明时unit
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", index+1)
				return err
			}
			// 把这一行内首尾[]去掉，去掉首尾的空格，取到中间的内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", index+1)
				return err
			}
			// 根据字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("Find=====>section name is %v, struct name is %v.\n", sectionName, structName)
				}
			}
		} else {
			// 2.3 如果不是”[“开头，则说明时”=“分割的键值对
			// 1. 以等号分割这一行，等号左边是key，右边是value
			idx := strings.Index(line, "=")
			if idx == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line: %d syntax error", index+1)
				return err
			}

			// 获取key和value的内容
			iniKey := strings.TrimSpace(line[:idx])
			iniValue := strings.TrimSpace(line[idx+1:])
			fmt.Println("Debug====2========>", idx, iniKey, iniValue)

			// 2. 根据structName去data里面把对应的嵌套结构体取出来
			v := reflect.ValueOf(data)

			fmt.Println("Debug====3========>", v)

			// 获取结构体的名称和类型
			structValue := v.Elem().FieldByName(structName) // 嵌套结构体的值信息
			structType := structValue.Type()                // 嵌套结构体的类型信息
			fmt.Println("Debug====3.1========>", structValue, structType)

			if structType.Kind() != reflect.Struct {
				err = fmt.Errorf("the member of %s from data should be the struct type", structName)
				return err
			}
			var fieldName string
			var fieldType reflect.StructField
			// 3. 遍历嵌套结构体的每一个字段，判断tag是不是等于key
			fmt.Println("Debug====3.2========>", structType.NumField())
			for i := 0; i < structType.NumField(); i++ {
				fieldTemp := structType.Field(i) // tag信息存储在类型信息中
				fieldType = fieldTemp
				fmt.Println("Debug====3.3========>", fieldTemp, fieldTemp.Tag.Get("ini"), fieldTemp.Name, iniKey)
				if fieldTemp.Tag.Get("ini") == iniKey {
					// 找到对应的字段
					fieldName = fieldTemp.Name
					break
				}
			}
			// 4. 如果key == tag，给这个字段赋值
			// 4.1 根据fieldName取出相应的字段
			// 如果在结构体中找不到对应的字段，则继续遍历
			if len(fieldName) == 0 {
				continue
			}
			obj := structValue.FieldByName(fieldName)
			tmpType := fieldType.Type.Kind()
			// 4.2 对其赋值
			fmt.Println("Debug====4==========>", obj, fieldName, tmpType)

			switch tmpType {
			case reflect.String:
				obj.SetString(iniValue[:])

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(iniValue, 10, 64)
				if err != nil {
					// 如果解析出错
					err = fmt.Errorf("line: %d value type error", index+1)
					return err
				}
				obj.SetInt(valueInt)

			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(iniValue)
				if err != nil {
					// 如果解析出错
					err = fmt.Errorf("line: %d value type error", index+1)
					return err
				}
				obj.SetBool(valueBool)

			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(iniValue, 64)
				if err != nil {
					// 如果解析出错
					err = fmt.Errorf("line: %d value type error", index+1)
					return err
				}
				obj.SetFloat(valueFloat)
			}

		}
	}
	return nil
}

func main() {
	var conf Config
	err := loadIniConfig("D:\\project\\go\\gopath\\src\\code.wing.com\\fundamental\\conf.ini", &conf)
	if err != nil {
		fmt.Printf("load ini config file failed, error is %v.\n", err)
		return
	}
	//fmt.Printf("MySQL config are: address = %s, port = %d, username = %s, password = %s.\n", conf.MySQLConfig.Address, conf.MySQLConfig.Port, conf.MySQLConfig.Username, conf.MySQLConfig.Password)
	fmt.Printf("%#v\n", conf)
}
