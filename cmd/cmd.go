package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/UndertaIe/passwd/local"
	"github.com/UndertaIe/passwd/server"
	"github.com/spf13/cobra"
)

func Run() {
	var runServeCmd = &cobra.Command{
		Use:   "serve",
		Short: "",
		Long:  "提供全面的账号密码添加方式",
		Run:   runServe,
	}
	var runLocalCmd = &cobra.Command{
		Use:   "local",
		Short: "",
		Long:  "提供简单的账号密码添加方式",
		Run:   runLocal,
	}
	runServeCmd.AddCommand(runLocalCmd)
	runServeCmd.Execute()

}

func runServe(cmd *cobra.Command, args []string) {
	s := server.NewServer()
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1) // 捕获SIGINT/SIGTERM信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	log.Println("server running...")
	<-quit
	log.Println("Shuting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}

func runLocal(cmd *cobra.Command, args []string) {
	local.RunLocal()
}
