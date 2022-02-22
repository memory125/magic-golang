package main

import "github.com/astaxie/beego"

// beego - hello world

// controller结构体
type HomeController struct {
	beego.Controller
}

// controller函数
func (c *HomeController) Get() {
	c.Ctx.WriteString("Beego, Hello World!")
}

func main() {
	// 注册路由
	beego.Router("/", &HomeController{})
	// 监听端口
	beego.Run(":8000")
}
