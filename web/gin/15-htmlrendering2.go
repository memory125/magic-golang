package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin - html模板渲染

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/**/*")
	router.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.html", gin.H{"title": "用户", "address": "www.test.com"})
	})
	router.Run(":8000")
}
