package main

import (
	"fmt"
	"os"
)

// 信息管理系统

// 书籍
type bookInfo struct {
	id int64
	name string
}

type booksManagerSystem struct {
	// 声明书库，类似与数据库中的表
	allBooks map[int64]*bookInfo
}

// 全局变量
var bms booksManagerSystem

// 构造函数
func newbookInfo(id int64, name string) *bookInfo  {
	return &bookInfo{
		id: id,
		name: name,
	}
}



// 显示所有书籍
func (b booksManagerSystem) showAllBooksFromBookStore() {
	fmt.Println("showAllBooksFromBookStore===========>", b.allBooks)
	for k, v := range b.allBooks {
		fmt.Printf("书的ID：%d, 书名：%s\n", k, v.name)
	}
}

// 添加书籍
func (b booksManagerSystem) addBookToBookStore()  {
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
	newBook := newbookInfo(id, name)
	// 添加到allBooks中
	b.allBooks[id] = newBook
}


// 删除书籍
func (b booksManagerSystem) editBookFromBookStore()  {
	// 根据书籍ID修改相应的书籍
	var id int64
	fmt.Print("请输入您要修改的书籍的ID：")
	fmt.Scanln(&id)

	bookObj, ok := b.allBooks[id]
	if !ok {
		fmt.Printf("找不到ID为%d的书籍！！！\n", id)
		return
	}
	var bookName string
	fmt.Print("请输入您将要书籍名更改为：")
	fmt.Scanln(&bookName)

	// 修改书籍的名字，并保存
	bookObj.name = bookName
	b.allBooks[id] = bookObj

	fmt.Printf("已经将ID为%d的书籍修改为%s！！！\n", id,bookName)
}

// 删除书籍
func (b booksManagerSystem) deleteBookFromBookStore()  {
	// 根据书籍ID删除allBooks中的书籍
	var id int64
	fmt.Print("请输入您要删除书籍的ID：")
	fmt.Scanln(&id)

	_, ok := b.allBooks[id]
	if !ok {
		fmt.Printf("找不到ID为%d的书籍！！！\n", id)
		return
	}

	delete(b.allBooks, id)
	fmt.Printf("ID为%d的书籍删除成功！！！\n", id)
}

func main()  {
	// 初始化，申请空间
	bms = booksManagerSystem{
		allBooks: make(map[int64]*bookInfo, 50),
	}

	for {
		// 打印菜单
		var option int
		fmt.Println("欢迎登录图书管理系统！")
		fmt.Println("1. 查看所有书籍")
		fmt.Println("2. 添加书籍")
		fmt.Println("3. 修改书籍")
		fmt.Println("4. 删除书籍")
		fmt.Println("5. 退出系统")
		fmt.Print("请选择您的操作：")
		fmt.Scanln(&option)
		fmt.Println("请选择是：", option)

		switch option {
		case 1:
			bms.showAllBooksFromBookStore()
		case 2:
			bms.addBookToBookStore()
		case 3:
			bms.editBookFromBookStore()
		case 4:
			bms.deleteBookFromBookStore()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("您的输入有误！请重新输入，谢谢！")
		}
	}
}

