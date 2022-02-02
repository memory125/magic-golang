package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin - Cookie demo

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("admin"); err == nil {
			if cookie == "admin" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "AuthMiddleWare ==== error"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}

func main() {
	// 1.创建路由
	router := gin.Default()
	router.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("admin", "admin", 300, "/",
			"localhost", false, true)
		// 返回信息
		c.String(200, "Login success!")
	})
	router.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})
	router.Run(":8000")
}
