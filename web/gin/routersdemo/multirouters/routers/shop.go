package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模拟shop的路由

func shopHomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Shop Home"})
}

func shopGoodsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Shop Goods"})
}

func shopCheckoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Shop Checkout"})
}

func LoadShop(e *gin.Engine) {
	e.GET("/home", shopHomeHandler)
	e.GET("/goods", shopGoodsHandler)
	e.GET("/checkout", shopCheckoutHandler)
}
