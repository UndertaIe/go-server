package internal

import (
	"context"

	"github.com/UndertaIe/passwd/global"
	"github.com/jinzhu/gorm"
)

type Service struct {
	ctx context.Context
	db  *gorm.DB
}

func NewService(ctx context.Context) *Service {
	return &Service{
		ctx: ctx,
		db:  global.DBEngine,
	}
}
