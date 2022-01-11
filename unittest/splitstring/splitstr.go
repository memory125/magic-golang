package splitstring

import (
	"fmt"
	"strings"
)

// 字符串切分

func SplitStr(str string, sep string) []string {
	var ret []string
	// 获得sep所在的index值
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index])
		fmt.Println("ret: ", ret)
		str = str[index+1:]
		fmt.Println("str: ", str)
		index = strings.Index(str, sep)
		fmt.Println("index: ", index)
	}
	ret = append(ret, str)
	return ret
}
