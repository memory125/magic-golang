package splitstring

import (
	"strings"
)

// 字符串切分
/*
  str: 原字符串
  sep: 切割的字符或字符串
*/
func SplitStr(str string, sep string) []string {
	var ret []string
	// 获得sep所在的index值
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index])
		// fmt.Println("ret: ", ret)
		str = str[index+len(sep):]
		// fmt.Println("str: ", str)
		index = strings.Index(str, sep)
		// fmt.Println("index: ", index)
	}
	ret = append(ret, str)
	return ret
}
