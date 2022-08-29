package model

import (
	"fmt"

	"github.com/UndertaIe/passwd/database"
	"github.com/UndertaIe/passwd/pkg/page"
	"gorm.io/gorm"
)

type Platform struct {
	*database.BaseModel
	PlatformId       int    `gorm:"column:platform_id"`
	PlatformName     string `gorm:"column:name"`
	PlatformAbbr     string `gorm:"column:abbr"`
	PlatformType     string `gorm:"column:type"`
	PlatformDesc     string `gorm:"column:description"`
	PlatformDomain   string `gorm:"column:domain"`
	PlatformImgUrl   string `gorm:"column:img_url"`
	PlatformLoginUrl string `gorm:"column:login_url"`
}

func (p Platform) TableName() string {
	return "passwd_platform"
}

func (p Platform) Get(db *gorm.DB) (Platform, error) {
	var pf Platform
	err := db.Where("platform_id = ? AND is_deleted = ?", p.PlatformId, false).Take(&pf).Error
	fmt.Println(err)
	return pf, err
}

func (p Platform) GetList(db *gorm.DB, pager *page.Pager) ([]Platform, error) {
	var platforms []Platform
	db = db.Offset(pager.Offset()).Limit(pager.Limit())
	err := db.Find(&platforms).Error
	if err != nil {
		return nil, err
	}
	return platforms, nil
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
	err := db.Model(&p).Where("platform_id = ? AND is_deleted = ?", p.PlatformId, false).Update("is_deleted", true).Error
	return &p, err
}
