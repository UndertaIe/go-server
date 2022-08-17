package cmd

import (
	"github.com/UndertaIe/passwd/local"
	"github.com/UndertaIe/passwd/server"
	"github.com/spf13/cobra"
)

func Run() {
	var runServeCmd = &cobra.Command{
		Use:   "serve",
		Short: "",
		Long:  "",
		Run:   server.RunServe,
	}
	var runLocalCmd = &cobra.Command{
		Use:   "local",
		Short: "",
		Long:  "",
		Run:   local.RunLocal,
	}
	runServeCmd.AddCommand(runLocalCmd)
	runServeCmd.Execute()

}
