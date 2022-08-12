package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/UndertaIe/passwd/server/router"
)

// server: host,port,timeout,
// mysql: host,port,user,pwd,charset,db,
// sqlite：
func NewServer() *http.Server {
	HttpPort := 7788
	ReadTimeout := 60 * time.Second
	WriteTimeout := 60 * time.Second
	r := router.NewRouter()

	s := &http.Server{
		Addr:           ":" + strconv.Itoa(HttpPort),
		Handler:        r,
		ReadTimeout:    ReadTimeout,
		WriteTimeout:   WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}
