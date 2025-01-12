package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

var (
	config *Config
)

// 把全局对象保护起来
func C() *Config {
	if config == nil {
		panic("请加载配置信息，LoadConfigFromToml/LoadConfigFromEnv")
	}
	return config
}

// config 进行加载
// toml 文件加载
// 使用到的第三方库: "github.com/BurntSushi/toml"
func LoadConfigFromToml(path string) (*Config, error) {
	//所有默认配置
	conf := DefaultConfig()
	// toml配置,如果有就覆盖conf
	_, err := toml.DecodeFile(path, conf)
	if err != nil {
		return nil, err
	}
	// 赋值给全局变量
	config = conf
	return conf, nil
}

// 通过环境变量读取配置
// 采用第三方库: "github.com/caarlos0/env/v6"
// 读取环境变量  环境变量 ---> config object
func LoadConfigFromEnv() (*Config, error) {
	conf := DefaultConfig()
	// 通过环境变量给conf赋值，如果没有就是默认值
	if err := env.Parse(conf); err != nil {
		return nil, err
	}
	// 赋值给全局变量
	config = conf
	return nil, nil
}

// mysql ORM初始化连接池
