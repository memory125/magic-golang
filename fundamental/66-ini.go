package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// ini文件解析

// MySQL配置
type MySQLConfig struct {
	Address string   `ini:"address"`
	Port int         `ini:"port"`
	Username string  `ini:"username"`
	Password string  `ini:"password"`
}

// Redis配置
type RedisConfig struct {
	Host string   `ini:"host"`
	Port int         `ini:"port"`
	Password string  `ini:"password"`
	Database string  `ini:"database"`
}

type Config struct {
	MySQLConfig  `ini:"mysql"`
	RedisConfig  `ini:"redis"`
}


func loadIniConfig(fileName string, data interface{}) error {
	// 0. 参数的校验
	// 0.1 传递的data参数必须是指针类型(因为需要在函数中对其赋值)
	t := reflect.TypeOf(data)
	fmt.Println(t.Kind())
	if t.Kind() != reflect.Ptr {
		err := errors.New("data parameter should be a pointer")    // 新创建一个错误
		return err
	}
	// 0.2 传递的data参数必须时结构体类型(因为配置文件中有键值对需要赋值给结构体的字段)
	if t.Elem().Kind() != reflect.Struct {
		err := errors.New("data parameter should be a struct pointer")    // 新创建一个错误
		return err
	}
	// 1. 读文件得到字节类型的数据
	bytesRead, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("read file failed, error is %v.\n", err)
		return err
	}
	linesSlice := strings.Split(string(bytesRead), "\r\n")
	var structName string
	// 2. 按行读取数据
	for index, line := range linesSlice {
		// 去掉首尾的空格
		line = strings.TrimSpace(line)
		// 2.1 如果是注释跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#")  {
			continue
		}
		// 2.2 如果时”[“开头说明时unit
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line) - 1] != ']'{
				err = fmt.Errorf("line:%d syntax error", index + 1)
				return err
			}
			// 把这一行内首尾[]去掉，去掉首尾的空格，取到中间的内容
			sectionName := strings.TrimSpace(line[1:len(line) - 1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", index + 1)
				return err
			}
			// 根据字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("Find----section name is %v, struct name is %v.\n", sectionName, structName)
				}
			}
		} else {
			// 2.3 如果不是”[“开头，则说明时”=“分割的键值对
			


		}
	}
	return nil
}

func main()  {
	var conf Config
	err := loadIniConfig("D:\\project\\go\\gopath\\src\\code.wing.com\\fundamental\\conf.ini", &conf)
	if err != nil {
		fmt.Printf("load ini config file failed, error is %v.\n", err)
		return
	}
	fmt.Printf("MySQL config are: address = %s, port = %d, username = %s, password = %s.\n", conf.MySQLConfig.Address, conf.MySQLConfig.Port, conf.MySQLConfig.Username, conf.MySQLConfig.Password)

}
