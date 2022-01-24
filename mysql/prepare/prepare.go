package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
)

// MySQL - 预处理

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

// 预处理查询示例
func prepareQuery(sqlStr string, args ...interface{}) {
	fmt.Println("==============prepareQuery=============")
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}

	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}

	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("user information: id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 预处理插入示例
func prepareInsert(sqlStr string, args ...interface{}) {
	fmt.Println("==============prepareInsert=============")
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	fmt.Println("insert data success.")
}

// 预处理更新和删除的操作和插入数据相似
// 预处理更新
func prepareUpdate(sqlStr string, args ...interface{}) {
	fmt.Println("==============prepareUpdate=============")
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}

	fmt.Println("update data success.")
}

// 预处理删除
func prepareDelete(sqlStr string, args ...interface{}) {
	fmt.Println("==============prepareDelete=============")
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		fmt.Printf("delete data failed, err:%v\n", err)
		return
	}

	fmt.Println("delete data success.")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Connect database failed, err is %v.\n", err)
	}

	fmt.Println("Connect to database success...")

	sqlStr := "select id, name, age from user where id > ?"
	prepareQuery(sqlStr, 0)

	sqlStr = "insert into user(name, age) values (?,?)"
	data := map[string]int{
		"Devin":  31,
		"Eric":   32,
		"Law":    30,
		"Edward": 36,
		"Blues":  34,
		"Sky":    25,
		"Sean":   27,
	}
	for k, v := range data {
		prepareInsert(sqlStr, k, v)
	}

}
