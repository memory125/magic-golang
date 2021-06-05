package main

import (
	"bufio"            //
	"fmt"
	"os"
)

func main()  {
	fmt.Println("请输入一个字符串：")
	// 从标准输入读取，如键盘输入
	// 创建bufio的Reader对象
	reader := bufio.NewReader(os.Stdin)
	// 获取字符串信息，直到换行符
	s1, _ := reader.ReadString('\n')
	fmt.Println("读到的数据：", s1)
}
