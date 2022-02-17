package main

import "fmt"

// 面试题

func swap(a, b *int) (*int, *int) {
	a, b = b, a
	return a, b
}

func main() {
	a, b := 5, 6
	/*
		1 === &a = 0xc00000a098, &b = 0xc00000a0b0.
	*/
	fmt.Printf("1 === &a = %v, &b = %v.\n", &a, &b)
	c, d := swap(&a, &b)
	/*
		2 === &c = 0xc00000a0b0, &d = 0xc00000a098.
	*/
	fmt.Printf("2 === &c = %v, &d = %v.\n", c, d)
	/*
	*c = 6, *d = 5
	 */
	fmt.Printf("*c = %v, *d = %v\n", *c, *d)

	a = *c // 这里是关键一步
	b = *d
	/*
		3 === &a = 0xc00000a098, &b = 0xc00000a0b0.
	*/
	fmt.Printf("3 === &a = %v, &b = %v.\n", &a, &b)
	/*
		a = 6, b = 6
	*/
	fmt.Printf("a = %v, b = %v\n", a, b)
}
