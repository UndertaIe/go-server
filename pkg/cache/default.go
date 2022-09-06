package cache

import (
	"strings"
	"time"
)

const (
	defaultExpire = 600
	defaultIdle   = 3

	DEFAULT = time.Duration(0)
	FOREVER = time.Duration(-1)
)

// errors:

var (
	UnKnownCacheError         = NewCacheError(11001, "unknown cache error")
	UnsupportedCacheTypeError = NewCacheError(11002, "unsupported cache type,supported cache type:) "+strings.Join(supportedCacheType, ","))
	AutoInitCacheError        = NewCacheError(11003, "auto init cache error")
	NoKeyCacheError           = NewCacheError(11004, "no key cache error")
)
