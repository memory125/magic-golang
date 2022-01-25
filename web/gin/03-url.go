package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin - URL参数

func main() {
	router := gin.Default()
	router.GET("/user", func(c *gin.Context) {
		//指定默认值
		//http://localhost:8080/user 才会打印出来默认的值
		name := c.DefaultQuery("name", "Jacky")
		c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
	})

	router.Run(":8000")
}
