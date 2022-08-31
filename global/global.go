package global

import (
	"github.com/UndertaIe/passwd/pkg/com/alibaba"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

var DBEngine *gorm.DB
var Tracer opentracing.Tracer
var SmsClient *alibaba.Client
