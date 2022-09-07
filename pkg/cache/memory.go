package cache

import (
	"reflect"
	"time"

	memCache "github.com/patrickmn/go-cache"
	"github.com/robfig/go-cache"
)

type MemoryInCache struct {
	mCache *memCache.Cache
	cfg    *MemoryInConfig
	Cache
}

// Get(key string, value interface{}) error
// Set(key string, value interface{}, expire time.Duration) error
// Add(key string, value interface{}, expire time.Duration) error
// Replace(key string, data interface{}, expire time.Duration) error
// Delete(key string) error
// Incr(key string, data int64) (int64, error)
// Flush() error

type MemoryInConfig struct {
	defaultExpireTime time.Duration
	cc                CacheConfig
}

func NewMemoryInCache(cc CacheConfig) (*MemoryInCache, error) {
	memInCache := &MemoryInCache{}
	err := memInCache.setup(cc)
	if err != nil {
		return nil, err
	}
	return memInCache, err
}

func defalutMemInConfig() *MemoryInConfig {
	return &MemoryInConfig{
		defaultExpireTime: 60 * time.Second,
		cc:                nil,
	}
}

func (mc *MemoryInCache) setup(cc CacheConfig) error {
	var err error
	if cc == nil {
		mc.cfg = defalutMemInConfig()
	} else {
		err = mc.setConfig(cc)
	}
	mc.mCache = memCache.New(mc.cfg.defaultExpireTime, time.Second)
	return err
}

func (mc *MemoryInCache) setConfig(cc CacheConfig) error {
	cfg := &MemoryInConfig{}
	m := cc()
	if ex, has := m["defaultExpireTime"]; has {
		if v, ok := ex.(int); ok {
			cfg.defaultExpireTime = time.Duration(v) * time.Second
		} else {
			return ExpireTimeTypeError
		}
	} else {
		return ExpireTimeNoKeyError
	}
	cfg.cc = cc
	mc.cfg = cfg
	return nil
}

func (mc *MemoryInCache) Get(key string, value interface{}) error {
	val, found := mc.mCache.Get(key)
	if !found {
		return NoKeyCacheError
	}

	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr && v.Elem().CanSet() {
		v.Elem().Set(reflect.ValueOf(val))
		return nil
	}
	return SetValueError
}

func (mc *MemoryInCache) Set(key string, value interface{}, expire time.Duration) error {
	mc.mCache.Set(key, value, expire)
	return nil
}

func (mc *MemoryInCache) Add(key string, value interface{}, expire time.Duration) error {
	err := mc.mCache.Add(key, value, expire)
	if err == cache.ErrKeyExists {
		return AddError.WithDetails(err.Error())
	}
	return err
}

func (mc *MemoryInCache) Replace(key string, value interface{}, expire time.Duration) error {
	err := mc.mCache.Replace(key, value, expire)
	if err != nil {
		return ReplaceError.WithDetails(err.Error())
	}
	return nil
}

func (mc *MemoryInCache) Delete(key string) error {
	mc.mCache.Delete(key)
	return nil
}

func (mc *MemoryInCache) Incr(key string, data int64) (int64, error) {
	val, err := mc.mCache.IncrementInt64(key, data)
	if err != nil {
		return -1, IncrError.WithDetails(err.Error())
	}
	return val, nil
}

func (mc *MemoryInCache) Flush() error {
	mc.mCache.Flush()
	return nil
}
