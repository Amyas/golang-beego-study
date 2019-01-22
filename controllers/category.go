package controllers

import (
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.TplName = "category.html"
	this.Data["IsLogin"] = checkAccount(this.Controller)
	this.Data["IsCategory"] = true
}
