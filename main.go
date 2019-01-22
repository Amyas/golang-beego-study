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
