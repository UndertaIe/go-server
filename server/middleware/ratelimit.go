package middleware

import (
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/UndertaIe/passwd/pkg/ratelimiter"
	"github.com/gin-gonic/gin"
)

func RateLimit(limiter ratelimiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := limiter.Key(c)
		bucket, has := limiter.GetBucket(key)
		if !has {
			opt := ratelimiter.BucketOption{
				Key: key,
			}
			limiter.AddBuckets(opt)
			bucket, _ = limiter.GetBucket(key)
		}

		if bucket.Take(1) {
			c.Next()
		} else {
			app.NewResponse(c).ToError(errcode.TooManyRequests)
			c.Abort()
		}
	}
}
