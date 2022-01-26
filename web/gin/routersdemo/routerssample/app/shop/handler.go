package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// shop处理函数

func shopHomeHandler1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Shop Home"})
}

func shopGoodsHandler1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Shop Goods"})
}

func shopCheckoutHandler1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Shop Checkout"})
}
