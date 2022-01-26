package main

import (
	"fmt"
	"wing.com/magic-golang/web/gin/routersdemo/routerssample/app/blog"
	"wing.com/magic-golang/web/gin/routersdemo/routerssample/app/shop"
	"wing.com/magic-golang/web/gin/routersdemo/routerssample/routers"
)

// router sample测试

func main() {
	// 加载多个APP的路由配置
	routers.Include(shop.Routers, blog.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(":8000"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
