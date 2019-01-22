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