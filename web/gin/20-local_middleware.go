package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// gin - 局部中间件

// 定义中间件
func LocalMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("Local() ====== 中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "Local() ======中间件")
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("Local() ======中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	router := gin.Default()
	//局部中间键使用
	router.GET("/local", LocalMiddleWare(), func(c *gin.Context) {
		// 取值
		req, exist := c.Get("request")
		if !exist {
			fmt.Println("the key is not exist!")
			return
		}
		fmt.Println("request:", req)
		// 页面接收
		c.JSON(200, gin.H{"request": req})
	})
	router.Run(":8000")
}
