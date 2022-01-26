package blog

import (
	"github.com/gin-gonic/gin"
)

// 模拟blog的路由

func Routers(e *gin.Engine) {
	e.GET("/post", blogPostHandler1)
	e.GET("/comment", blogCommentHandler1)
}
