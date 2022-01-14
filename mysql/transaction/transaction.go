package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
)

// MySQL - 事务示例

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

// 事务处理示例
func transactionHandle(sqlStr1, sqlStr2 string) {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("Transaction begin failed, err is %v.\n", err)
		return
	}

	ret1, err := tx.Exec(sqlStr1)
	if err != nil {
		// 执行失败，需要回滚
		tx.Rollback()
		fmt.Printf("SQL 1 execution failed, err is %v.\n", err)
		return
	}

	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	ret2, err := tx.Exec(sqlStr2)
	if err != nil {
		// 执行失败，需要回滚
		tx.Rollback()
		fmt.Printf("SQL 2 execution failed, err is %v.\n", err)
		return
	}

	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	fmt.Println("Affected row: ", affRow1, affRow2)
	// 如果是
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("Transaction committed...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("Transaction rollback...")
	}

	fmt.Println("Transaction execution succeed!!!!!!")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Connect database failed, err is %v.\n", err)
	}

	fmt.Println("Connect to database success...")

	sqlStr1 := "Update user set age=age - 2 where id=3"
	sqlStr2 := "Update user set age=age + 2 where id=4"
	transactionHandle(sqlStr1, sqlStr2)
}
