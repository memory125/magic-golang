package main

import "fmt"

func main()  {
	traversalString()
}

// 遍历字符串
func traversalString() {
	s := "hello - にいはお"
	/*
		104(h) 101(e) 108(l) 108(l) 111(o) 32( ) 45(-) 32( ) 227(ã) 129() 171(«) 227(ã) 129() 132() 227(ã) 129() 175(¯) 227(ã) 129() 138()
	*/
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	/*
		104(h) 101(e) 108(l) 108(l) 111(o) 32( ) 45(-) 32( ) 12395(に) 12356(い) 12399(は) 12362(お)
	*/
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
