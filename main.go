package main

import (
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
)

var (
	port       int
	runMode    string
	configPath string
)

func init() { // 初始化工作
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}
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

}

func main() {
	cmd.Run()
}

func setupFlag() error {
	flag.IntVar(&port, "port", 7788, "启动端口7788")
	flag.StringVar(&runMode, "mode", "debug", "启动模式(debug,prod)")
	flag.StringVar(&configPath, "config", "./", "配置文件路径,当前路径下")
	flag.Parse()

	return nil
}

func setupSetting() error {
	s, err := config.NewSetting(strings.Split(configPath, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSettings)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.APPSettings)
	if err != nil {
		return err
	}
	err = s.ReadSection("MySQL", &global.DatabaseSettings)
	if err != nil {
		return err
	}
	err = s.ReadSection("SQLITE3", &global.DatabaseSettings)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &global.EmailSettings)
	if err != nil {
		return err
	}
	err = s.ReadSection("SmsService", &global.SmsServiceSettings)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global.JwtSettings)
	if err != nil {
		return err
	}
	err = s.ReadSection("Redis", &global.RedisSettings)
	if err != nil {
		return err
	}
	err = s.ReadSection("MemoryInCache", &global.MemoryInCacheSettings)
	if err != nil {
		return err
	}

	err = s.ReadSection("MemCache", &global.MemCacheSettings)
	if err != nil {
		return err
	}

	// 将时间单位ns转化为s
	global.APPSettings.DefaultContextTimeout *= time.Second
	global.ServerSettings.ReadTimeout *= time.Second
	global.ServerSettings.WriteTimeout *= time.Second
	global.JwtSettings.Expire *= time.Second
	global.SmsServiceSettings.DefaultExpireTime *= time.Second

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
	srv, err := sms.NewSmsCodeService(global.Cacher, cli, ls.DefaultExpireTime, ls.Prefix, ls.CodeLen)
	global.SmsService = srv
	return err
}

func setupSentry() error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://8a5adff2f48d407da9992ce08c1254ec@o1401849.ingest.sentry.io/6733237", // TODO: 需要初始化
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	return err
}
