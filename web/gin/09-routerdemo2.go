package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin -router拆分

// routerHandler 路由处理函数
func routerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello Router World!",
	})
}

// 构造路由
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/test", routerHandler)
	return r
}

func main() {
	r := setupRouter()
	if err := r.Run(":8000"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
