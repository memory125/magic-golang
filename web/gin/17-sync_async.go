package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// gin - 异步 & 同步

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	router := gin.Default()
	// 1.异步
	router.GET("/async", func(c *gin.Context) {
		// 需要搞一个副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
	})

	// 2.同步
	router.GET("/sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
		c.JSON(http.StatusOK, gin.H{"msg": "sync data"})
	})

	router.Run(":8000")
}
