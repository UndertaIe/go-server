package ratelimiter

import (
	"context"

	"github.com/gin-gonic/gin"
)

func IP(c context.Context) string {
	switch c := c.(type) {
	case *gin.Context:
		return c.RemoteIP()
	default:
		panic("unsupported context")
	}
}

func Router(c context.Context) string {
	switch c := c.(type) {
	case *gin.Context:
		return c.FullPath()
	default:
		panic("unsupported context")
	}
}
