package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin - html模板渲染

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/common/*")
	router.GET("/common", func(c *gin.Context) {
		c.HTML(http.StatusOK, "common.html", gin.H{"title": "Template测试", "common": "渲染", "ce": "123456"})
	})
	router.Run(":8000")
}
