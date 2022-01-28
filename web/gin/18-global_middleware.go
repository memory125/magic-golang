package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// gin - 全局中间件

// 定义中间件
func GlobalMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("Global() ===== 中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "Global() ===== 中间件")
		status := c.Writer.Status()
		fmt.Println("Global() ===== 中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	router := gin.Default()
	// 注册中间件
	router.Use(GlobalMiddleWare())
	// {}为了代码规范
	{
		router.GET("/mid", func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			// 页面接收
			c.JSON(200, gin.H{"request": req})
		})

	}
	router.Run(":8000")
}
