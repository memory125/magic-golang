package shop

import (
	"github.com/gin-gonic/gin"
)

// 模拟shop的路由

func Routers(e *gin.Engine) {
	e.GET("/home", shopHomeHandler1)
	e.GET("/goods", shopGoodsHandler1)
	e.GET("/checkout", shopCheckoutHandler1)
}
