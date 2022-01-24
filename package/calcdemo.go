package main

/*
	包 - package
    - 包的路径有`GOPATH/src`后面的路径开始写起，路径分隔符用`/`
    - 想被别的包调用的标识符都要首字母大写
    - 单行导入和多行导入
    - 导入包的时候可以指定别名
    - 导入包不想使用包内部的标识符，需要使用匿名导入
    - 每个包导入的时候会自动执行一个名为`init()`的函数，它没有参数也没有返回值，不能手动调用
    - `init()`的执行顺序：变量声明 > init() > main()
*/

import (
	calc "code.wing.com/package/calc"
	"fmt"
)

// 全局变量
var x int = 20

const pi = 3.1415926

func init() {
	fmt.Printf("x = %v, pi = %v.\n", x, pi)
}

func main() {
	ret := calc.Add(50, 80)
	/*
		-------------------------Module cals---------------------
		x = 20, pi = 3.1415926.
		ret:  130
	*/
	fmt.Println("ret: ", ret)
}
