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
