package main

import "fmt"

func main()  {
	mapDemo()
}

// map
func mapDemo()  {
	// 声明map，这里声明一个键值是string，值是int型的map变量m1
	// 注意：初次声明的map未分配空间，所以不能使用
	var m1 map[string] int

	// 估算map容量，避免在程序运行期间动态扩容
	m1 = make(map[string]int, 10)

	// map使用
	m1["David"] = 35
	m1["Jack"] = 38

	// 输出map
	/*
		map[David:35 Jack:38]
	 */
	fmt.Println(m1)

	// 判断map中有效值
	/*
		找不到这个Key!
	 */
	value, ok := m1["Lily"]
	if !ok {
		fmt.Println("找不到这个Key!")
	} else {
		fmt.Println(value)
	}

	// map的遍历
	/*
		Jack 38
		David 35
	 */
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 删除操作，如果删除的键为nil或者不存在时，delete不做操作
	delete(m1, "Jack")
	/*
		map[David:35]
	 */
	fmt.Println(m1)
}
