package main

import "fmt"

// 结构体示例

type Person2 struct {
	name, gender string
}

// Go语言中函数传递参数永远传的是拷贝
func update1(person Person2)  {
	person.gender = "Female"        // 此处修改的是副本的gender
}

func update2(person *Person2)  {
	// (*person).gender = "Female"  根据内存地址找到原变量，修改的就是原来的变量
	person.gender = "Female"      // 语法糖，自动根据指针找对应的变量
}

func main()  {
	var p Person2
	p.name = "Kevin"
	p.gender = "Male"
	/*
		Type: main.Person2, Value: {Kevin Male}
	*/
	fmt.Printf("Type: %T, Value: %v \n", p, p)

	update1(p)
	/*
		Male
	 */
	fmt.Println(p.gender)
	update2(&p)
	/*
		Female
	 */
	fmt.Println(p.gender)

	// 返回结构体Person2的指针
	var p2 = new(Person2)
	/*
		Type: *main.Person2
		Value: &{ }
	 */
	fmt.Printf("Type: %T\n", p2)
	fmt.Printf("Value: %x\n", p2)
}
