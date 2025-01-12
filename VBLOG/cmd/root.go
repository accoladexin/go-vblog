package cmd

import (
	"fmt"
	"github.com/accoladexin/vblog/cmd/start"
	"github.com/accoladexin/vblog/common/logger"
	"github.com/spf13/cobra"
)

// 1. 创建一个根命令
var rootCmd = &cobra.Command{
	Use:   "vlog-main",
	Short: "A simple CLI vblog-application",
	Long:  "This is a simple vblog CLI application built with Cobra",
	Run: func(cmd *cobra.Command, args []string) {
		// 具体指定逻辑
		fmt.Println("vlog main part")
		logger.L().Info().Str("main para", fmt.Sprintf("%s", args)).Msg("main")
	},
}

// 执行命令
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// 注册子的命令

func init() {
	rootCmd.AddCommand(start.StartCmd)
}
