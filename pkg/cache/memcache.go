package cache

type MemCache struct {
	Cache
}

func MemCacheConfig() map[string]interface{} {
	var m map[string]interface{}
	return m
}

func NewMemCache(cc CacheConfig) (*MemCache, error) {
	var err error
	return &MemCache{}, err
}
