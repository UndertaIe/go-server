package ratelimiter

import (
	"context"

	"github.com/gin-gonic/gin"
)

type KeyFunc func(c context.Context) string

func (f KeyFunc) apply(lo LimiterOption) LimiterOption {
	lo.keyFunc = f
	return lo
}

func WithIPKey() KeyFunc {
	f := func(c context.Context) string {
		switch c := c.(type) {
		case *gin.Context:
			return c.RemoteIP()
		default:
			panic("unsupported context")
		}
	}
	return KeyFunc(f)
}

func WithRouterKey() KeyFunc {
	f := func(c context.Context) string {
		switch c := c.(type) {
		case *gin.Context:
			return c.FullPath() // "/limit/router/user/:id"
		default:
			panic("unsupported context")
		}
	}
	return KeyFunc(f)
}
