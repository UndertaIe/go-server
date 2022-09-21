package middleware

import (
	"time"

	"github.com/UndertaIe/passwd/global"
	"github.com/UndertaIe/passwd/pkg/cache"
	"github.com/gin-gonic/gin"
)

// wrap cache with global.logger
func CachePageWithTracing(store cache.Cache, expire time.Duration, handle gin.HandlerFunc) gin.HandlerFunc {
	return cache.CachePageWithTracing(store, expire, global.Logger, handle)
}
