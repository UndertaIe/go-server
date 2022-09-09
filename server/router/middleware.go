package router

import (
	"time"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetMiddlewares(r *gin.Engine) {
	switch gin.Mode() {
	case gin.DebugMode:
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	case gin.ReleaseMode:
		r.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))
		r.Use(sentrygin.New(sentrygin.Options{}))
	}
}
