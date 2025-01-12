package conf_test

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/accoladexin/vblog/conf"
	"testing"
)

func TestConfig(t *testing.T) {
	// TODO
	// 测试配置文件 创建一个空的 conf.Config
	confObj := &conf.Config{}
	// 将配置文件解析到 confObj
	file, err := toml.DecodeFile("test/config.toml", confObj) //test/config.toml 表示当前文件夹下的test/config.toml
	if err != nil {
		t.Error(err)
	}
	fmt.Println("content:==》", file)
	fmt.Println(confObj)

}

func TestConfig2(t *testing.T) {
	// TODO
	// 不是并发，先不用考虑锁的问题
	_, err := conf.LoadConfigFromToml("test/config.toml")
	if err != nil {
		panic(err)

	}
	fmt.Println("全局变量==>\n", conf.C())
	//fmt.Println("Test_config", fromToml)

}

// 测试数据连接池

func TestGorm(t *testing.T) {
	var err error
	conf, err := conf.LoadConfigFromToml("test/config.toml")
	if err != nil {
		panic(err)
	}
	fmt.Println(conf.MySQL.DB)

}
