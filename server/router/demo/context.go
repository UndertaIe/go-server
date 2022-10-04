package demo

import (
	"errors"
	"net/http"
	"time"

	"github.com/UndertaIe/go-server-env/global"
	"github.com/UndertaIe/go-server-env/pkg/app"
	"github.com/UndertaIe/go-server-env/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func ContextTimeout(c *gin.Context) {
	time.Sleep(global.APPSettings.DefaultContextTimeout + time.Second)
	resp := app.NewResponse(c)
	if err := AssertTimeoutError(c); err != nil {
		resp.ToError(errcode.RequestTimeout)
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "The request did not time out"})
}

func ContextNoTimeout(c *gin.Context) {
	time.Sleep(global.APPSettings.DefaultContextTimeout)
	resp := app.NewResponse(c)
	if err := AssertTimeoutError(c); err != nil {
		resp.ToError(errcode.RequestTimeout)
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "The request did not time out"})
}

func AssertTimeoutError(c *gin.Context) error {
	select {
	case <-c.Request.Context().Done():
		return errors.New("context timeout error")
	default:
	}
	return nil
}
