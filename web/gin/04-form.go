package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin - 表单参数

func main() {
	router := gin.Default()
	// post请求
	router.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("password")
		// c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})
	router.Run(":8000")
}
