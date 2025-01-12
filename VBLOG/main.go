package main

import (
	"github.com/accoladexin/vblog/cmd"
	"github.com/spf13/cobra"
)

func main() {
	// 通过命令启动
	// go run main.go
	err := cmd.Execute()
	cobra.CheckErr(err)

}
