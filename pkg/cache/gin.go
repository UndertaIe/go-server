package cache

import (
	"bytes"
	"crypto/sha1"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

var CACHE_MIDDLEWARE_KEY = "gin.cache"
var PageCachePrefix = "gin.page.cache"

type responseCache struct {
	Status int
	Header http.Header
	Data   []byte
}

type cachedWriter struct {
	gin.ResponseWriter
	status  int
	written bool
	store   Cache
	expire  time.Duration
	key     string
}

func urlEscape(prefix string, u string) string {
	key := url.QueryEscape(u)
	if len(key) > 200 {
		h := sha1.New()
		io.WriteString(h, u)
		key = string(h.Sum(nil))
	}
	var buffer bytes.Buffer
	buffer.WriteString(prefix)
	buffer.WriteString(":")
	buffer.WriteString(key)
	return buffer.String()
}

func newCachedWriter(store Cache, expire time.Duration, writer gin.ResponseWriter, key string) *cachedWriter {
	return &cachedWriter{writer, 0, false, store, expire, key}
}

func (w *cachedWriter) WriteHeader(code int) {
	w.status = code
	w.written = true
	w.ResponseWriter.WriteHeader(code)
}

func (w *cachedWriter) Status() int {
	return w.status
}

func (w *cachedWriter) Written() bool {
	return w.written
}

func (w *cachedWriter) Write(data []byte) (int, error) {
	ret, err := w.ResponseWriter.Write(data)
	if err == nil {
		//cache response
		store := w.store
		val := responseCache{
			w.status,
			w.Header(),
			data,
		}
		err = store.Set(w.key, val, w.expire)
		// if err != nil {
		// 	need logger
		// }
	}
	return ret, err
}

// Cache Middleware
func GinCache(store Cache, key ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var Key = CACHE_MIDDLEWARE_KEY
		if len(key) == 1 {
			Key = key[0]
		}
		c.Set(Key, store)
		c.Next()
	}
}

func SiteCache(store Cache, expire time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cache responseCache
		key := CacheKey(c)
		if err := store.Get(key, &cache); err != nil {
			c.Next()
		} else {
			c.Writer.WriteHeader(cache.Status)
			for k, vals := range cache.Header {
				for _, v := range vals {
					c.Writer.Header().Add(k, v)
				}
			}
			c.Writer.Write(cache.Data)
		}
	}
}

func CacheKey(c *gin.Context) string {
	key := urlEscape(PageCachePrefix, c.Request.URL.RequestURI())
	return key
}

func GetCache(c *gin.Context, cacheName ...string) Cache {
	var name = CACHE_MIDDLEWARE_KEY
	if len(cacheName) == 1 {
		name = cacheName[0]
	}
	return c.Keys[name].(Cache)
}

// for restful api
func DeleteCache(c *gin.Context, cacheName ...string) error {
	key := CacheKey(c)
	return GetCache(c, cacheName...).Delete(key)
}

// Cache Decorator
func CachePage(store Cache, expire time.Duration, handle gin.HandlerFunc) gin.HandlerFunc {

	return func(c *gin.Context) {
		var cache responseCache
		key := CacheKey(c)
		if err := store.Get(key, &cache); err != nil {
			// replace writer
			writer := newCachedWriter(store, expire, c.Writer, key)
			c.Writer = writer // 将responseWriter替换。write方法实现了写入client fd后写入缓存
			handle(c)
		} else {
			c.Writer.WriteHeader(cache.Status)
			for k, vals := range cache.Header {
				for _, v := range vals {
					c.Writer.Header().Add(k, v)
				}
			}
			c.Writer.Write(cache.Data)
		}
	}
}
