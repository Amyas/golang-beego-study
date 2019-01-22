## 创建控制器(controller)
和首页相同，直接贴代码
``` go
// controllers/login.go
package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}
```

## 创建路由(router)
和首页相同，直接贴代码
``` go
// routers/router.go
...
func init() {
  ...
	beego.Router("/login", &controllers.LoginController{})
}
```

## 创建模板(template)
和首页项目，直接贴代码

?> `html/css/js` 不是本次学习的重点，所以基础部分会直接贴代码，需要注意的地方会提示。
``` html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <link rel="stylesheet" href="/static/css/bootstrap.min.css">
  <title>登录 - 我的 beego 博客</title>
</head>
<body>
  <div class="container" style="width:500px;margin-top: 200px;">
    <form method="POST" action="/login">
      <div class="form-group">
        <label>账号</label>
        <input id="uname" name="uname" class="form-control" placeholder="账号">
      </div>
      <div class="form-group">
        <label>密码</label>
        <input id="pwd" name="pwd" type="password" class="form-control" placeholder="密码">
      </div>
      <div class="checkbox">
        <label>
          <input name="autoLogin" type="checkbox"> 自动登录
        </label>
      </div>
      <button type="submit" class="btn btn-default">登录</button>
    </form>
  </div>

  <script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
  <script src="/static/js/bootstrap.min.js"></script>
</body>
</html>
```

* 这里我们创建了一个表单 `form`
* 提交方式 `method` 为 `POST`
* 提交地址为 `/login`
* 一共 `3` 个字段，分别为 账号`uname`、密码`pwd`、自动登录`autoLogin`

## 提取公用部分

写到这里有一个问题，需要我们解决一下。

首页的 `header` 部分和 登录页的 大致相同。

我们可以把这部分`抽象出来`，作为一个单独的模板。所有用到这部分的都引入这个模板就好了。

接下来我们就开始干这件事。

#### 模板使用方法
  * 创建模板文件：`.tpl` 后缀为模板文件
  * 定义模板文件：`{{define "template name"}} 模板内容 {{end}}` 
  * 使用模板文件：`{{template "template name"}}`

在 `views` 目录下创建 `template.header.tpl` 和 `template.footer.tpl` 文件

``` html
<!-- views/template.header.tpl -->
{{define "header"}}
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <link rel="stylesheet" href="/static/css/bootstrap.min.css">
{{end}}
```

``` html
<!-- views/template.footer.tpl -->
{{define "footer"}}
  <script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
  <script src="/static/js/bootstrap.min.js"></script>
</body>
</html>
{{end}}
```

使用模板（只展示登录页，其他页面相同）

``` html
<!-- views/index.html -->
{{template "header"}}
  <title>登录 - 我的 beego 博客</title>
</head>
<body>
  ...
{{template "footer"}}
```

打开 [localhost:8080/login](http://localhost:8080/login) 看到如下效果就ok啦！

![login](_images/login.png 'login')


## 登录功能实现

我们为了方便起见，就只配置一个管理员账号

在 `conf` 目录下的 `app.conf` 文件中 添加管理员账号、密码。

``` conf
uname = admin
pwd = admin
```

配置完成后，就可以开始写登录的 `POST` 请求逻辑了。

``` go
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
```

写好后，登录一下试试，但是我们还没有写登录后的效果，所以在UI上还看不出来。

打开开发者工具，在 `Application` 菜单下的 `Cookies` 中查看是否存在 `uname` `pwd` 

效果如下：

![cookie](_images/cookie.png 'cookie')

我们希望在登录成功后，管理员登录按钮变成退出按钮，让我们可以退出登录

在开始写退出登录的逻辑前，我们还需要说一下模板变量的使用

``` go
// controllers/login.go
...
func (this *LoginController) Get() {
	// 声明一个变量 a string 类型
	this.Data["a"] = "变量a"
	// 声明一个变量 IsLogin bool 类型
	this.Data["IsLogin"] = true
	// 声明一个长度为10的 slice 类型
	this.Data["nums"] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	this.TplName = "login.html"
}
...
``` 

``` html
<!-- views/login.html -->
...
<body>
  <!-- 使用变量 -->
  <p>{{.a}}</p>
  <p>{{.nums}}</p>

  <!-- 判断 -->
  <p>
    {{if .IsLogin}}
    已登录
    {{else}}
    未登录
    {{end}}
  </p>

  <!-- 循环 -->
  <p>
    {{range .nums}}
    <!-- `.`为current item -->
    {{.}}
    {{end}}
  </p>
  ...
``` 

效果如下

![template](_images/template.png 'template')

接下来我们实现下具体逻辑

``` go
// controllers/login.go

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
```

``` go
// controllers/home.go

...
func (this *HomeController) Get() {
  ...
  this.Data["IsLogin"] = checkAccount(this.Controller)
}

```

``` html
<!-- views/index.html -->

...
<body>
  ...
  <div class="navbar navbar-default navbar-fixed-top">
    <div class="container">
      ...
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
        ...
```

## 退出功能实现

``` go
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
  ...
}
```

登录与退出搞定~


<!-- ## 数据库操作

我们的登录页模板已经搞定，接下来就是创建、连接以及操作数据库。

在 `models` 目录下 创建 `models.go` 文件 -->

