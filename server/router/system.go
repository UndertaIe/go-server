package router

import (
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/gin-gonic/gin"
)

func healthz(c *gin.Context) {
	resp := app.Response{Ctx: c}
	resp.Ok()
}
