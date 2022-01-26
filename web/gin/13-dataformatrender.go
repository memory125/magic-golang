package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

// gin - 各种数据格式的响应

// 多种响应方式
func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	router := gin.Default()
	// 1.json
	router.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "json rendering", "status": 200})
	})
	// 2. 结构体响应
	router.GET("/struct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "struct"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(200, msg)
	})
	// 3.XML
	router.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{"msg": "xml rendering"})
	})
	// 4.YAML响应
	router.GET("/yaml", func(c *gin.Context) {
		c.YAML(200, gin.H{"msg": "yaml rendering"})
	})
	// 5.protobuf格式,谷歌开发的高效存储读取的工具
	// 数组？切片？如果自己构建一个传输格式，应该是什么格式？
	router.GET("/protobuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		label := "protobuf"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})

	router.Run(":8000")
}
