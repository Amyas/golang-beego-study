package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	// 获取query string `exit` 是否为 true
	isExit := this.Input().Get("exit") == "true"

	// 如果 `isExit` 为true，把cookie清空，并将maxAge 设置为 -1，然后重定向回首页
	if isExit {
		this.Ctx.SetCookie("uname", "1", -1, "/")
		this.Ctx.SetCookie("pwd", "1", -1, "/")
		this.Redirect("/", 301)
		return
	}
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	// beego.AppConfig 为 `conf` 目录下的 `app.cong` 文件内容
	config := beego.AppConfig
	// `this.Input()` 为获取表单
	input := this.Input()

	// 获取指定字段的方式为：`input.Get("key")`
	uname := input.Get("uname")
	pwd := input.Get("pwd")
	// checkbox 类型的 input 如果勾选会返回字符串 `on`
	autoLogin := input.Get("autoLogin") == "on"

	// 判断用户输入的 `uname` `pwd` 与 `app.conf` 中的是否相等
	// 相等代表输入正确，反之错误直接返回首页
	if config.String("uname") == uname &&
		config.String("pwd") == pwd {
		// 设置cookie的时间
		maxAge := 0

		//如果是自动登录，给一个超长时间
		if autoLogin {
			maxAge = 1<<31 - 1
		}

		//设置cookie
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}

	// 重定向到首页
	this.Redirect("/", 301)
}

// 判断是否登录
func checkAccount(controller beego.Controller) bool {
	config := beego.AppConfig

	// 获取用户请求cookie中是否带有uname
	ck, err := controller.Ctx.Request.Cookie("uname")
	// 没有uname就是没登录
	if err != nil {
		return false
	}
	// 有uname就赋值
	uname := ck.Value

	// 获取用户请求cookie中是否带有pwd
	ck, err = controller.Ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	// 校验 `app.conf` 中的配置和用户cookie是否一致
	return config.String("uname") == uname &&
		config.String("pwd") == pwd
}
