package controllers

import (
	"github.com/astaxie/beego"
)

// 声明一个 `HomeController` 结构体
type HomeController struct {
	// 将 `beego.Controller` 中的方法 挂在到 `HomeController` 上
	beego.Controller
}

// 覆盖 `beego.Controller` 中的 `Get` 方法
func (this *HomeController) Get() {
	// 指定首页的模板文件
	// 模板文件存下于 `views` 目录中
	this.TplName = "index.html"

	this.Data["IsLogin"] = checkAccount(this.Controller)
	this.Data["IsHome"] = true
}
