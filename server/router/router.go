package router

import (
	"time"

	"github.com/UndertaIe/passwd/global"
	"github.com/UndertaIe/passwd/pkg/cache"
	"github.com/UndertaIe/passwd/pkg/ratelimiter"
	"github.com/UndertaIe/passwd/server/middleware"
	"github.com/UndertaIe/passwd/server/router/demo"
	v1 "github.com/UndertaIe/passwd/server/router/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	SetMiddlewares(r)
	apiv1 := r.Group("api/v1")
	{
		user := v1.NewUser()
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

	SmsRouters(r)
	AuthRouters(r)
	CacheRouters(r)
	SentryRouters(r)
	OtelRouters(r)
	RateLimitRouters(r)

	return r
}

func RateLimitRouters(r *gin.Engine) {
	igroup := r.Group("/limit/ip/")
	igroup.Use(
		middleware.RateLimit(
			ratelimiter.NewRateLimiter(
				ratelimiter.WithIPKey(),
				// ratelimiter.WithPool(),
				ratelimiter.WithRefreshInterval(time.Minute),
			),
		),
	)
	{ // 同一个ip访问r1,r2,r3接口使用一个Group的ip限流
		igroup.GET("/r1", demo.RateLimit)
		igroup.GET("/r2", demo.RateLimit)
	}

	igroup2 := r.Group("/limit/router/")
	igroup2.Use(
		middleware.RateLimit(
			ratelimiter.NewRateLimiter(
				ratelimiter.WithRouterKey(),
				ratelimiter.WithPool(),
				ratelimiter.WithRefreshInterval(time.Minute),
			),
		),
	)
	{ // 同一个ip访问r1,r2,r3接口使用一个Group的ip限流
		igroup2.GET("/r1", demo.RateLimit)
		igroup2.GET("/r2", demo.RateLimit)
		user := v1.NewUser()
		igroup2.GET("/user/:id", user.Get)
		igroup2.GET("/user", user.List)
	}
}

func OtelRouters(r *gin.Engine) {
	c := r.Group("/otel/api/")
	c.GET("/now", demo.Now)
	c.GET("/cnow", middleware.CachePageWithTracing(global.Cacher, cache.DEFAULT, demo.CacheNow))
	c.GET("/user/:id", middleware.CachePageWithTracing(global.Cacher, cache.DEFAULT, demo.GetUser))
	c.DELETE("/user/:id", demo.DeleteUserWithTracing)
	c.PUT("/user/:id", demo.DeleteUserWithTracing)
}

func SentryRouters(r *gin.Engine) {
	r.GET("/sentry", demo.Sentry)
}

func SmsRouters(r *gin.Engine) {
	sms := demo.NewSMS()
	r.GET("/sms/code/:phone", sms.PhoneCode)
	r.POST("/sms/auth", sms.PhoneAuth)
}

func AuthRouters(r *gin.Engine) {
	{
		r.POST("/jwt/auth", demo.Auth) // appkey,appsecret获取token
		admin := r.Group("/jwt/")
		admin.Use(middleware.JWT()) // 使用jwt鉴权如下接口
		admin.GET("/admin", demo.PassAuth)
		admin.POST("/admin", demo.PassAuth)
	}
	{
		uAuth := r.Group("/jwt/")
		r.POST("/user/auth", demo.UserAuth) // 使用user_id password获取token
		uAuth.Use(middleware.UserJwt())     // 使用jwt鉴权并在ctx中设置 user_id值
		uAuth.GET("/user/secret", demo.PassUserAuth)
		uAuth.POST("/user/secret", demo.PassUserAuth)
	}
}

// 集成redis，memory-in，memcached缓存中间件
func CacheRouters(r *gin.Engine) {
	{ // redis cache
		c := r.Group("/cache/api/")
		c.Use(cache.GinCache(global.Cacher)) // c.Keys["cache"] = global.Cacher
		// c.Use(cache.SiteCache())
		c.GET("/now", demo.Now)
		c.GET("/cnow", middleware.CachePage(global.Cacher, cache.DEFAULT, demo.CacheNow))
		c.GET("/user/:id", middleware.CachePage(global.Cacher, cache.DEFAULT, demo.GetUser))
		c.DELETE("/user/:id", demo.DeleteUser)
		c.PUT("/user/:id", demo.UpdateUser)
	}
	{ // memory-in cache
		c2 := r.Group("/memorycache/api/")
		c2.Use(cache.GinCache(global.MemInCacher))
		c2.GET("/now", demo.Now)
		c2.GET("/cnow", middleware.CachePage(global.MemInCacher, cache.DEFAULT, demo.CacheNow))
		c2.GET("/user/:id", middleware.CachePage(global.MemInCacher, cache.DEFAULT, demo.GetUser))
		c2.DELETE("/user/:id", demo.DeleteUser)
		c2.PUT("/user/:id", demo.UpdateUser)
	}
	{ // memcached cache
		c3 := r.Group("/memcached/api/")
		c3.Use(cache.GinCache(global.MemCacher))
		c3.GET("/now", demo.Now)
		c3.GET("/cnow", middleware.CachePage(global.MemCacher, cache.DEFAULT, demo.CacheNow))
		c3.GET("/user/:id", middleware.CachePage(global.MemCacher, cache.DEFAULT, demo.GetUser))
		c3.DELETE("/user/:id", demo.DeleteUser)
		c3.PUT("/user/:id", demo.UpdateUser)
	}
}
