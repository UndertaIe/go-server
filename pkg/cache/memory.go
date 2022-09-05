package cache

type MemoryInCache struct {
	Cache
}

func MemoryInConfig() map[string]interface{} {
	var m map[string]interface{}
	return m
}

func NewMemoryInCache(cc CacheConfig) (*MemoryInCache, error) {
	var err error
	return &MemoryInCache{}, err
}
