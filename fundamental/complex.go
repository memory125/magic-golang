package main

import "fmt"

func main()  {
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 3 + 4i

	/*
		(1+2i)
		(3+4i)
	 */
	fmt.Println(c1)
	fmt.Println(c2)
}
