package model

import (
	"github.com/UndertaIe/passwd/database"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/utils"
	"gorm.io/gorm"
)

type Platform struct {
	*database.BaseModel
	PlatformId       int    `gorm:"column:platform_id"`
	PlatformName     string `gorm:"column:name"`
	PlatformAbbr     string `gorm:"column:abbr"`
	PlatformType     string `gorm:"column:type; default:''"`
	PlatformDesc     string `gorm:"column:description; default:''"`
	PlatformDomain   string `gorm:"column:domain; default:''"`
	PlatformImgUrl   string `gorm:"column:img_url; default:''"`
	PlatformLoginUrl string `gorm:"column:login_url; default:''"`
}

func (p Platform) TableName() string {
	return "passwd_platform"
}

func (p Platform) Get(db *gorm.DB) (Platform, error) {
	var pf Platform
	err := db.Where("platform_id = ? AND is_deleted = ?", p.PlatformId, false).Take(&pf).Error
	return pf, err
}

func (p Platform) GetList(db *gorm.DB, pager *app.Pager) ([]Platform, error) {
	var platforms []Platform
	db = pager.Use(db)
	err := db.Table(p.TableName()).Where("is_deleted = ?", false).Find(&platforms).Error
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
	err := db.Model(&p).Where("platform_id = ? AND is_deleted = ?", p.PlatformId, 0).Updates(values).Error
	return err
}

func (p Platform) Delete(db *gorm.DB) (*Platform, error) {
	err := db.Model(&p).Where("platform_id = ? AND is_deleted = ?", p.PlatformId, false).Update("is_deleted", true).Error
	return &p, err
}

func (p Platform) IsExistsName(db *gorm.DB) (bool, error) {
	err := db.Where("name = ?", p.PlatformName).Take(&p).Error
	return utils.IsExistsRecord(err)
}

func (p Platform) IsExistsAbbr(db *gorm.DB) (bool, error) {
	err := db.Where("abbr = ?", p.PlatformAbbr).Take(&p).Error
	return utils.IsExistsRecord(err)
}

func (p Platform) Count(db *gorm.DB) (int, error) {
	var rows int64
	err := db.Table(p.TableName()).Where("is_deleted = ?", false).Count(&rows).Error
	return int(rows), err
}
