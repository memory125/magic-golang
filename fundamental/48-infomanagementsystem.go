package main

import (
	"fmt"
	"os"
)

// 信息管理系统

// 书籍
type book struct {
	id int64
	name string
}

// 构造函数
func newBook(id int64, name string) *book  {
	return &book{
		id: id,
		name: name,
	}
}

var (
	// 声明书籍存储的变量，类似与数据库中的表
	allBooks map[int64]*book
)

// 显示所有书籍
func showAllBooks() {
	fmt.Println("showAllBooks===========>", allBooks)
	for k, v := range allBooks {
		fmt.Printf("书的ID：%d, 书名：%s\n", k, v.name)
	}
}

// 添加书籍
func addBook()  {
	// 向allBooks中添加新的书籍
	var (
		id int64
		name string
	)
	fmt.Print("请输入书籍的ID：")
	fmt.Scanln(&id)
	fmt.Print("请输入书名：")
	fmt.Scanln(&name)

	// 根据输入的信息构造book结构体实例
	newBook := newBook(id, name)
	// 添加到allBooks中
	allBooks[id] = newBook
}

// 删除书籍
func deleteBook()  {
	// 根据书籍ID删除allBooks中的书籍
	var id int64
	fmt.Print("请输入您要删除书籍的ID：")
	fmt.Scanln(&id)

	delete(allBooks, id)
}

func main()  {
	// 初始化，申请空间
	allBooks = make(map[int64]*book, 50)
	for {
		// 打印菜单
		var option int
		fmt.Println("欢迎登录图书管理系统！")
		fmt.Println("1. 查看所有书籍")
		fmt.Println("2. 添加书籍")
		fmt.Println("3. 删除书籍")
		fmt.Println("4. 退出系统")
		fmt.Print("请选择您的操作：")
		fmt.Scanln(&option)
		fmt.Println("请选择是：", option)

		switch option {
		case 1:
			showAllBooks()
		case 2:
			addBook()
		case 3:
			deleteBook()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("您的输入有误！请重新输入，谢谢！")
		}
	}
}

