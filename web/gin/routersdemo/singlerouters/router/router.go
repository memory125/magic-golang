package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin - router统一管理路由

func helloHandler3(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello Router",
	})
}

// SetupRouter1 配置路由信息
func SetupRouter2() *gin.Engine {
	router := gin.Default()
	router.GET("/router", helloHandler3)
	return router
}
