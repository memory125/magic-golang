package main

import "fmt"

// 面试题

// 结构体Staff
type Staff struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*Staff)
	staffs := []Staff{
		{"Lee", 30},
		{"Jack", 33},
		{"Joe", 36},
	}

	for _, stf := range staffs {
		m[stf.Name] = &stf
		/*
			stf = {Lee 30}.
			stf = {Jack 33}.
			stf = {Joe 36}.
		*/
		fmt.Printf("stf = %v.\n", stf)
	}

	/*
		m = map[Jack:0xc000004078 Joe:0xc000004078 Lee:0xc000004078], len(m) = 3
	*/
	fmt.Printf("m = %v, len(m) = %d\n", m, len(m))

	for k, v := range m {
		/*
			k = Lee, v = &{Joe 36}.
			k = Jack, v = &{Joe 36}.
			k = Joe, v = &{Joe 36}.
		*/
		fmt.Printf("k = %v, v = %v. \n", k, v)
	}
}
