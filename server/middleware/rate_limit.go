package middleware

import (
	"github.com/UndertaIe/passwd/pkg/ratelimiter"
	"github.com/gin-gonic/gin"
)

func RateLimit(limiter ratelimiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
