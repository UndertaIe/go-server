package router

import (
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/UndertaIe/go-server-env/global"
	"github.com/UndertaIe/go-server-env/pkg/cache"
	"github.com/UndertaIe/go-server-env/server/middleware"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

const (
	UTCTime       = false
	TimeFormat    = time.RFC3339
	SentryRepanic = false
)

// TODO: 打印middlewares
func SetMiddlewares(r *gin.Engine) {
	// gin.SetMode(gin.ReleaseMode)
	switch gin.Mode() {
	case gin.DebugMode:
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	case gin.ReleaseMode: // export GIN_MODE=release
		r.Use(ginrus.Ginrus(global.Logger, TimeFormat, UTCTime))
		r.Use(sentrygin.New(sentrygin.Options{Repanic: SentryRepanic})) // sentry deal with panic
		r.Use(middleware.ContextTimeout(global.APPSettings.DefaultContextTimeout))
		r.Use(otelgin.Middleware(global.ServiceName)) // 添加 tracing TODO: 增加 trace_id字段
		r.Use(cache.GinCache(global.Cacher))
	}
	PrintMiddlewares(r)

}

func PrintMiddlewares(r *gin.Engine) {
	global.Logger.Out.Write([]byte("Gin Using middlewares: \n"))
	for _, h := range r.Handlers {
		p := reflect.ValueOf(h).Pointer()
		rfunc := runtime.FuncForPC(p)
		fn := rfunc.Name()
		f, ln := rfunc.FileLine(p)
		sb := strings.Builder{}
		sb.WriteString("middleware: ")
		sb.WriteString(fn)
		sb.WriteString(" | ")
		sb.WriteString(f)
		sb.WriteString("/")
		sb.WriteString(cast.ToString(ln))
		global.Logger.Info(sb.String())
	}
}
