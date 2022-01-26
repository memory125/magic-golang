package blog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// blog处理函数

func blogPostHandler1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Blog Post"})
}

func blogCommentHandler1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Blog Comment"})
}
