package routers

import (
	// 将 `controllers` 包引入
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 参数1：路由地址(route path)
	// 参数2：控制器(controller)
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
}
