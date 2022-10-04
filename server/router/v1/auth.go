package v1

import (
	"github.com/UndertaIe/go-server-env/internal/service"
	"github.com/UndertaIe/go-server-env/pkg/app"
	"github.com/UndertaIe/go-server-env/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// Auth godoc
// @Summary      用户登录/认证
// @Description  通过账号密码或验证码等方式登录
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user_id  		 body     int 	  false  "用户 ID"
// @Param        user_name  	 body     string  false  "用户名"
// @Param        phone_number  	 body     string  false  "手机号"
// @Param        email   		 body     string  false  "邮件"
// @Param        password  	 	 body     string  false  "用户密码"
// @Param        code   		 body     string  false  "验证码"
// @Success      200  {object}  string  		"成功"
// @Failure      400  {object}  errcode.Error "请求错误"
// @Failure      500  {object}  errcode.Error "内部错误"
// @Router       /api/v1/auth [post]
func (User) Auth(c *gin.Context) {
	resp := app.NewResponse(c)
	srv := service.NewService(c)
	param := service.AuthParam{}
	authType, err := service.BindAuth(c, &param)
	if err != nil {
		resp.ToError(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	token, nErr := srv.Auth(&param, authType)
	if err != nil {
		resp.ToError(nErr)
		return
	}

	resp.To(gin.H{"token": token})
}

// SendPhoneCode godoc
// @Summary      发送验证码
// @Description  用户请求验证码发送到手机，用于后续登录认证
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        phone_number	body	string  true  "手机号"
// @Success      200  {object}  string 		"成功"
// @Failure      500  {object}  errcode.Error "内部错误"
// @Router       /api/v1/auth/phone [post]
func (User) SendPhoneCode(c *gin.Context) {
	resp := app.NewResponse(c)
	srv := service.NewService(c)
	param := service.SendPhoneCodeParam{}
	err := srv.SendPhoneCode(param)
	if err != nil {
		resp.ToError(err)
		return
	}
	resp.Ok()
}

// SendEmailCode godoc
// @Summary      发送验证码
// @Description  用户请求验证码发送到邮箱，用于后续登录认证
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        email	body	string  true  "邮箱"
// @Success      200  {object}  string  		"成功"
// @Failure      500  {object}  errcode.Error "内部错误"
// @Router       /api/v1/auth/email [post]
func (User) SendEmailCode(c *gin.Context) {
	resp := app.NewResponse(c)
	srv := service.NewService(c)
	param := service.SendEmailCodeParam{}
	err := srv.SendEmailCode(param)
	if err != nil {
		resp.ToError(err)
		return
	}
	resp.Ok()
}

// SendEmailCode godoc
// @Summary      求验证链接
// @Description  用户请求验证链接发送到邮箱，用于后续登录认证
// @Tags         Auth
// @Produce      json
// @Param        email	body	string  true  "邮箱"
// @Success      200  {object} string 		"成功"
// @Failure      500  {object}  errcode.Error "内部错误"
// @Router       /api/v1/auth/link/:link [get]
func (User) SendEmailLink(c *gin.Context) {
	resp := app.NewResponse(c)
	srv := service.NewService(c)
	param := service.SendEmailLinkParam{}
	err := srv.SendEmailLink(param)
	if err != nil {
		resp.ToError(err)
		return
	}
	resp.Ok()
}
