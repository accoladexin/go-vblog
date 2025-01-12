package start

import (
	"fmt"
	"github.com/accoladexin/vblog/apps/blog/api"
	"github.com/accoladexin/vblog/apps/blog/impl"
	"github.com/accoladexin/vblog/common/logger"
	"github.com/accoladexin/vblog/conf"
	"github.com/accoladexin/vblog/protocol"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "vlog-main-start", // 命令名称 go run main.go  vlog-main-start  -h
	Short: "A simple main-start",
	Long:  "This is a simple vblog main-start CLI application built with Cobra",
	Run: func(cmd *cobra.Command, args []string) {
		// 具体指定逻辑
		fmt.Println("Welcome to main-start part")
		logger.L().Info().Str("main-start para", fmt.Sprintf("%s", args)).Msg("main-start")

		// TODO 1.程序的入口2.加载配置文件
		logger.L().Info().Str("main-start  configType：", fmt.Sprintf("%s", configType)).Msg("加载配置信息")
		switch configType {
		case "file":
			toml, err := conf.LoadConfigFromToml(configFile)
			cobra.CheckErr(err)
			logger.L().Info().Str("main-start file config", fmt.Sprintf("%s", toml)).Msg("main-start")
		case "env":
			env, err := conf.LoadConfigFromEnv()
			cobra.CheckErr(err)
			logger.L().Info().Str("main-start env  config", fmt.Sprintf("%s", env)).Msg("main-start")
		default:
			env, err := conf.LoadConfigFromEnv()
			cobra.CheckErr(err)
			logger.L().Info().Str("main-start default  config", fmt.Sprintf("%s", env)).Msg("main-start")
		}
		// todo 3.加载业务逻辑实现
		// 业务模块
		blogService := impl.NewImpl()
		err2 := blogService.Init()
		cobra.CheckErr(err2)
		apiHander := api.NewHandlerWithObj(blogService)

		// todo 4.启动服务
		r := gin.Default()
		// 注册路由
		apiHander.Registry(r)
		httpsever := protocol.NewHttp(r)
		err := httpsever.Start()
		cobra.CheckErr(err)

	},
}

// 执行命令
func Execute() error {
	if err := StartCmd.Execute(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

var (
	configType string
	configFile string
)

func init() {
	// 使用方式 go run main.go  vlog-main-start  --config-type -t env
	// Flag --config-type -t file -f etc/config.toml
	StartCmd.Flags().StringVarP(&configType, "config-type", "t", "file", "程序加载配置的方式")
	// &configType 指向的变量
	// "config-type" 参数 --config-type
	// "t" 参数 -t
	// "file" 默认值 -t file
	//"程序加载配置的方式" 说明
	StartCmd.Flags().StringVarP(&configFile, "config-file", "f", "etc/config.toml", "配置文件的路径")
}
