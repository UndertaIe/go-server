package router

import "github.com/gin-gonic/gin"

func SetMiddlewares(r *gin.Engine) {
	switch gin.Mode() {
	case gin.DebugMode:
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	case gin.ReleaseMode:
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	}
}
