package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/UndertaIe/passwd/global"
	"github.com/UndertaIe/passwd/server/router"
	"github.com/spf13/cobra"
)

func NewServer() *http.Server {
	handlers := router.NewRouter()
	s := &http.Server{
		Addr:           ":" + strconv.Itoa(global.ServerSettings.HttpPort),
		Handler:        handlers,
		ReadTimeout:    global.ServerSettings.ReadTimeout,
		WriteTimeout:   global.ServerSettings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}

func RunServe(cmd *cobra.Command, args []string) {
	s := NewServer()
	cfg := *global.ServerSettings

	go func() {
		var err error
		if cfg.EnabledTls {
			err = s.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile)
		} else {
			err = s.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()
	quit := make(chan os.Signal, 1) // 捕获SIGINT/SIGTERM信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("server running on [%v]; mode:[%v] \n", s.Addr, global.ServerSettings.RunMode)
	<-quit
	log.Println("Shuting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
