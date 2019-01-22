我们接下来就要开始实现分类以及文章模块啦。在开始前呢，我们先关联、创建数据库。方便我们后续操作。

## 模型创建

在 `models` 目录下创建 `models.go` 文件
``` go
package models

import (
	"os"
	"path"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	// 定义数据库路径已经名称
	_DB_NAME = "data/blog.db"
	// 定义所使用的数据库
	_SQLITE3_DRIVER = "sqlite3"
)

// 定义分类表
type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

// 定义文章表
type Topic struct {
	Id               int64
	Uid              int64
	Title            string
	Content          string `orm:"size(5000)"`
	Attachment       string
	Created          time.Time `orm:"index"`
	Updated          time.Time `orm:"index"`
	Views            int64     `orm:"index"`
	Author           string
	ReplyTime        time.Time `orm:"index"`
	ReplyCount       int64
	RepleyLastUserId int64
}

func RegisterDB() {
	// 如果数据库目录不存在
	_, err := os.Stat(_DB_NAME)
	IsExist := err == nil || os.IsExist(err)
	if !IsExist {
		// 创建目录
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		// 创建文件
		os.Create(_DB_NAME)
	}

	// 注册定义的model
	orm.RegisterModel(new(Category), new(Topic))
	// 注册所使用的数据库驱动
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRMySQL)
	// 注册默认数据库
	// orm必须注册一个别名为 `default` 的数据库，作为默认使用
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}
```

## 数据库注册

在 `main.go` 文件中添加注册、创建、启动数据库操作

``` go
package main

import (
	"blog/models"
	_ "blog/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {
	// 开启 `orm` 调试模式
	orm.Debug = true

	// 启动自动建表
	// 参数1：表名
	// 参数2：是否每次启动都重新创建表，我们只需要第一次创建就好
	// 参数3：是否打印建表信息
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
```

## sqlite可视化

[DB Browser for SQLite 官网](https://sqlitebrowser.org/)