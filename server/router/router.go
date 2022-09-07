package router

import (
	"github.com/UndertaIe/passwd/global"
	"github.com/UndertaIe/passwd/pkg/cache"
	"github.com/UndertaIe/passwd/server/middleware"
	v1 "github.com/UndertaIe/passwd/server/router/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	SetMiddlewares(r)
	apiv1 := r.Group("api/v1")
	{
		user := v1.User{}
		apiv1.GET("/user/:id", user.Get)
		apiv1.GET("/user", user.List)
		apiv1.POST("/user", user.Create)
		apiv1.DELETE("/user/:id", user.Delete)
		apiv1.PUT("/user/:id", user.Update)

		userSignup := v1.NewUserSignUp()
		apiv1.POST("/user/phone", userSignup.PhoneExists)
		apiv1.POST("/user/email", userSignup.EmailExists)
		apiv1.POST("/user/name", userSignup.UserNameExists)
	}
	{
		platform := v1.NewPlatform()
		apiv1.GET("/platform/:id", platform.Get)
		apiv1.GET("/platform", platform.List)
		apiv1.POST("/platform", platform.Create)
		apiv1.DELETE("/platform/:id", platform.Delete)
		apiv1.PUT("/platform/:id", platform.Update)
	}
	{
		userPasswd := v1.NewUserPasswd()
		apiv1.GET("/userpasswd", userPasswd.All)
		apiv1.GET("/userpasswd/:user_id", userPasswd.List)
		apiv1.GET("/userpasswd/:user_id/:platform_id", userPasswd.Get)
		apiv1.POST("/userpasswd", userPasswd.Create)
		apiv1.DELETE("/userpasswd/:user_id/:platform_id", userPasswd.Delete)
		apiv1.DELETE("/userpasswd/:user_id", userPasswd.DeleteList)
		apiv1.PUT("/userpasswd/:user_id/:platform_id", userPasswd.Update)
	}
	{
		sms := v1.NewSMS()
		apiv1.GET("/sms", sms.Get)
	}

	AuthRouters(r)
	CacheRouters(r)

	return r
}

func AuthRouters(r *gin.Engine) {
	{
		r.POST("/jwt/auth", Auth) // appkey,appsecret获取token
		admin := r.Group("/jwt/")
		admin.Use(middleware.JWT()) // 使用jwt鉴权如下接口
		admin.GET("/admin", PassAuth)
		admin.POST("/admin", PassAuth)
	}
	{
		uAuth := r.Group("/jwt/")
		r.POST("/user/auth", UserAuth)  // 使用user_id password获取token
		uAuth.Use(middleware.UserJwt()) // 使用jwt鉴权并在ctx中设置 user_id值
		uAuth.GET("/user/secret", PassUserAuth)
		uAuth.POST("/user/secret", PassUserAuth)
	}
}

// 集成redis，memory-in，memcached缓存中间件
func CacheRouters(r *gin.Engine) {
	{ // redis cache
		c := r.Group("/cache/api/")
		c.Use(cache.GinCache(global.Cacher)) // c.Keys["cache"] = global.Cacher
		// c.Use(cache.SiteCache())
		c.GET("/now", v1.Now)
		c.GET("/cnow", cache.CachePage(global.Cacher, cache.DEFAULT, v1.CacheNow))
		c.GET("/user/:id", cache.CachePage(global.Cacher, cache.DEFAULT, v1.GetUser))
		c.DELETE("/user/:id", v1.DeleteUser)
		c.PUT("/user/:id", v1.UpdateUser)
	}
	{ // memory-in cache
		c2 := r.Group("/memorycache/api/")
		c2.Use(cache.GinCache(global.MemInCacher))
		c2.GET("/now", v1.Now)
		c2.GET("/cnow", cache.CachePage(global.MemInCacher, cache.DEFAULT, v1.CacheNow))
		c2.GET("/user/:id", cache.CachePage(global.MemInCacher, cache.DEFAULT, v1.GetUser))
		c2.DELETE("/user/:id", v1.DeleteUser)
		c2.PUT("/user/:id", v1.UpdateUser)
	}
	{ // memcached cache
		c3 := r.Group("/memcached/api/")
		c3.Use(cache.GinCache(global.MemCacher))
		c3.GET("/now", v1.Now)
		c3.GET("/cnow", cache.CachePage(global.MemCacher, cache.DEFAULT, v1.CacheNow))
		c3.GET("/user/:id", cache.CachePage(global.MemCacher, cache.DEFAULT, v1.GetUser))
		c3.DELETE("/user/:id", v1.DeleteUser)
		c3.PUT("/user/:id", v1.UpdateUser)
	}
}
