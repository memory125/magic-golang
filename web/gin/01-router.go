package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// router - 路由
/*
   gin 框架中采用的路由库是基于httprouter做的
   https://github.com/julienschmidt/httprouter
*/

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})
	r.POST("/xxxpost", nil)
	r.PUT("/xxxput", nil)
	//监听端口默认为8080
	r.Run(":8000")
}
