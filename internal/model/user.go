package model

import (
	"github.com/UndertaIe/go-server-env/database"
	"github.com/UndertaIe/go-server-env/pkg/app"
	"github.com/UndertaIe/go-server-env/pkg/utils"
	"gorm.io/gorm"
)

type User struct {
	UserId        int     `json:"user_id" gorm:"column:user_id"`
	UserName      string  `json:"user_name" gorm:"column:user_name"`
	Password      string  `json:"password" gorm:"column:password"`
	Salt          string  `json:"salt" gorm:"column:salt"`
	PhoneNumber   string  `json:"phone_number" gorm:"column:phone_number"`
	Email         *string `json:"email" gorm:"column:email; default:null"`
	ShareMode     int8    `json:"share_mode" gorm:"column:share_mode; default:0"`
	Role          int8    `json:"role" gorm:"column:role; default:0"`
	ProfileImgUrl string  `json:"profile_img_url" gorm:"column:profile_img_url; default:''"`
	Description   string  `json:"description" gorm:"column:description; default:''"`
	Sex           int8    `json:"sex" gorm:"column:sex; default:0"`
	*database.BaseModel
}

func (u User) TableName() string {
	return "passwd_user"
}

func (u User) Get(db *gorm.DB) (User, error) {
	err := db.Where("user_id = ? AND is_deleted = ?", u.UserId, false).Take(&u).Error
	return u, err
}

func (u User) GetUserByPhone(db *gorm.DB) (User, error) {
	err := db.Where("phone_number = ? AND is_deleted = ?", u.PhoneNumber, false).Take(&u).Error
	return u, err
}

func (u User) PhoneExists(db *gorm.DB) (bool, error) {
	err := db.Where("phone_number = ? AND is_deleted = ?", u.PhoneNumber, false).Take(&u).Error
	return utils.IsExistsRecord(err)
}

func (u User) NameExists(db *gorm.DB) (bool, error) {
	err := db.Where("user_name = ? AND is_deleted = ?", u.UserName, false).Take(&u).Error
	return utils.IsExistsRecord(err)
}

func (u User) EmailExists(db *gorm.DB) (bool, error) {
	err := db.Where("email = ? AND is_deleted = ?", u.Email, false).Take(&u).Error
	return utils.IsExistsRecord(err)
}

func (u User) GetUserByEmail(db *gorm.DB) (User, error) {
	err := db.Where("email = ? AND is_deleted = ?", u.Email, false).Take(&u).Error
	return u, err
}

func (u User) GetUserByName(db *gorm.DB) (User, error) {
	err := db.Where("user_name = ? AND is_deleted = ?", u.UserName, false).Take(&u).Error
	return u, err
}

func (u User) Create(db *gorm.DB) error {
	err := db.Omit("user_id").Create(&u).Error
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

func (u User) Count(db *gorm.DB) (int64, error) {
	var n int64
	err := db.Table(u.TableName()).Where("is_deleted = ?", false).Count(&n).Error
	return n, err
}

func (u User) GetUserList(db *gorm.DB, pager *app.Pager) ([]User, error) {
	db = pager.Use(db)
	var users []User
	tx := db.Table(u.TableName()).Where("is_deleted = ?", false).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}
