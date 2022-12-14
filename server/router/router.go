package router

import (
	"github.com/UndertaIe/go-eden/auth"
	"github.com/UndertaIe/go-server/database"
	_ "github.com/UndertaIe/go-server/docs"
	"github.com/UndertaIe/go-server/global"
	"github.com/UndertaIe/go-server/internal/service"
	"github.com/UndertaIe/go-server/server/middleware"
	"github.com/UndertaIe/go-server/server/router/demo"
	v1 "github.com/UndertaIe/go-server/server/router/v1"
	v2 "github.com/UndertaIe/go-server/server/router/v2"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	UseGlobalLogger()

	r := gin.New()
	SetMiddlewares(r)
	cached := middleware.DefaultCachePageWithTracing //  add cache decorater by default cacher, expire, log
	dcache := middleware.DeleteCachePageWithTracing  // delete cache decorater by default cacher, expire, log

	// adminAuth := middleware.RoleAuth(auth.Admin) // 管理员权限
	// userAuth := middleware.RoleAuth(auth.User)   // 用户权限
	public := middleware.RoleAuth(auth.Public) // 公共权限

	adminAuth := middleware.RoleAuth(auth.Public) // 管理员权限
	userAuth := middleware.RoleAuth(auth.Public)  // 公共权限

	apiv1 := r.Group("api/v1")
	{
		user := v1.NewUser()
		apiv1.POST("/user", public, user.Create)
		apiv1.GET("/user/:id", userAuth, cached(user.Get))
		apiv1.GET("/user", adminAuth, cached(user.List))
		apiv1.PUT("/user/:id", userAuth, dcache(user.Update))
		apiv1.DELETE("/user/:id", userAuth, dcache(user.Delete))

		apiv1.POST("/user/auth", public, user.Auth)
		apiv1.POST("/user/auth/phone", public, user.SendPhoneCode)
		apiv1.POST("/user/auth/email", public, user.SendEmailCode)
		apiv1.POST("/user/auth/link/:link", public, user.SendEmailLink)
		apiv1.POST("/user/exists/phone", public, user.PhoneExists)
		apiv1.POST("/user/exists/email", public, user.EmailExists)
		apiv1.POST("/user/exists/username", public, user.UserNameExists)

	}
	{
		platform := v1.NewPlatform()
		apiv1.POST("/platform", adminAuth, platform.Create)
		apiv1.GET("/platform/:id", public, cached(platform.Get))
		apiv1.GET("/platform", public, cached(platform.List))
		apiv1.PUT("/platform/:id", adminAuth, dcache(platform.Update))
		apiv1.DELETE("/platform/:id", adminAuth, dcache(platform.Delete))
	}
	{
		userPasswd := v1.NewUserPasswd()
		apiv1.POST("/userpasswd", userAuth, userPasswd.Create)
		apiv1.GET("/userpasswd", userAuth, cached(userPasswd.All))
		apiv1.GET("/userpasswd/:user_id", userAuth, cached(userPasswd.List))
		apiv1.GET("/userpasswd/:user_id/:platform_id", userAuth, cached(userPasswd.Get))
		apiv1.PUT("/userpasswd/:user_id/:platform_id", userAuth, dcache(userPasswd.Update))
		apiv1.DELETE("/userpasswd/:user_id/:platform_id", userAuth, dcache(userPasswd.Delete))
		apiv1.DELETE("/userpasswd/:user_id", userAuth, dcache(userPasswd.DeleteList))
	}
	{ // system api
		r.GET("/healthz", healthz)
	}

	Demo(r)

	return r
}

// set global logger for handler, gin Writer, service
func UseGlobalLogger() {
	gin.DefaultWriter = global.Logger.Out
	gin.DefaultErrorWriter = global.Logger.Out

	v1.UseLog(global.Logger)
	v2.UseLog(global.Logger)
	demo.UseLog(global.Logger)

	service.UseLog(global.Logger)

	database.UselLogger(global.Logger)
}
