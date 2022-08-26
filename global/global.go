package global

import (
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

var DBEngine *gorm.DB
var Tracer opentracing.Tracer
