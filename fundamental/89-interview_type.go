package main

import "fmt"

// 面试 - type

// 声明类型
type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}

// 添加元素 Add
func (s Slice) AddElement(elem int) *Slice {
	s = append(s, elem)
	fmt.Print(elem)
	return &s
}

func main() {
	s := NewSlice()
	/*
		12435
	*/
	defer s.AddElement(1).AddElement(2).AddElement(4).AddElement(5)
	s.AddElement(3)
}
