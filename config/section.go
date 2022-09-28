package config

import (
	"time"
)

type mode string

const (
	Debug      mode = "debug"
	Production mode = "prod"
)

func Mode(m string) mode {
	switch mode(m) {
	case Debug:
		return Debug
	case Production:
		return Production
	default:
		return Debug
	}
}

func IsDebug(m mode) bool {
	return m == Debug
}

func IsProduction(m mode) bool {
	return m == Production
}

type ServerSetting struct {
	RunMode      mode // debug or prod; default: debug
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	EnabledTls   bool
	CertFile     string
	KeyFile      string
}

type AppSetting struct {
	DefaultPageSize       int
	MaxPageSize           int
	DefaultContextTimeout time.Duration
	TraceSavePath         string
	LocalTime             bool
	LogSavePath           string
	LogFileName           string
	LogFileExt            string
	LogFormat             string
	LogMaxSize            int
	LogMaxBackup          int
	LogMaxAge             int
	LogCompress           bool
	UploadSavePath        string
	UploadServerUrl       string
	UploadImageMaxSize    int
	UploadImageAllowExts  []string
}

type EmailSetting struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

type SmsServiceSetting struct {
	Origin            string
	AccessKey         string
	AccessSecret      string
	DefaultExpireTime time.Duration
	Prefix            string
	CodeLen           int
}

type SentrySetting struct {
	Dsn string
}

type JWTSetting struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type DatabaseSetting struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type RedisSetting struct {
	Host              string
	Db                int
	Password          string
	DefaultExpireTime int
}

type MemoryInCacheSetting struct {
	DefaultExpireTime int
}

type MemCacheSetting struct {
	Hosts             []string
	DefaultExpireTime int
}

type TracingSetting struct {
	Enabled  bool
	EndPoint string
}

var sections = make(map[string]interface{})

type Hook func()

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReadSections(m map[string]interface{}, hooks ...Hook) error {
	var err error
	for k, p := range m {
		err = s.ReadSection(k, p)
		if err != nil {
			return err
		}
	}
	for _, hook := range hooks {
		hook()
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
