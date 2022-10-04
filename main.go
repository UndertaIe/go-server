package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/UndertaIe/go-server-env/cmd"
	"github.com/UndertaIe/go-server-env/config"
	"github.com/UndertaIe/go-server-env/database"
	"github.com/UndertaIe/go-server-env/global"
	"github.com/UndertaIe/go-server-env/pkg/app"
	"github.com/UndertaIe/go-server-env/pkg/cache"
	"github.com/UndertaIe/go-server-env/pkg/com/alibaba"
	"github.com/UndertaIe/go-server-env/pkg/email"
	"github.com/UndertaIe/go-server-env/pkg/sms"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	port       int
	runMode    string
	configPath string
)

func init() { // 初始化工作(有序初始化)
	var err error

	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupCacher()
	if err != nil {
		log.Fatalf("init.setupCacher err: %v", err)
	}
	err = setupSmsService()
	if err != nil {
		log.Fatalf("init.setupSmsService err: %v", err)
	}
	err = setupSentry()
	if err != nil {
		log.Fatalf("init.setupSentry err: %v", err)
	}
	setupEmailService()

	setupAppPagination()

}

// @title          passwd API
// @version        1.0
// @description    This is a passwd server, for saving platform password.
// @termsOfService http://swagger.io/terms/
// @tag.name 	   Go Eden
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
func main() {
	cmd.Run()
}

func setupSetting() error {

	flag.IntVar(&port, "port", 0, "默认启动端口0")
	flag.StringVar(&runMode, "mode", "", "启动模式(debug,prod)")
	flag.StringVar(&configPath, "config", "./", "配置文件路径:默认当前路径下所有yml文件")
	flag.Parse()

	s, err := config.NewSetting(strings.Split(configPath, ",")...)
	if err != nil {
		return err
	}
	sections := map[string]interface{}{
		"Server": &global.ServerSettings,
		"App":    &global.APPSettings,
		"MySQL":  &global.DatabaseSettings,
		// "SQLITE3":       &global.DatabaseSettings,
		"Email":         &global.EmailSettings,
		"SmsService":    &global.SmsServiceSettings,
		"JWT":           &global.JwtSettings,
		"Sentry":        &global.SentrySettings,
		"Redis":         &global.RedisSettings,
		"MemoryInCache": &global.MemoryInCacheSettings,
		"MemCache":      &global.MemCacheSettings,
		"Tracing":       &global.TracingSettings,
	}
	hooks := func() {
		global.APPSettings.DefaultContextTimeout *= time.Second
		global.ServerSettings.ReadTimeout *= time.Second // TODO: viper支持time.Duration默认为Seconds？ 不支持打个补丁
		global.ServerSettings.WriteTimeout *= time.Second
		global.JwtSettings.Expire *= time.Second
		global.SmsServiceSettings.DefaultExpireTime *= time.Second
		if port != 0 {
			global.ServerSettings.HttpPort = port
		}
		if runMode != "" {
			global.ServerSettings.RunMode = config.Mode(runMode)
		}
	}
	s.ReadSections(sections, hooks)
	return err
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = database.NewDBEngine(global.DatabaseSettings)
	return err
}

func setupTracer() error {
	var err error
	var exporter sdktrace.SpanExporter
	if config.IsDebug(global.ServerSettings.RunMode) {
		var f *os.File
		f, err = os.Create(global.APPSettings.TraceSavePath)
		if err != nil {
			return err
		}
		exporter, err = stdout.New(
			stdout.WithWriter(f),
			stdout.WithPrettyPrint(),
			stdout.WithoutTimestamps(),
		)
	} else {
		exporter, err = jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(global.TracingSettings.EndPoint)))
	}
	if err != nil {
		return err
	}
	r, err := resource.Merge(resource.Default(), resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(global.ServiceName),
		semconv.ServiceVersionKey.String(global.ServiceVersion),
	))
	if err != nil {
		return err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(r),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return err
}

func setupCacher() error {
	var err error
	cc := func() map[string]any {
		return map[string]any{
			"host":              global.RedisSettings.Host,
			"db":                global.RedisSettings.Db,
			"password":          global.RedisSettings.Password,
			"defaultExpireTime": global.RedisSettings.DefaultExpireTime,
		}
	}
	global.Cacher, err = cache.NewCache(cache.RedisT, cc)
	// cacher, err := cache.NewCache(cache.RedisT, nil) //使用默认配置
	if err != nil {
		return err
	}

	global.MemInCacher, err = cache.NewCache(cache.MemoryInT, nil) //使用默认配置
	if err != nil {
		return err
	}
	global.MemCacher, err = cache.NewCache(cache.MemCacheT, nil) //使用默认配置
	return err

}

func setupSmsService() error {
	ls := global.SmsServiceSettings
	cli, err := alibaba.NewClient(ls.AccessKey, ls.AccessSecret)
	if err != nil {
		return err
	}
	if global.Cacher == nil {
		log.Fatal("setup SmsService error: global.cacher is nil")
	}
	srv, err := sms.NewSmsCodeService(global.Cacher, cli, ls.DefaultExpireTime, ls.Prefix, ls.CodeLen)
	global.AuthCodeService = srv
	return err
}

func setupEmailService() {
	cfg := global.EmailSettings
	opt := email.Options{
		MailHost: cfg.Host,
		MailPort: cfg.Port,
		MailUser: cfg.UserName,
		MailPass: cfg.Password,
	}
	global.EmailClient = email.NewEmailClient(opt)
}

func setupSentry() error {
	isDebug := config.IsDebug(global.ServerSettings.RunMode)
	if global.SentrySettings.Dsn == "" {
		return errors.New("sentry Dsn is nil")
	}
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   global.SentrySettings.Dsn,
		Debug: isDebug,
	})
	sentry.Logger.SetOutput(global.Logger.Out)
	return err
}

func setupLogger() error {
	cfg := *global.APPSettings
	fileName := cfg.LogSavePath + "/" + cfg.LogFileName + cfg.LogFileExt
	global.Logger = logrus.New()
	if cfg.LogFormat == "json" {
		global.Logger.SetFormatter(&logrus.JSONFormatter{})
	}
	var out io.Writer
	if config.IsProduction(global.ServerSettings.RunMode) {
		out = &lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    cfg.LogMaxSize, // megabytes
			MaxBackups: cfg.LogMaxBackup,
			MaxAge:     cfg.LogMaxAge,   //days
			Compress:   cfg.LogCompress, // disabled by default
			LocalTime:  cfg.LocalTime,
		}
	} else { // 标准输出
		out = os.Stdout
	}
	global.Logger.SetOutput(out)

	return nil
}

func setupAppPagination() {
	app.SetPagerOption(app.PagerOption{
		DefaultPageSize:    global.APPSettings.DefaultPageSize,
		DefaultMaxPageSize: global.APPSettings.MaxPageSize,
	})
}
