package cache

import (
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type MemCache struct {
	cache *memcache.Client
	cfg   *MemCacheConfig
	Cache
}

type MemCacheConfig struct {
	hosts             []string
	defaultExpireTime time.Duration
	cc                CacheConfig
}

func NewMemCache(cc CacheConfig) (*MemCache, error) {
	mc := &MemCache{}
	err := mc.setup(cc)
	if err != nil {
		return nil, err
	}
	return mc, nil
}

func defalutMemCacheConfig() *MemCacheConfig {
	return &MemCacheConfig{
		hosts:             []string{"127.0.0.1:11211"},
		defaultExpireTime: 60 * time.Second,
		cc:                nil,
	}
}

func (mc *MemCache) setup(cc CacheConfig) error {
	var err error
	if cc == nil {
		mc.cfg = defalutMemCacheConfig()
	} else {
		err = mc.setConfig(cc)
	}
	mc.cache = memcache.New(mc.cfg.hosts...)
	return err
}

func (mc *MemCache) setConfig(cc CacheConfig) error {
	cfg := &MemCacheConfig{}
	m := cc()
	if hosts, has := m["hosts"]; has {
		if v, ok := hosts.([]string); ok {
			cfg.hosts = v
		} else {
			return HostsTypeError
		}
	} else {
		return HostsNoNilError
	}
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

func (mc *MemCache) Set(key string, value interface{}, expires time.Duration) error {
	return mc.set((*memcache.Client).Set, key, value, expires)
}

func (mc *MemCache) Add(key string, value interface{}, expires time.Duration) error {
	return mc.set((*memcache.Client).Add, key, value, expires)
}

func (mc *MemCache) Replace(key string, value interface{}, expires time.Duration) error {
	return mc.set((*memcache.Client).Replace, key, value, expires)
}

func (mc *MemCache) Get(key string, value interface{}) error {
	item, err := mc.cache.Get(key)
	if err != nil {
		return convertMemcacheError(err)
	}
	return deserialize(item.Value, value)
}

func (mc *MemCache) Delete(key string) error {
	return convertMemcacheError(mc.cache.Delete(key))
}

func (mc *MemCache) Increment(key string, delta uint64) (uint64, error) {
	newValue, err := mc.cache.Increment(key, delta)
	return newValue, convertMemcacheError(err)
}

func (mc *MemCache) Decrement(key string, delta uint64) (uint64, error) {
	newValue, err := mc.cache.Decrement(key, delta)
	return newValue, convertMemcacheError(err)
}

func (mc *MemCache) Flush() error {
	return NotSupportError
}

func (mc *MemCache) set(do func(*memcache.Client, *memcache.Item) error,
	key string, value interface{}, expire time.Duration) error {
	switch expire {
	case DEFAULT:
		expire = mc.cfg.defaultExpireTime
	case FOREVER:
		expire = time.Duration(0)
	}

	b, err := serialize(value)
	if err != nil {
		return err
	}
	return convertMemcacheError(do(mc.cache, &memcache.Item{
		Key:        key,
		Value:      b,
		Expiration: int32(expire / time.Second),
	}))
}

func convertMemcacheError(err error) error {
	switch err {
	case nil:
		return nil
	case memcache.ErrCacheMiss:
		return NoKeyCacheError
	case memcache.ErrNotStored:
		return SetValueError
	}
	return err
}
