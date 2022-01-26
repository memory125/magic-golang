package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wing.com/magic-golang/web/gin/routersdemo/multirouters/routers"
)

// gin - multi routers(多路由)示例

func main() {
	r := gin.Default()
	routers.LoadBlog(r)
	routers.LoadShop(r)
	if err := r.Run(":8000"); err != nil {
		fmt.Println("Start service failed, err:%v\n", err)
	}
}
