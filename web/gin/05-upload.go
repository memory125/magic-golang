package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 文件上传

func main() {
	router := gin.Default()
	//限制上传最大尺寸
	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}
		// c.JSON(200, gin.H{"message": file.Header.Context})
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})
	router.Run(":8000")
}
