package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模拟blog的路由

func blogPostHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Blog Post"})
}

func blogCommentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Blog Comment"})
}

func LoadBlog(e *gin.Engine) {
	e.GET("/post", blogPostHandler)
	e.GET("/comment", blogCommentHandler)
}
