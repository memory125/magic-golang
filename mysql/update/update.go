package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // conf()
)

// MySQL - Update更新示例

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

// 更新数据
func updateValues(sqlStr string, args ...interface{}) {
	fmt.Println("===========updateValues============")
	ret, err := db.Exec(sqlStr, args...)
	if err != nil {
		fmt.Printf("update data failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Connect database failed, err is %v.\n", err)
	}

	fmt.Println("Connect to database success...")

	sqlStr := "update user set age=? where id = ?"
	// 创建一个interface类型的切片，存储要相关字段的值
	args := make([]interface{}, 2)
	args[0] = 35
	args[1] = 7
	updateValues(sqlStr, args...)
}
