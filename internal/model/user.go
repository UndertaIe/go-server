package model

import (
	"github.com/UndertaIe/passwd/database"
	"github.com/UndertaIe/passwd/pkg/page"
	"gorm.io/gorm"
)

type User struct {
	*database.BaseModel
	UserId      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	ShareMode   int    `json:"share_mode"`
	Sex         int    `json:"sex"`
	Description string `json:"description"`
	Role        int    `json:"role"`
}

func (u User) TableName() string {
	return "passwd_user"
}

func (u User) Get(db *gorm.DB) (User, error) {
	err := db.Where("user_id = ? AND is_deleted = ?", u.UserId, false).Take(&u).Error
	return u, err
}
func (u User) Create(db *gorm.DB) error {
	err := db.Create(&u).Error
	return err
}

func (u User) Update(db *gorm.DB, values interface{}) error {
	err := db.Model(&u).Where("user_id = ? AND is_deleted = ?", u.UserId, false).Updates(values).Error
	return err
}

func (u User) Delete(db *gorm.DB) error {
	err := db.Model(&u).Where("user_id = ? AND is_deleted = ?", u.UserId, false).Update("is_deleted", true).Error
	return err
}

type UserRow struct {
	UserId      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	ShareMode   int    `json:"share_mode"`
	Sex         int    `json:"sex"`
	Description string `json:"description"`
	Role        int    `json:"role"`
}

func (u User) GetUserList(db *gorm.DB, pager *page.Pager) ([]UserRow, error) {
	db = db.Offset(pager.Offset()).Limit(pager.Limit())
	var users []UserRow
	tx := db.Model(&User{}).Where(map[string]interface{}{"is_deleted": false}).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}
