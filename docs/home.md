## 删除默认配置

beego为我们提供了一个基础的模板，我们需要修改一些自定义的内容

* 删除 `controllers` 目录下的 `default.go`
* 删除 `views` 目录下的 `index.tpl`
* 删除 `routers` 目录下的 `router.go` 文件中的 `init()` 函数内容
* 删除 `static` 目录下的 `js` 目录下的 `reload.min.js`

## 创建控制器(controller)

在 `controllers` 目录下 创建 `home.go` 文件

``` go
package controllers

import (
	"github.com/astaxie/beego"
)

// 声明一个 `HomeController` 结构体
type HomeController struct {
	// 将 `beego.Controller` 中的方法 挂载到 `HomeController` 上
	beego.Controller
}

// 覆盖 `beego.Controller` 中的 `Get` 方法
func (this *HomeController) Get() {
	// 指定首页的模板文件
	// 模板文件存下于 `views` 目录中
	this.TplName = "index.html"
}
```

## 创建路由(router)

在 `routers` 目录下的 `router.go` 文件中添加首页路由
``` go
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
}
```

## 创建模板(template)

在 `views` 目录下 创建 `index.html` 文件

``` html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>首页 - 我的 beego 博客</title>
</head>
<body>
  首页
</body>
</html>
```

打开 [localhost:8080](http://localhost:8080/) 看到你输入的html内容就ok啦！

## 引入bootstrap和jquery

该项目主要以学习 `beego` 框架为主，所以就直接用 `bootstrap`，简单又方便~

* [bootstrap3下载](https://github.com/twbs/bootstrap/releases/download/v3.3.7/bootstrap-3.3.7-dist.zip)

下载解压后，将 `bootstrap` 中

* `css` 目录下的 `bootstrap.min.css` 
* `js` 目录下的 `bootstrap.min.js` 

放到 `你的项目/static/对应目录(css || js)`

?> `/static` 为项目静态文件目录，项目中的 `css` `img` `js` 都放在该目录下。

在 `index.html` 中引入 `jquery` 和 `bootstrap`
``` html
<head>
	...
  <link rel="stylesheet" href="/static/css/bootstrap.min.css">
  <title>blog</title>
</head>
<body>
  首页

  <script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
  <script src="/static/js/bootstrap.min.js"></script>
</body>
```

## 创建导航栏

在引入 `bootstrap` 后我们就开始进行导航栏的实现

``` html
<!-- index.html -->
...
<body>
  <!-- 导航栏 -->
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
          <li><a href="/login">管理员登录</a></li>
        </ul>
      </div>
      
    </div>
  </div>
	...
</body>
```

打开 [localhost:8080](http://localhost:8080/) 看到如下效果就ok啦！

![navbar](_images/navbar.png 'navbar')