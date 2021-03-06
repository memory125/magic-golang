package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin - Hello World

/*
gin框架：
  安装：go get -u github.com/gin-gonic/gin
*/

func main() {
	// 1.创建路由
	router := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	router.GET("/wing", func(c *gin.Context) {
		c.String(http.StatusOK, "The first sub web page!")
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	router.Run(":8000")
}
