package router

import (
	"github.com/UndertaIe/go-server-env/pkg/app"
	"github.com/gin-gonic/gin"
)

func healthz(c *gin.Context) {
	resp := app.NewResponse(c)
	resp.Ok()
}
