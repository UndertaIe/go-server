package main

import (
	"errors"
	"flag"
	"log"
	"strings"
	"time"

	"github.com/UndertaIe/passwd/cmd"
	"github.com/UndertaIe/passwd/config"
	"github.com/UndertaIe/passwd/database"
	"github.com/UndertaIe/passwd/global"
	"github.com/UndertaIe/passwd/pkg/cache"
	"github.com/UndertaIe/passwd/pkg/com/alibaba"
	"github.com/UndertaIe/passwd/pkg/sms"
	"github.com/UndertaIe/passwd/pkg/tracer"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	port       int
	runMode    string
	configPath string
)

func init() { // 初始化工作
	var err error

	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
	err = setupCacher()
	if err != nil {
		log.Fatalf("init.setupCacher err: %v", err)
	}
	err = setupMemoryInCacher()
	if err != nil {
		log.Fatalf("init.setupMemoryInCacher err: %v", err)
	}
	err = setupMemCacher()
	if err != nil {
		log.Fatalf("init.setupMemCacher err: %v", err)
	}
	err = setupSmsService()
	if err != nil {
		log.Fatalf("init.setupSmsService err: %v", err)
	}
	err = setupSentry()
	if err != nil {
		log.Fatalf("init.setupSentry err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

func main() {
	cmd.Run()
}

func setupSetting() error {

	flag.IntVar(&port, "port", 7788, "启动端口7788")
	flag.StringVar(&runMode, "mode", "debug", "启动模式(debug,prod)")
	flag.StringVar(&configPath, "config", "./", "配置文件路径,当前路径下")
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
	}
	hook := func() {
		global.APPSettings.DefaultContextTimeout *= time.Second
		global.ServerSettings.ReadTimeout *= time.Second
		global.ServerSettings.WriteTimeout *= time.Second
		global.JwtSettings.Expire *= time.Second
		global.SmsServiceSettings.DefaultExpireTime *= time.Second
	}
	s.ReadSections(sections, hook)
	return err
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = database.NewDBEngine(global.DatabaseSettings) // 在对全局变量赋值时不要使用 :=, 否则会导致左侧变量变为nil
	return err
}

func setupTracer() error {
	tracer, _, err := tracer.NewJaegerTracer("passwd-service", "127.0.0.1:6831")
	global.Tracer = tracer
	return err
}

func setupCacher() error {
	// cc := func() map[string]any {
	// 	return map[string]any{``
	// 		"host": global.RedisSettings.Host,
	// 		"db":   global.RedisSettings.Db,
	// 		// "password":          global.RedisSettings.Password,
	// 		"defaultExpireTime": global.RedisSettings.DefaultExpireTime,
	// 	}
	// }
	// cacher, err := cache.NewCache(cache.RedisT, cc)
	cacher, err := cache.NewCache(cache.RedisT, nil) //使用默认配置
	global.Cacher = cacher
	return err
}

func setupMemoryInCacher() error {
	// cc := func() map[string]any {
	// 	return map[string]any{
	// 		"defaultExpireTime": global.RedisSettings.DefaultExpireTime,
	// 	}
	// }
	// cacher, err := cache.NewCache(cache.RedisT, cc)
	cacher, err := cache.NewCache(cache.MemoryInT, nil) //使用默认配置
	global.MemInCacher = cacher
	return err
}

func setupMemCacher() error {
	// cc := func() map[string]any {
	// 	return map[string]any{``
	// 		"host": global.RedisSettings.Host,
	// 		"db":   global.RedisSettings.Db,
	// 		// "password":          global.RedisSettings.Password,
	// 		"defaultExpireTime": global.RedisSettings.DefaultExpireTime,
	// 	}
	// }
	// cacher, err := cache.NewCache(cache.RedisT, cc)
	cacher, err := cache.NewCache(cache.MemCacheT, nil) //使用默认配置
	global.MemCacher = cacher
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
	global.SmsService = srv
	return err
}

func setupSentry() error {
	mode := gin.Mode()
	isDebug := true
	if mode == gin.ReleaseMode {
		isDebug = false
	}
	if global.SentrySettings.Dsn == "" {
		return errors.New("sentry Dsn is nil")
	}
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   global.SentrySettings.Dsn,
		Debug: isDebug,
	})
	return err
}

func setupLogger() error {
	cfg := *global.APPSettings
	fileName := cfg.LogSavePath + "/" + cfg.LogFileName + cfg.LogFileExt
	global.Logger = logrus.New()
	if cfg.LogFormat == "json" {
		global.Logger.SetFormatter(&logrus.JSONFormatter{})
	}
	out := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    cfg.LogMaxSize, // megabytes
		MaxBackups: cfg.LogMaxBackup,
		MaxAge:     cfg.LogMaxAge,   //days
		Compress:   cfg.LogCompress, // disabled by default
		LocalTime:  cfg.LocalTime,
	}
	global.Logger.SetOutput(out)

	return nil
}
