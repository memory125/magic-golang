package main

import (
	"fmt"
	"strconv"
)

// strconv示例

func main()  {
	// 字符串数字转数值类型
	// 方式一
	str1 := "888"
	intValue1, _ := strconv.Atoi(str1)
	/*
		Atoi Valus is 888, Type is int
	*/
	fmt.Printf("Atoi Valus is %v, Type is %T \n", intValue1, intValue1)

	// 方式二
	intValue2, err := strconv.ParseInt(str1, 10, 64)
	if err != nil {
		fmt.Printf("strconv ParseInt faliled, error is %v.\n", err)
		return
	}
	/*
		ParseInt Value is 888, Type is int64
	*/
	fmt.Printf("ParseInt Value is %v, Type is %T \n", intValue2, intValue2)

	// 字符串数字转浮点类型
	str2 := "3.2545114"
	floatValue, err := strconv.ParseFloat(str2, 32)
	if err != nil {
		fmt.Printf("strconv ParseFloat faliled, error is %v.\n", err)
		return
	}
	/*
		ParseFloat Value is 3.2545113563537598, Type is float64
	*/
	fmt.Printf("ParseFloat Value is %v, Type is %T \n", floatValue, floatValue)

	// 字符串bool值转bool类型
	str3 := "true"
	boolValue, err := strconv.ParseBool(str3)
	if err != nil {
		fmt.Printf("strconv ParseBool faliled, error is %v.\n", err)
		return
	}
	/*
		ParseBool Value is true, Type is bool
	*/
	fmt.Printf("ParseBool Value is %v, Type is %T \n", boolValue, boolValue)

	// 数值转字符串
	i := 666888
	strValue := strconv.Itoa(i)
	/*
		Itoa Value is 666888, Type is string
	*/
	fmt.Printf("Itoa Value is %v, Type is %T \n", strValue, strValue)

	// 浮点数转字符串类型
	f := 666.888
	floatStr := fmt.Sprintf("%f", f)
	/*
		Sprintf Value is 666.888000, Type is string
	*/
	fmt.Printf("Sprintf Value is %s, Type is %T \n", floatStr, floatStr)
}
