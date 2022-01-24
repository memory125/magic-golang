package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // conf()
)

// MySQL - Insert插入示例

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

// 插入数据
func insertValues(sqlStr string, args ...interface{}) {
	fmt.Println("===========insertValues============")
	ret, err := db.Exec(sqlStr, args...)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Connect database failed, err is %v.\n", err)
	}

	fmt.Println("Connect to database success...")

	sqlStr := "insert into user(name, age) values (?,?)"
	// 创建一个interface类型的切片，存储要相关字段的值
	args := make([]interface{}, 2)
	args[0] = "Jone"
	args[1] = 32
	insertValues(sqlStr, args...)
}
