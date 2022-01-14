package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
	"github.com/jmoiron/sqlx"
)

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

// sqlx 查询：单个
func sqlxQuerySingle(sqlStr string, args interface{}) {
	var u user
	err := db.Get(&u, sqlStr, args)
	if err != nil {
		fmt.Printf("Get date failed, err:%v\n", err)
		return
	}
	fmt.Printf("User == id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
}

// sqlx 查询：多个
func sqlxQueryMulti(sqlStr string, args interface{}) {
	var u []user
	err := db.Select(&u, sqlStr, args)
	if err != nil {
		fmt.Printf("Select data failed, err:%v\n", err)
		return
	}
	fmt.Printf("User Information: %v\n", u)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Connect database failed, err is %v.\n", err)
	}

	fmt.Println("Connect to database success...")

	sqlStr1 := "select id, name, age from user where id = ?"

	sqlxQuerySingle(sqlStr1, 1)

	sqlStr2 := "select id, name, age from user where id > ?"
	sqlxQueryMulti(sqlStr2, 0)
}
