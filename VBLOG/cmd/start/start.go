package start

import (
	"fmt"
	_ "github.com/accoladexin/vblog/apps"
	"github.com/accoladexin/vblog/common/logger"
	"github.com/accoladexin/vblog/conf"
	"github.com/accoladexin/vblog/ioc"
	"github.com/accoladexin/vblog/protocol"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"strings"
)

var StartCmd = &cobra.Command{
	Use:   "vlog-main-start", // 命令名称 go run main.go  vlog-main-start  -h
	Short: "A simple main-start",
	Long:  "This is a simple vblog main-start CLI application built with Cobra",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
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
			logger.L().Info().Str("main-start default  config", fmt.Sprintf("%s", env)).Msg("main-start<>")
		}
		// todo 3.加载业务逻辑实现
		r := gin.Default()
		// 初始化ioc，所有ioc加载到 ioc容器中
		// import _ "github.com/accoladexin/vblog/apps" 已完成

		// 初始化ioc容器，对象自己的方法
		err = ioc.InitServiceIoc() // service初始化方法，初始化service自己的配置
		cobra.CheckErr(err)
		names := ioc.ShowServiceIoc()
		logger.L().Info().Str("Service IOC List", strings.Join(names, ", ")).Msg("main-start<>")
		// gin 挂载路由
		err = ioc.InitController("/vblog/api/v1", r) //1. controller 自己会绑定 service，2.里面有分组的逻辑 3，注册到r里面
		cobra.CheckErr(err)
		logger.L().Info().Str("Controller IOC List", strings.Join(ioc.ShowControllerIoc(), ", ")).Msg("main-start<>")
		//

		// 所有ioc执行自己的初化方法
		// 业务模块
		// v1 非ioc实现方法
		//blogService := impl.NewImpl()
		//err = blogService.Init() // 获得业务逻辑实现（service对象）
		//cobra.CheckErr(err)
		//apiHander := api.NewHandlerWithObj(blogService)
		// 注册路由
		// apiHander.Registry(r)

		// v2 ioc实现方法
		// 前面初始化完成了，所有这里不要任何操作

		// todo 4.启动服务
		httpsever := protocol.NewHttp(r) //
		err = httpsever.Start()
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
