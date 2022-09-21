package router

import (
	"time"

	"github.com/UndertaIe/passwd/global"
	"github.com/UndertaIe/passwd/pkg/cache"
	"github.com/UndertaIe/passwd/server/middleware"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

const (
	ContextTimeOut = time.Second * 60
	UTCTime        = false
	TimeFormat     = time.RFC3339
	SentryRepanic  = false
	ServiceName    = "passwd"
)

func SetMiddlewares(r *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	switch gin.Mode() {
	case gin.DebugMode:
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	case gin.ReleaseMode: // export GIN_MODE=release
		r.Use(sentrygin.New(sentrygin.Options{Repanic: SentryRepanic})) // sentry deal with panic
		r.Use(cache.GinCache(global.Cacher))                            // cache context propagation
		r.Use(middleware.ContextTimeout(ContextTimeOut))
		r.Use(ginrus.Ginrus(global.Logger, TimeFormat, UTCTime))
		r.Use(otelgin.Middleware(ServiceName)) // 添加 tracing
	}
}
