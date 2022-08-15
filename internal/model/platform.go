package model

import (
	"github.com/UndertaIe/passwd/database"
	"github.com/jinzhu/gorm"
)

type Platform struct {
	*database.BaseModel
	PlatformId       int    `json:"platform_id"`
	PlatformType     string `json:"platform_type"`
	PlatformName     string `json:"platform_name"`
	PlatformLoginUrl string `json:"platform_login_url"`
	PlatformDomain   string `json:"platform_domain"`
	PlatformDesc     string `json:"platform_desc"`
}

func (p Platform) TableName() string {
	return "passwd_platform"
}

func (p Platform) Get(db *gorm.DB) (Platform, error) {
	var pf Platform
	err := db.Where("platform_id = ? AND is_deleted = ?", pf.PlatformId, false).First(&pf).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return pf, err
	}
	return pf, nil
}

func (p Platform) Create(db *gorm.DB) (Platform, error) {
	err := db.Create(&p).Error
	return p, err
}

func (p Platform) Update(db *gorm.DB, values interface{}) error {
	err := db.Model(&p).Where("platform_id = ? AND is_del = ?", p.PlatformId, 0).Updates(values).Error
	return err
}

func (p Platform) Delete(db *gorm.DB) (*Platform, error) {
	err := db.Where("platform_id=? and is_deleted=?", p.PlatformId, false).Delete(&p).Error
	return &p, err
}
