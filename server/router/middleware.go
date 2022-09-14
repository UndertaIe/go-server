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
)

func SetMiddlewares(r *gin.Engine) {
	// gin.SetMode(gin.ReleaseMode)
	switch gin.Mode() {
	case gin.DebugMode:
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	case gin.ReleaseMode:
		r.Use(middleware.ContextTimeout(ContextTimeOut))
		r.Use(ginrus.Ginrus(global.Logger, TimeFormat, UTCTime))
		r.Use(sentrygin.New(sentrygin.Options{Repanic: true})) // sentry接入报警
	}
}
