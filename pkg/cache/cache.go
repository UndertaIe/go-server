package cache

import (
	"time"
)

type cacheType string

const (
	MemoryInT cacheType = "memory"
	RedisT    cacheType = "redis"
	MemCacheT cacheType = "memcache"
)

var supportedCacheType = []string{string(MemoryInT), string(RedisT), string(MemCacheT)}

type Cache interface {
	Get(key string, value interface{}) error
	Set(key string, value interface{}, expire time.Duration) error
	Add(key string, value interface{}, expire time.Duration) error
	Replace(key string, data interface{}, expire time.Duration) error
	Delete(key string) error
	Incr(key string, data int64) (int64, error)
	Flush() error
}

type CacheConfig func() map[string]interface{}

// supported cache type: Redis，MemCache，MemoryIn
func NewCache(c cacheType, cc CacheConfig) (Cache, error) {
	switch c {
	case RedisT:
		return NewRedisCache(cc)
	case MemCacheT:
		return NewMemCache(cc)
	case MemoryInT:
		return NewMemoryInCache(cc)
	default:
		return nil, UnsupportedCacheTypeError
	}
}

// auto select cache type: Redis，MemCache，MemoryIn
func NewAutoCache(cc CacheConfig) (Cache, error) {
	var err error
	switch RedisT {
	case RedisT:
		c, err := NewRedisCache(cc)
		if err == nil {
			return c, nil
		}
		fallthrough
	case MemCacheT:
		c, err := NewMemCache(cc)
		if err == nil {
			return c, nil
		}
		fallthrough
	case MemoryInT:
		c, err := NewMemoryInCache(cc)
		if err == nil {
			return c, nil
		}
	}
	return nil, AutoInitCacheError.WithDetails(err.Error())
}
