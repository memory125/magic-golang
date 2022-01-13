package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
)

// MySQL - Query查询示例

// 全局变量
var db *sql.DB

// user结构体
type user struct {
	id   int
	name string
	age  int
}

func initDB() (err error) {
	// 数据库连接信息
	dataSource := "root:123456@tcp(192.168.40.133:3306)/go_test"
	// 连接数据库
	/*
		备注：
		   open()：不会校验用户名和密码是否正确
		   dataSource：格式不正确呃时候会报错
	*/
	db, err = sql.Open("mysql", dataSource)
	if err != nil {
		return
	}

	// 尝试连接数据库
	err = db.Ping()
	if err != nil {
		return
	}

	return
}

// 查询单个
func query(queryStr string, args ...interface{}) {
	var u user

	row := db.QueryRow(queryStr, args...)
	err := row.Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("row scan failed, error is %v.\n", err)
		return
	}

	fmt.Printf("user information: %#v\n", u)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Connect database failed, err is %v.\n", err)
	}

	fmt.Println("Connect to database success...")
	str := "select id, name, age from user where id = ?;"
	query(str, 1)
	query(str, 3)
	query(str, 5)
}
