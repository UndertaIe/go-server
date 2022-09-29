package v1

import (
	"github.com/UndertaIe/passwd/internal/service"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
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

func AuthByUserEmailLink(c *gin.Context) {

}
