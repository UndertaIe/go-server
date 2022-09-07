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
	ExpireTimeTypeError       = NewCacheError(11005, "default must be int")
	ExpireTimeNoKeyError      = NewCacheError(11006, "map that cc returned must has defaultExpireTime")
	SetValueError             = NewCacheError(11007, "set value error")
	IncrError                 = NewCacheError(11008, "incr error")
	ReplaceError              = NewCacheError(11009, "replace error")
	AddError                  = NewCacheError(11010, "add error")
)
