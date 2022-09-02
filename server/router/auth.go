package router

import (
	"fmt"
	"strconv"

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

func PassAuth(c *gin.Context) {
	resp := app.NewResponse(c)
	resp.To((gin.H{"msg": "通过鉴权"}))
}

func UserAuth(c *gin.Context) {
	resp := app.NewResponse(c)
	srv := service.NewService(c)
	param := service.UserAuthRequest{}
	err := c.Bind(&param)
	if err != nil {
		resp.ToError(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	right, err := srv.UserAuth(&param)
	if err != nil { // 找不到user_id对应的记录
		resp.ToError(errcode.ErrorUserRecordNotFound.WithDetails(err.Error()))
		return
	}

	var token string
	if right {
		token, err = auth.GenerateUserToken(strconv.Itoa(param.UserId), global.NewGlobal())
		if err != nil {
			resp.ToError(errcode.UnauthorizedTokenGenerate.WithDetails(err.Error()))
			return
		}
	} else {
		resp.ToError(errcode.ErrorUserAuth)
		return
	}
	resp.To(gin.H{"token": token})
}

func PassUserAuth(c *gin.Context) {
	resp := app.NewResponse(c)
	msg := fmt.Sprintf("用户认证通过, 用户ID:%d", c.GetInt("user_id"))
	resp.To((gin.H{"msg": msg}))
}
