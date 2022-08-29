package model

import (
	"github.com/UndertaIe/passwd/database"
	"github.com/UndertaIe/passwd/pkg/page"
	"gorm.io/gorm"
)

type UserAccount struct {
	*database.BaseModel
	UserId     int    `json:"user_id"`
	PlatformId int    `json:"platform_id"`
	Password   string `json:"password"`
}

type UserAccountRow struct {
	UserId           int    `gorm:"column:user_id"`
	PlatformId       int    `gorm:"column:platform_id"`
	PlatformName     string `gorm:"column:name"`
	PlatformAbbr     string `gorm:"column:abbr"`
	Password         string `gorm:"column:password"`
	PlatformType     string `gorm:"column:type"`
	PlatformDesc     string `gorm:"column:description"`
	PlatformDomain   string `gorm:"column:domain"`
	PlatformImgUrl   string `gorm:"column:img_url"`
	PlatformLoginUrl string `gorm:"column:login_url"`
}

func (ua UserAccount) TableName() string {
	return "passwd_user_account"
}

func (ua UserAccount) GetAll(db *gorm.DB, pager *page.Pager) ([]UserAccountRow, error) {
	var resp []UserAccountRow
	db = db.Offset(pager.Offset()).Limit(pager.Limit())
	err := db.Model(&ua).Where(
		"passwd_user_account.is_deleted = ?", false).Joins(
		"join passwd_platform p on passwd_user_account.platform_id = p.platform_id").Select(
		"user_id", "p.platform_id", "name", "abbr", "password", "type", "description", "domain", "img_url", "login_url").Find(&resp).Error
	return resp, err
}

func (ua UserAccount) Get(db *gorm.DB) (UserAccountRow, error) {
	var resp UserAccountRow
	err := db.Model(&ua).Where(
		"user_id = ? and p.platform_id = ? and passwd_user_account.is_deleted = ?", ua.UserId, ua.PlatformId, false).Joins(
		"join passwd_platform p on passwd_user_account.platform_id = p.platform_id").Select(
		"user_id", "p.platform_id", "name", "abbr", "password", "type", "description", "domain", "img_url", "login_url").Find(&resp).Error
	return resp, err
}

func (ua UserAccount) Create(db *gorm.DB) (UserAccount, error) {
	err := db.Create(&ua).Error
	return ua, err
}

func (ua UserAccount) Update(db *gorm.DB, values interface{}) (UserAccount, error) {
	err := db.Model(&ua).Where("user_id = ? AND platform_id = ? AND is_deleted = ?", ua.UserId, ua.PlatformId, false).Updates(values).Error
	return ua, err
}

func (ua UserAccount) Delete(db *gorm.DB) error {
	err := db.Model(&ua).Where("user_id = ? AND platform_id = ? AND is_deleted = ?", ua.UserId, ua.PlatformId, false).Update("is_deleted", true).Error
	return err
}

func (ua UserAccount) DeleteList(db *gorm.DB) error {
	err := db.Model(&ua).Where("user_id = ? AND is_deleted = ?", ua.UserId, false).Update("is_deleted", true).Error
	return err
}

func (ua UserAccount) GetAccountsByUserID(db *gorm.DB, pager *page.Pager) ([]UserAccountRow, error) {
	var resp []UserAccountRow
	db = db.Offset(pager.Offset()).Limit(pager.Limit())
	err := db.Model(&ua).Where(
		"user_id = ? and passwd_user_account.is_deleted = ?", ua.UserId, false).Joins(
		"join passwd_platform p on passwd_user_account.platform_id = p.platform_id").Select(
		"user_id", "p.platform_id", "name", "abbr", "password", "type", "description", "domain", "img_url", "login_url").Find(&resp).Error
	return resp, err
}
