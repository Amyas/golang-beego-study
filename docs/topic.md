## 基础布局
在开始写业务之前，我们先写一下基础布局和相关路由配置
``` html
<!-- views/topic.html -->

{{template "header"}}
  <title>文章 - 我的 beego 博客</title>
</head>
<body>

  {{template "nav" .}}

  <div class="container" style="margin-top:50px;">
    <h1>文章列表</h1>
    <table class="table table-striped">
      <thead>
        <tr>
          <th>#</th>
          <th>文章名称</th>
          <th>文章内容</th>
          <th>创建时间</th>
          <th>更新时间</th>
          <th>浏览次数</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody>
        {{range .TopicList}}
        <tr>
          <th>{{.Id}}</th>
          <th>{{.Title}}</th>
          <th>{{.Content}}</th>
          <th>{{.Created}}</th>
          <th>{{.Updated}}</th>
          <th>{{.Views}}</th>
          <th>
            <a href="">删除</a>
          </th>
        </tr>
        {{end}}
      </tbody>
    </table>
  </div>
  
{{template "footer"}}
```

``` go
// controllers/topic.go

package controllers

import (
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Controller)
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"
}
```

``` go
// routers/router.go

...

func init() {
  ...

	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
}
```

?> `AutoRouter` 是 beego 为我们提供的一个自动匹配路由的方法，只需要将 `controller` 传入，beego 会自动解析该 controller 下对应的方法。
例如：`topicController` 下有方法 `Add()`，beego 就会自动创建 `/topic/add` 路由。
文章的新增、编辑与删除我们将使用这种方法

## 创建文章
``` html
<!-- views/topic_add.html -->

...
  <div class="container" style="margin-top:50px;">
    <h1>新增文章</h1>
    <form method="POST" action="/topic">
      <div class="form-group">
        <label>文章名称</label>
        <input type="text" name="title" class="form-control">
      </div>
      <div class="form-group">
        <label>文章内容</label>
        <textarea name="content" cols="30" rows="10" class="form-control"></textarea>
      </div>
      <button class="btn btn-primary">新增文章</button>
    </form>
  </div>
...
```

``` go
// controllers/topic.go

...

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
...
```

``` go
// models/models.go

...
func AddTopic(title, content string) error {
	o := orm.NewOrm()

	topic := &Topic{
		Title:     title,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}

	_, err := o.Insert(topic)
	return err
}
```

## 获取文章列表
``` go
// controllers/topic.go

...
func (this *TopicController) Get() {
  ...

	var err error
	this.Data["TopicList"], err = models.GetTopicList()
	if err != nil {
		beego.Error(err)
	}
}
```

``` go
// models/models.go

...
func GetTopicList() ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	_, err := qs.All(&topics)
	return topics, err
}
```

## 删除文章
``` html
<!-- views/topic.html -->

...
<tbody>
  ...
    <th>
      <a href="/topic/delete/{{.Id}}">删除</a>
    </th>
  ...
</tbody>
...
```

``` go
// controllers/topic.go

...
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
...
```

``` go
// models/models.go

...
func DelTopic(id string) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: tid}
	_, err = o.Delete(topic)
	return err
}
```

## 查看文章
``` html
<!-- views/topic_view.html -->

...
{{template "header"}}
  <title>{{.Topic.Title}} - 我的 beego 博客</title>
</head>
<body>

  {{template "nav" .}}

  <div class="container" style="margin-top:50px;">
    <h1>{{.Topic.Title}}</h1>
    <div>{{.Topic.Content}}</div>
  </div>
  
{{template "footer"}}
```

``` go
// controllers/topic.go

...
func (this *TopicController) View() {
	this.Data["IsLogin"] = checkAccount(this.Controller)
	this.Data["IsTopic"] = true
	this.TplName = "topic_view.html"

	var err error
	this.Data["Topic"], err = models.GetTopic(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
	}
}
```

``` go
// models/models.go

...
func GetTopic(id string) (*Topic, error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: tid}

	qs := o.QueryTable("topic")
	err = qs.Filter("title", tid).One(topic)
	return topic, err
}
```

## 编辑文章