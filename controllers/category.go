package controllers

import (
	"blog/models"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.TplName = "category.html"
	this.Data["IsLogin"] = checkAccount(this.Controller)
	this.Data["IsCategory"] = true

	input := this.Input()

	switch input.Get("option") {
	case "add":
		name := input.Get("name")
		// 如果name字段为空，跳出
		if len(name) == 0 {
			break
		}

		// 操作models，创建分类
		err := models.AddCategory(name)

		// 如果创建分类失败，交给beego error处理
		if err != nil {
			beego.Error(err)
		}

		// 创建成功，重定向回到分类页
		this.Redirect("/category", 301)
		return
	case "del":
		id := input.Get("id")
		// 如果id字段为空，跳出
		if len(id) == 0 {
			break
		}

		err := models.DelCategory(id)

		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 301)
		return
	}

	//获取分类列表
	categoryList, err := models.GetCategoryList()

	this.Data["CategoryList"] = categoryList
	//获取失败，交给beego Error
	if err != nil {
		beego.Error(err)
	}

}
