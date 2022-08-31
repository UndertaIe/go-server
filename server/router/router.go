package router

import (
	v1 "github.com/UndertaIe/passwd/server/router/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	setMiddlewares(r)
	apiv1 := r.Group("api/v1")
	{
		user := v1.User{}
		apiv1.GET("/user/:id", user.Get)
		apiv1.GET("/user", user.List)
		apiv1.POST("/user", user.Create)
		apiv1.DELETE("/user/:id", user.Delete)
		apiv1.PUT("/user/:id", user.Update)
	}
	{
		platform := v1.Platform{}
		apiv1.GET("/platform/:id", platform.Get)
		apiv1.GET("/platform", platform.List)
		apiv1.POST("/platform", platform.Create)
		apiv1.DELETE("/platform/:id", platform.Delete)
		apiv1.PUT("/platform/:id", platform.Update)
	}
	{
		userPasswd := v1.UserPasswd{}
		apiv1.GET("/userpasswd", userPasswd.All)
		apiv1.GET("/userpasswd/:user_id", userPasswd.List)
		apiv1.GET("/userpasswd/:user_id/:platform_id", userPasswd.Get)
		apiv1.POST("/userpasswd", userPasswd.Create)
		apiv1.DELETE("/userpasswd/:user_id/:platform_id", userPasswd.Delete)
		apiv1.DELETE("/userpasswd/:user_id", userPasswd.DeleteList)
		apiv1.PUT("/userpasswd/:user_id/:platform_id", userPasswd.Update)
	}
	{
		sms := v1.NewSms()
		apiv1.GET("/sms", sms.Get) // demo

	}

	return r
}

func setMiddlewares(r *gin.Engine) {
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())
}
