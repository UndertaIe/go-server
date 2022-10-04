package middleware

import (
	"time"

	"github.com/UndertaIe/go-eden/cache"
	"github.com/UndertaIe/go-server/global"
	"github.com/gin-gonic/gin"
)

// wrap cache with global.logger
func CachePageWithTracing(store cache.Cache, expire time.Duration, handle gin.HandlerFunc) gin.HandlerFunc {
	return cache.CachePageWithTracing(store, expire, global.Logger, handle)
}

func DefaultCachePageWithTracing(handle gin.HandlerFunc) gin.HandlerFunc {
	return cache.CachePageWithTracing(global.Cacher, defaultExpire, global.Logger, handle)
}

var DeleteCachePageWithTracing = cache.DeleteCachePageWithTracing
