package model

import (
	"github.com/UndertaIe/internal/db"
)

type Platform struct {
	*db.BaseModel
	platformId       int    `json:"platform_id"`
	platformType     string `json:"platform_type"`
	platformName     string `json:"platform_name"`
	platformLoginUrl string `json:"platform_login_url"`
	platformDomain   string `json:"platform_domain"`
	platformDesc     string `json:"platform_desc"`
}
