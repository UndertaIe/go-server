package global

import (
	"github.com/UndertaIe/passwd/pkg/cache"
	"github.com/UndertaIe/passwd/pkg/email"
	"github.com/UndertaIe/passwd/pkg/sms"
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
