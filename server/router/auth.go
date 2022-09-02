package router

import (
	"github.com/UndertaIe/passwd/global"
	"github.com/UndertaIe/passwd/internal/service"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/auth"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// jwt鉴权:
// 使用app_key,app_secret从服务端获取token
func Auth(c *gin.Context) {
	resp := app.NewResponse(c)
	srv := service.NewService(c)
	param := service.AuthRequest{}
	err := c.Bind(&param)
	if err != nil {
		resp.ToError(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	err = srv.CheckAuth(&param)
	if err != nil {
		resp.ToError(errcode.UnauthorizedTokenError)
		return
	}
	token, err := auth.GenerateToken(param.AppKey, param.AppSecret, global.NewGlobal())
	if err != nil {
		resp.ToError(errcode.UnauthorizedTokenGenerate)
		return
	}
	resp.To(gin.H{"token": token})
}

func AuthPass(c *gin.Context) {
	resp := app.NewResponse(c)
	resp.To((gin.H{"msg": "通过鉴权"}))
}
