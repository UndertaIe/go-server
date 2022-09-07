package cache

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisCache struct {
	pool *redis.Pool
	cfg  *RedisConfig
	Cache
}

func NewRedisCache(cc CacheConfig) (*RedisCache, error) {
	var err error
	rc := &RedisCache{}
	rc.setup(cc)
	return rc, err
}

type RedisConfig struct {
	db                int
	host              string
	password          string
	defaultExpireTime time.Duration
	cc                CacheConfig
}

// redis默认配置信息
func defaultRedisConfig() *RedisConfig {
	return &RedisConfig{
		db:                0,
		host:              "127.0.0.1:6379",
		password:          "",
		defaultExpireTime: defaultExpire * time.Second,
		cc:                nil,
	}
}

// https://github.com/gin-gonic/contrib/blob/master/cache/redis.go

func (rc *RedisCache) setConfig(cc CacheConfig) {
	m := cc()
	cfg := &RedisConfig{}
	if db, ok := m["db"]; ok {
		cfg.db = db.(int)
	}
	cfg.host = m["host"].(string)
	if pwd, ok := m["password"]; ok {
		cfg.password = pwd.(string)
	}
	if ex, ok := m["defaultExpireTime"]; ok {
		cfg.defaultExpireTime = time.Second * time.Duration(ex.(int))
	} else {
		cfg.defaultExpireTime = time.Second * defaultExpire
	}
	cfg.cc = cc
	rc.cfg = cfg
}

func (rc *RedisCache) setup(cc CacheConfig) {
	if cc != nil {
		rc.setConfig(cc)
	} else {
		rc.cfg = defaultRedisConfig()
	}
	rc.pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   512,
		Wait:        false,
		IdleTimeout: defaultIdle * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", rc.cfg.host)
			if err != nil {
				return nil, err
			}
			if len(rc.cfg.password) > 0 {
				if _, err := c.Do("AUTH", rc.cfg.password); err != nil {
					c.Close()
					return nil, err
				}
			} else {
				// check with PING
				if _, err := c.Do("PING"); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", rc.cfg.db); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if _, err := c.Do("PING"); err != nil {
				return err
			}
			return nil
		},
	}
}

func (rc *RedisCache) Conn() redis.Conn {
	return rc.pool.Get()
}

func (rc *RedisCache) Get(key string, value interface{}) error {
	conn := rc.Conn()
	defer conn.Close()
	raw, err := conn.Do("GET", key)
	if raw == nil {
		return NoKeyCacheError
	}
	item, err := redis.Bytes(raw, err)
	if err != nil {
		return err
	}
	return deserialize(item, value)
}

func (rc *RedisCache) Set(key string, value interface{}, expire time.Duration) error {
	conn := rc.Conn()
	defer conn.Close()
	return rc.set(conn.Do, key, value, expire)
}

func (rc *RedisCache) Add(key string, value interface{}, expire time.Duration) error {
	conn := rc.Conn()
	defer conn.Close()
	if rc.exists(conn.Do, key) {
		return KeyExistsError
	}
	return rc.set(conn.Do, key, value, expire)
}

func (rc *RedisCache) Replace(key string, value interface{}, expire time.Duration) error {
	conn := rc.Conn()
	defer conn.Close()
	if !rc.exists(conn.Do, key) {
		return NoKeyCacheError
	}
	err := rc.set(conn.Do, key, value, expire)
	return err
}

func (rc *RedisCache) Delete(key string) error {
	conn := rc.Conn()
	defer conn.Close()
	if !rc.exists(conn.Do, key) {
		return NoKeyCacheError
	}
	_, err := conn.Do("DEL", key)
	return err
}

func (rc *RedisCache) Incr(key string, delta int64) (int64, error) {
	conn := rc.Conn()
	defer conn.Close()
	val, err := conn.Do("GET", key)
	if err != nil {
		return 0, UnKnownCacheError.WithDetails(err.Error())
	}
	if val == nil {
		return 0, NoKeyCacheError
	}
	curVal, err := redis.Int64(val, nil)
	if err != nil {
		return 0, err
	}
	sum := curVal + delta
	_, err = conn.Do("SET", key, sum)
	if err != nil {
		return 0, err
	}
	return sum, nil
}

func (rc *RedisCache) Flush() error {
	conn := rc.Conn()
	defer conn.Close()
	_, err := conn.Do("FLUSHALL")
	return err
}

func (rc *RedisCache) exists(do func(string, ...interface{}) (interface{}, error), key string) bool {
	has, _ := redis.Bool(do("EXISTS", key))
	return has
}

func (rc *RedisCache) set(do func(string, ...interface{}) (interface{}, error), key string, value interface{}, expires time.Duration) error {
	if expires == DEFAULT {
		expires = rc.cfg.defaultExpireTime
	}

	b, err := serialize(value)
	if err != nil {
		return err
	}
	if expires > 0 {
		_, err := do("SETEX", key, int32(expires/time.Second), b)
		return err
	} else {
		_, err := do("SET", key, b)
		return err
	}
}
