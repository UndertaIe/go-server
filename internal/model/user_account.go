package model

import (
	"github.com/UndertaIe/passwd/database"
	"github.com/jinzhu/gorm"
)

type UserAccount struct {
	*database.BaseModel
	UserId     int    `json:"user_id"`
	PlatformId int    `json:"platform_id"`
	Password   string `json:"password"`
}

func (ua UserAccount) TableName() string {
	return "passwd_user_account"
}

func (ua UserAccount) Get(db gorm.DB) (UserAccount, error) {
	err := db.Where("user_id = ? AND platformId = ? AND is_deleted = ?", ua.UserId, ua.PlatformId, false).First(&ua).Error
	return ua, err
}

func (ua UserAccount) Create(db gorm.DB) (UserAccount, error) {
	err := db.Create(&ua).Error
	return ua, err
}

func (ua UserAccount) Update(db gorm.DB, values interface{}) error {
	err := db.Model(&ua).Where("user_id = ? AND platformId = ? AND is_deleted = ?", ua.UserId, ua.PlatformId, false).Updates(values).Error
	return err
}

func (ua UserAccount) Delete(db gorm.DB) error {
	err := db.Where("user_id = ? AND platformId = ? AND is_deleted = ?", ua.UserId, ua.PlatformId, false).Delete(&ua).Error
	return err
}

