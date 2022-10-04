package cache

import (
	"bytes"
	"crypto/sha1"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/UndertaIe/go-server-env/pkg/logger"
	"github.com/gin-gonic/gin"
)

const (
	CACHE_MIDDLEWARE_KEY = "gin.cache"
	PageCachePrefix      = "gin.page.cache"
	cacheFlag            = "cache.flag"
)

type responseCache struct {
	Status int
	Header http.Header
	Data   []byte
}

type cacheWritePool struct {
	pool *sync.Pool
}

func newCacheWritePool() *cacheWritePool {
	p := &sync.Pool{New: func() any {
		return newCachedWriter()
	}}
	return &cacheWritePool{p}
}

func (p *cacheWritePool) Get() *cachedWriter {
	c, _ := p.pool.Get().(*cachedWriter)
	return c
}

func (p *cacheWritePool) Put(cw *cachedWriter) {
	p.pool.Put(cw)
}

var writerPool = newCacheWritePool()

type cachedWriter struct {
	gin.ResponseWriter
	ctx     *gin.Context
	status  int
	written bool
	store   Cache
	expire  time.Duration
	key     string
	log     logger.Log
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

func newCachedWriter() *cachedWriter {
	return &cachedWriter{}
}

func (w *cachedWriter) WriteHeader(code int) {
	w.status = code
	w.written = true
	w.ResponseWriter.WriteHeader(code)
}

func (w *cachedWriter) Status() int {
	return w.status
}

func (w *cachedWriter) Reset(ctx *gin.Context, store Cache, expire time.Duration, writer gin.ResponseWriter, key string, log logger.Log) {
	w.status = 0
	w.written = false

	w.ctx = ctx
	w.store = store
	w.expire = expire
	w.ResponseWriter = writer
	w.key = key
	w.log = log
}

func (w *cachedWriter) Written() bool {
	return w.written
}

func (w *cachedWriter) cacheEnabled() bool {
	if v, ok := w.ctx.Get(cacheFlag); ok {
		switch v := v.(type) {
		case CacheFlag:
			if IsDisableCache(v) {
				return false
			}
		}
	}
	return true
}

func (w *cachedWriter) Write(data []byte) (int, error) {
	ret, err := w.ResponseWriter.Write(data)
	if err == nil && w.cacheEnabled() {
		//cache response
		store := w.store
		val := responseCache{
			w.status,
			w.Header(),
			data,
		}
		err = store.Set(w.key, val, w.expire)
		if err != nil {
			w.log.Error("set cache error")
		}
	}
	return ret, err
}

type CacheFlag int8

const (
	enableCache CacheFlag = iota
	disableCache
)

func IsDisableCache(f CacheFlag) bool {
	return f == disableCache
}

func IsEnableCache(f CacheFlag) bool {
	return f == enableCache
}

func EnableCache(c *gin.Context) {
	c.Set(cacheFlag, enableCache)
}
func DisableCache(c *gin.Context) {
	c.Set(cacheFlag, disableCache)
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

func SiteCache(store Cache) gin.HandlerFunc {
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
	if _, ok := c.Keys[name]; ok {
		return c.Keys[name].(Cache)
	}
	return nil
}

// for call in gin.HandlerFunc
func DeleteCache(c *gin.Context, cacheName ...string) error {
	key := CacheKey(c)
	cacheImpl := GetCache(c, cacheName...)
	if cacheImpl == nil {
		return nil
	}
	return cacheImpl.Delete(key)
}

// for DeleteCache Decorator
func DeleteCachePage(handle gin.HandlerFunc, cacheName ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		handle(c)
		DeleteCache(c, cacheName...)
	}
}

// Cache Decorator
func CachePage(store Cache, expire time.Duration, log logger.Log, handle gin.HandlerFunc) gin.HandlerFunc {

	return func(c *gin.Context) {
		var cache responseCache
		key := CacheKey(c)
		if err := store.Get(key, &cache); err != nil {
			writer := writerPool.Get()
			writer.Reset(c, store, expire, c.Writer, key, log)
			c.Writer = writer // 将responseWriter替换。write方法实现了写入client fd后写入缓存
			EnableCache(c)
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
