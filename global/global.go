package global

import (
	"github.com/UndertaIe/passwd/pkg/cache"
	"github.com/UndertaIe/passwd/pkg/com/alibaba"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

var (
	DBEngine  *gorm.DB
	Tracer    opentracing.Tracer
	SmsClient *alibaba.Client
)

//: cache
var (
	Cacher      cache.Cache // redis
	MemInCacher cache.Cache // memory-in
	MemCacher   cache.Cache // memcached
)
