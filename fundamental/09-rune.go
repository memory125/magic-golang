package main

import "fmt"

func main()  {
	fmt.Println("================= traversalString ==============")
	traversalString()

	fmt.Println("================= modifyString ==============")
	modifyString()
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

// 修改字符串
func modifyString()  {
	s1 := "Hello"
	s2 := []rune(s1)             // 把字符串s1强制转换成了rune切片，此时： s2 ====> ['H', 'e', 'l', 'l', 'o']
	s2[0] = 'W'                  // s2 ====> ['W', 'e', 'l', 'l', 'o']

	/*
		Hello
		Wello
	 */
	fmt.Println(s1)
	fmt.Println(string(s2))
}
