package main

import (
	"fmt"
	"wing.com/magic-golang/web/gin/routersdemo/singlerouters/router"
)

// gin - router拆分示例

func main() {
	r := router.SetupRouter2()
	if err := r.Run(":8000"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
