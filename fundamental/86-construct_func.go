package main

import "fmt"

// 设计模式 - 选择模式

// 构造结构体：实现动态构造结构体
type Options struct {
	str1 string
	str2 string
	str3 string
	str4 string
	i1   int
	i2   int
	i3   int
	i4   int
}

// 构造类型
type OptionFunc func(opt *Options)

// 通过封包来实现动态初始化结构体
func InitOptions(opts ...OptionFunc) {
	option := &Options{}
	for _, opt := range opts {
		opt(option)
	}
	fmt.Printf("Option is %#v.\n", option)
}

// 封包逐个实现结构体成员初始化
// string类型成员初始化
func initStr1(str string) OptionFunc {
	return func(opt *Options) {
		opt.str1 = str
	}
}

func initStr2(str string) OptionFunc {
	return func(opt *Options) {
		opt.str2 = str
	}
}

func initStr3(str string) OptionFunc {
	return func(opt *Options) {
		opt.str3 = str
	}
}

func initStr4(str string) OptionFunc {
	return func(opt *Options) {
		opt.str4 = str
	}
}

// ...

// int类型成员初始化
func initInt1(i int) OptionFunc {
	return func(opt *Options) {
		opt.i1 = i
	}
}

func initInt2(i int) OptionFunc {
	return func(opt *Options) {
		opt.i2 = i
	}
}

func initInt3(i int) OptionFunc {
	return func(opt *Options) {
		opt.i3 = i
	}
}

func initInt4(i int) OptionFunc {
	return func(opt *Options) {
		opt.i4 = i
	}
}

func main() {
	// 调用函数
	/*
		Option is &main.Options{str1:"test1", str2:"test2", str3:"test3", str4:"", i1:1, i2:2, i3:3, i4:0}.
	*/
	InitOptions(initStr1("test1"), initStr2("test2"), initStr3("test3"),
		initInt1(1), initInt2(2), initInt3(3))
}
