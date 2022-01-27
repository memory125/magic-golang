package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin - 重定向

func main() {
	router := gin.Default()
	router.GET("/to", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	router.Run(":8000")
}
