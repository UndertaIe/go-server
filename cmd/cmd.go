package cmd

import (
	"github.com/UndertaIe/passwd/local"
	"github.com/UndertaIe/passwd/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "个人密码管理工具",
	Long:  "个人密码管理工具",
}

func Run() {
	var runServeCmd = &cobra.Command{
		Use:   "serve",
		Short: "服务模式",
		Long:  "",
		Run:   server.RunServe,
	}
	var runLocalCmd = &cobra.Command{
		Use:   "local",
		Short: "命令行模式",
		Long:  "",
		Run:   local.RunLocal,
	}
	rootCmd.AddCommand(runServeCmd)
	rootCmd.AddCommand(runLocalCmd)

	rootCmd.Execute()
}
