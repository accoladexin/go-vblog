package impl_test

import (
	"github.com/accoladexin/vblog/apps/blog"
	"github.com/accoladexin/vblog/apps/blog/impl"
	"github.com/accoladexin/vblog/conf"
	"os"
)

// 初始化需要被测试对象

var (
	// 被测试的对象需要从Ioc中获取

	controller blog.Service
)

func init() {

	// 从环境变量中获取配置,单测一般不用配置文件
	os.Setenv("MYSQL_MAX_OPEN_CONN", "12")
	os.Setenv("MYSQL_DB", "12")
	_, err2 := conf.LoadConfigFromEnv()
	if err2 != nil {
		panic(err2)
	}
	// TODO 初始化imp对象
	obj := impl.NewImpl()
	err := obj.Init()
	if err != nil {
		panic(err)
	}
	controller = obj

}
