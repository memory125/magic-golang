package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin - 路由注册

// 处理函数
func routerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello wing!",
	})
}

func main() {
	router := gin.Default()
	router.GET("/wing", routerHandler)
	err := router.Run(":8000")
	if err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
