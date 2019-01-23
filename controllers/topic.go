package controllers

import (
	"blog/models"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Controller)
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"

	var err error
	this.Data["TopicList"], err = models.GetTopicList()
	if err != nil {
		beego.Error(err)
	}
}

func (this *TopicController) Post() {
	input := this.Input()

	title := input.Get("title")
	content := input.Get("content")

	err := models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 301)
}

func (this *TopicController) Add() {
	this.Data["IsLogin"] = checkAccount(this.Controller)
	this.Data["IsTopic"] = true
	this.TplName = "topic_add.html"
}

func (this *TopicController) Delete() {
	// 获取id
	// 我们在 `topic.html` 中的删除按钮地址为： `/topic/delete/{{.Id}}`
	// 获取delete后面的参数方法为：`this.Ctx.Input.Param("key")`
	id := this.Ctx.Input.Param("0")

	err := models.DelTopic(id)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 301)
}
