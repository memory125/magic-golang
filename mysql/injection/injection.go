package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // conf()
	"github.com/jmoiron/sqlx"
)

// MySQL - SQL注入示例

// MySQL - sqlx示例

// 全局变量
var db *sqlx.DB

// user结构体
type user struct {
	Id   int
	Name string
	Age  int
}

func initDB() (err error) {
	// 数据库连接信息
	dataSource := "root:123456@tcp(192.168.40.133:3306)/go_test"

	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dataSource)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20) // 设置最大可以打开的连接数
	db.SetMaxIdleConns(10) // 设置最大的空闲连接数

	return
}

// SQL注入示例
func sqlInjection(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var u []user
	err := db.Select(&u, sqlStr)
	if err != nil {
		fmt.Printf("Get data failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", u)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Connect database failed, err is %v.\n", err)
	}

	fmt.Println("Connect to database success...")

	// 这里会查询到所有的user信息
	sqlInjection("xxx' or 1=1 #")
}
