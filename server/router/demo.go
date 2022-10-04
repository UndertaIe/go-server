package router

import (
	"time"

	"github.com/UndertaIe/go-eden/auth"
	"github.com/UndertaIe/go-eden/cache"
	"github.com/UndertaIe/go-eden/ratelimiter"
	"github.com/UndertaIe/go-eden/swagger"
	"github.com/UndertaIe/go-server/global"
	"github.com/UndertaIe/go-server/server/middleware"
	"github.com/UndertaIe/go-server/server/router/demo"
	v1 "github.com/UndertaIe/go-server/server/router/v1"
	"github.com/gin-gonic/gin"
)

// ======================================
// 				module demo
// ======================================
func Demo(r gin.IRouter) {
	r = r.Group("/demo/")
	SmsRouter(r)
	AuthRouter(r)
	CacheRouter(r)
	SentryRouter(r)
	OtelRouter(r)
	RateLimitRouter(r)
	SwaggerRouter(r)
	TimeoutRouter(r)
}

func TimeoutRouter(r gin.IRouter) {
	r.GET("/context/timeout", demo.ContextTimeout)
	r.GET("/context/notimeout", demo.ContextNoTimeout)
}

func SwaggerRouter(r gin.IRouter) {
	r.GET("/swagger/*any", swagger.HandlerFunc)
}

func RateLimitRouter(r gin.IRouter) {
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

func OtelRouter(r gin.IRouter) {
	c := r.Group("/otel/api/")
	c.GET("/now", demo.Now)
	c.GET("/cnow", middleware.CachePageWithTracing(global.Cacher, cache.DEFAULT, demo.CacheNow))
	c.GET("/user/:id", middleware.CachePageWithTracing(global.Cacher, cache.DEFAULT, demo.GetUser))
	c.DELETE("/user/:id", demo.DeleteUserWithTracing)
	c.PUT("/user/:id", demo.DeleteUserWithTracing)
}

func SentryRouter(r gin.IRouter) {
	r.GET("/sentry", demo.Sentry)
}

func SmsRouter(r gin.IRouter) {
	sms := demo.NewSMS()
	r.GET("/sms/code/:phone", sms.PhoneCode)
	r.POST("/sms/auth", sms.PhoneAuth)
}

func AuthRouter(r gin.IRouter) {
	{
		r.POST("/jwt/auth", demo.Auth) // appkey,appsecret获取token
		admin := r.Group("/jwt/")
		admin.Use(middleware.JWT()) // 使用jwt鉴权如下接口
		admin.GET("/admin", demo.PassAuth)
		admin.POST("/admin", demo.PassAuth)
	}
	{
		uAuth := r.Group("/jwt/")
		r.POST("/user/auth", demo.UserAuth)         // 使用user_id password获取token
		uAuth.Use(middleware.RoleAuth(auth.Public)) // 使用jwt鉴权并在ctx中设置 user_id值
		uAuth.GET("/user/secret", demo.PassUserAuth)
		uAuth.POST("/user/secret", demo.PassUserAuth)
	}
}

// 集成redis，memory-in，memcached缓存中间件
func CacheRouter(r gin.IRouter) {
	{ // redis cache
		c := r.Group("/cache/api/")
		// c.Use(cache.GinCache(global.Cacher)) // c.Keys["cache"] = global.Cacher
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
