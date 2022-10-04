package service

import (
	"context"

	"github.com/UndertaIe/go-server/global"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Service struct {
	Ctx context.Context
	Db  *gorm.DB
}

func NewService(ctx context.Context) *Service {
	return &Service{
		Ctx: ctx,
		Db:  global.DBEngine.WithContext(ctx), // otel context propagation
	}
}

var log *logrus.Logger

func UseLog(l *logrus.Logger) {
	log = l
}
