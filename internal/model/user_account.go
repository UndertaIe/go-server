package model

import (
	"github.com/UndertaIe/internal/db"
)

type UserAccount struct {
	*db.BaseModel
	userId     int    `json:"user_id"`
	platformId int    `json:"platform"`
	password   string `json:"password"`
}

func (u User) TableName() string {
	return "passwd_user"
}
