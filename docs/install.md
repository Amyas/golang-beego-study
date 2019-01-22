## 安装 beego

``` bash
go get github.com/astaxie/beego
```

## 安装 bee

``` bash
go get github.com/beego/bee
```

?> `bee` 工具是一个为了协助快速开发 beego 项目而创建的项目，通过 bee 您可以很容易的进行 beego 项目的创建、热编译、开发、测试、和部署。

## 创建项目

进入 `$GOPATH/src` 所在目录
``` bash
cd go/src/
```

创建 `blog` 项目
``` bash
bee new blog
```

启动项目
``` bash
cd blog/
bee run
```

打开 [localhost:8080](http://localhost:8080/) 看到 `Welcome to Beego` 就ok啦！

<!-- ?> 更多 `beego` 与 `bee` 用法请查阅官网 [beego.me](https://beego.me)。 -->