## 抽离导航栏为公用模板

在 `views` 目录下创建 `category.html` 模板文件

`category.html` 文件也存在导航栏，所以我们需要把导航栏抽离出来。

在 `views` 目录下创建 `template.nav.tpl` 模板文件

``` html
<!-- views/template.nav.tpl -->

{{define "nav"}}
<div class="navbar navbar-default navbar-fixed-top">
  <div class="container">
    
    <a class="navbar-brand" href="/">我的博客</a>
    
    <ul class="nav navbar-nav">
      <li class="active"><a href="/">首页</a></li>
      <li><a href="/category">分类</a></li>
      <li><a href="/topic">文章</a></li>
    </ul>

    <div class="pull-right">
      <ul class="nav navbar-nav">
        <li>
          {{if .IsLogin}}
            <a href="/login?exit=true">退出</a>
          {{else}}
            <a href="/login">管理员登录</a>
          {{end}}
        </li>
      </ul>
    </div>
    
  </div>
</div>
{{end}}
```

在 `views` 目录下的 `index.html` 、 `category.html` 中 引入该模板文件

``` html
<!-- views/category.tpl -->

{{template "header"}}
  <title>分类 - 我的 beego 博客</title>
</head>
<body>

  {{template "nav" .}}
  
{{template "footer"}}
```

?> `{{template "nav" .}}` 最后面有一个 `.` ，这个 `.` 的意思就是把当前页面 `controllers` 中的 `Data` 传入 `nav` 模板文件

## 动态显示导航栏Actice

到现在为止，导航栏的 `Actice` 还是写死在 `index.html` 上的

接下来我们实现动态 `Actice`

``` go
// controllers/category.go

...
func (this *CategoryController) Get() {
  ...
	this.Data["IsLogin"] = checkAccount(this.Controller)
	this.Data["IsCategory"] = true
}
```

``` html
<!-- views/category.html -->

...
<ul class="nav navbar-nav">
  <li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
  <li {{if .IsCategory}}class="active"{{end}}><a href="/category">分类</a></li>
  <li {{if .IsTopic}}class="active"{{end}}><a href="/topic">文章</a></li>
</ul>
...
```

这里只演示了 `category` 分页的写法，其他项目需求页面，逻辑相同即可


打开 [localhost:8080/category](http://localhost:8080/category) 看到如下效果就ok啦！

![navactive](_images/navactive.png 'navactive')

## 创建分类

``` html
<!-- views/category.html -->

...
<div class="container" style="margin-top:50px;">
  <h1>分类列表</h1>
  <form method="GET" action="/category">
    <div class="form-group">
      <label>分类名称</label>
      <input id="name" name="name" class="form-control" placeholder="分类名称">
    </div>
    <input type="hidden" name="option" value="add">
    <button type="submit" class="btn btn-default">添加</button>
  </form>
</div>
...
```

``` go
// models/models.go

...
// 创建分类
func AddCategory(name string) error {
	// 创建一个orm模型
	o := orm.NewOrm()

	// 创建一个分类
	cate := &Category{
		Title:     name,
		Created:   time.Now(),
		TopicTime: time.Now(),
	}

	// 获取 `category` 表
	qs := o.QueryTable("category")
	// 查看表中是否已经存在相同分类
	err := qs.Filter("title", name).One(cate)

	// 如果没有err，就代表已经存在相同分类，return nil 即可
	if err == nil {
		return err
	}

	// 向表中拆入一个分类
	_, err = o.Insert(cate)

	// 如果err != nil ，就代表插入失败了，return err
	if err != nil {
		return err
	}

	// 成功插入，return nil
	return nil
}
...
```

``` go
// controllers/category.go
...
func (this *CategoryController) Get() {
  ...
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
	}

}
```

## 获取分类

``` html
<!-- views/category.html -->
...
<div class="container" style="margin-top:50px;">
  ...
  <table class="table table-striped">
    <thead>
      <tr>
        <th>#</th>
        <th>名称</th>
        <th>文章数</th>
        <th>操作</th>
      </tr>
    </thead>
    <tbody>
      {{range .CategoryList}}
      <tr>
        <th>{{.Id}}</th>
        <th>{{.Title}}</th>
        <th>{{.TopicCount}}</th>
        <th>
          <a href="/category?option=del&id={{.Id}}">删除</a>
        </th>
      </tr>
      {{end}}
    </tbody>
  </table>
</div>
```

``` go
// models/models.go

// 获取列表
func GetCategoryList() ([]*Category, error) {
	// 创建一个orm模型
	o := orm.NewOrm()
	// 创建一个 `Category` 切片
	cates := make([]*Category, 0)
	// 获取 `category` 表
	qs := o.QueryTable("category")
	// 获取所有数据
	_, err := qs.All(&cates)
	return cates, err
}
```

``` go
// controllers/category.go

func (this *CategoryController) Get() {
	...

	//获取分类列表
	categoryList, err := models.GetCategoryList()

	this.Data["CategoryList"] = categoryList
	//获取失败，交给beego Error
	if err != nil {
		beego.Error(err)
	}

}

```

## 删除分类

``` go
// models/models.go

// 删除分类
func DelCategory(id string) error {
	//将id 转化为 int64
	cid, err := strconv.ParseInt(id, 10, 64)

	// 转化失败，return err
	if err != nil {
		return err
	}

	// 创建一个orm模型
	o := orm.NewOrm()

	// 构造一个分类
	cate := &Category{Id: cid}

	// 删除相同id的分类
	_, err = o.Delete(cate)

	// 删除失败，return err
	if err != nil {
		return err
	}
	return nil
}
```

``` go
// controllers/category.go

...
func (this *CategoryController) Get() {
	...

	switch input.Get("option") {
	...
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
  ...
}
```