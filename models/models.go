package models

import (
	"os"
	"path"
	"strconv"
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

// 创建文章
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

// 删除文章
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

// 获取分类列表
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

// 获取文章列表
func GetTopicList() ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	_, err := qs.All(&topics)
	return topics, err
}
