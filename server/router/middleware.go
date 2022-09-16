package router

import (
	"time"

	"github.com/UndertaIe/passwd/global"
	"github.com/UndertaIe/passwd/server/middleware"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
)

const (
	ContextTimeOut = time.Second * 60
	UTCTime        = false
	TimeFormat     = time.RFC3339
	SentryRepanic  = false
)

func SetMiddlewares(r *gin.Engine) {
	// gin.SetMode(gin.ReleaseMode)
	// export GIN_MODE=release
	switch gin.Mode() {
	case gin.DebugMode:
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	case gin.ReleaseMode:
		r.Use(sentrygin.New(sentrygin.Options{Repanic: SentryRepanic})) // sentry异常处理
		r.Use(middleware.ContextTimeout(ContextTimeOut))                // 超时处理
		r.Use(ginrus.Ginrus(global.Logger, TimeFormat, UTCTime))        // TODO: 没有做异常处理
	}
}
