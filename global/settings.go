package global

import (
	"time"

	"github.com/UndertaIe/passwd/config"
)

var (
	DatabaseSettings *config.DatabaseSetting
	APPSettings      *config.AppSetting
	ServerSettings   *config.ServerSetting
	EmailSettings    *config.EmailSetting
	SmsSettings      *config.SmsSetting
	JwtSettings      *config.JWTSetting
)

type Globals struct{}

func NewGlobal() *Globals {
	return &Globals{}
}

func (g Globals) GetJWTSecret() []byte {
	return []byte(JwtSettings.Secret)
}

func (g Globals) GetJWTIssuer() string {
	return JwtSettings.Issuer
}

func (g Globals) GetJWTExpireTime() int64 {
	now := time.Now()
	return now.Add(JwtSettings.Expire).Unix()
}
