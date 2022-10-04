package middleware

import (
	"time"

	"github.com/UndertaIe/go-eden/cache"
	"github.com/UndertaIe/go-server/global"
	"github.com/gin-gonic/gin"
)

// wrap cache with global.logger
func CachePage(store cache.Cache, expire time.Duration, handle gin.HandlerFunc) gin.HandlerFunc {
	return cache.CachePage(store, expire, global.Logger, handle)
}

const (
	defaultExpire = time.Minute * 5
)

func DefaultCachePage(handle gin.HandlerFunc) gin.HandlerFunc {
	return cache.CachePage(global.Cacher, defaultExpire, global.Logger, handle)
}
