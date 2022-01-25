package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin - 上传多个文件

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	router := gin.Default()
	// 限制表单上传大小 8MB，默认为32MB
	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", func(c *gin.Context) {
		// 多个表单
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有文件
		files := form.File["files"]
		// 遍历所有文件
		for _, file := range files {
			// 一个一个保存
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})
	//默认端口号是8080
	router.Run(":8000")
}
