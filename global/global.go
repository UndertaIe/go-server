package global

import (
	"github.com/UndertaIe/go-eden/cache"
	"github.com/UndertaIe/go-eden/email"
	"github.com/UndertaIe/go-eden/sms"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	ServiceName    = "passwd"
	ServiceVersion = "v1.0.0"
)

var (
	DBEngine *gorm.DB
	Tracer   opentracing.Tracer
	Logger   *logrus.Logger
)

var (
	AuthCodeService *sms.SmsService
	EmailClient     email.Client
)

//: cache
var (
	Cacher      cache.Cache // redis
	MemInCacher cache.Cache // memory-in
	MemCacher   cache.Cache // memcached
)
