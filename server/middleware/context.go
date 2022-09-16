package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

var enableTimeout = true

func DisableContextTimeout() {
	enableTimeout = false
}
func EnableContextTimeout() {
	enableTimeout = true
}

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		if enableTimeout {
			ctx, cancel := context.WithTimeout(c.Request.Context(), t)
			defer cancel()
			c.Request = c.Request.WithContext(ctx)
		}
		c.Next()
	}
}
