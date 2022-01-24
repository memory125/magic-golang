package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // conf()
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

// sqlx 插入数据
func insertRow(sqlStr string, args ...interface{}) {
	ret, err := db.Exec(sqlStr, args...)
	if err != nil {
		fmt.Printf("Insert data failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("Get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("Insert data success, the id is %d.\n", theID)
}

// 更新数据
func updateRow(sqlStr string, args ...interface{}) {
	ret, err := db.Exec(sqlStr, args...)
	if err != nil {
		fmt.Printf("Update data failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("Get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("Update data success, affected rows:%d\n", n)
}

// 删除数据
func deleteRow(sqlStr string, args interface{}) {
	ret, err := db.Exec(sqlStr, args)
	if err != nil {
		fmt.Printf("Delete data failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("Get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("Delete data success, affected rows:%d\n", n)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Connect database failed, err is %v.\n", err)
	}

	fmt.Println("Connect to database success...")

	// 查询单个数据
	sqlStr1 := "select id, name, age from user where id = ?"
	sqlxQuerySingle(sqlStr1, 1)

	// 查询多个数据
	sqlStr2 := "select id, name, age from user where id > ?"
	sqlxQueryMulti(sqlStr2, 0)

	// 插入数据
	sqlStr3 := "insert into user(name, age) values (?,?)"
	data := []interface{}{
		"Peter",
		34,
	}
	insertRow(sqlStr3, data...)

	// 查询多个数据
	sqlxQueryMulti(sqlStr2, 0)

	// 更新数据
	sqlStr4 := "update user set age=? where id = ?"
	data = []interface{}{
		37,
		17, // "Peter", 37
	}
	updateRow(sqlStr4, data...)

	// 查询多个数据
	sqlxQueryMulti(sqlStr2, 0)

	// 删除数据
	sqlStr5 := "delete from user where id = ?"
	deleteRow(sqlStr5, 17)
}
