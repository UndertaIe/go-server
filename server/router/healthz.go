package router

import (
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	resp := app.Response{Ctx: c}
	resp.Ok()
}
