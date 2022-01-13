package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
)

// MySQL - 连接示例

/*
步骤：
   1. 下载mysql driver：go get -u github.com/go-sql-driver/mysql
      dataSource格式：username:password@tcp(database server ip:port)/database name
   2. 创建go文件，调用database/sql包面的接口
*/

func main() {
	// 数据库连接信息
	dataSource := "root:123456@tcp(192.168.40.133:3306)/go_test"
	// 连接数据库
	/*
		备注：
		   open()：不会校验用户名和密码是否正确
		   dataSource：格式不正确呃时候会报错
	*/
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		fmt.Printf("dataSource is invalid, error is %v.\n", err)
		return
	}

	// 尝试连接数据库
	err = db.Ping()
	if err != nil {
		fmt.Printf("open database failed, error is %v.\n", err)
		return
	}
	fmt.Println("Connect to database success...")
}
