package service

import (
	"context"

	"github.com/UndertaIe/passwd/global"
	"gorm.io/gorm"
)

type Service struct {
	Ctx context.Context
	Db  *gorm.DB
}

func NewService(ctx context.Context) *Service {
	return &Service{
		Ctx: ctx,
		Db:  global.DBEngine,
	}
}
