package database

import (
	"fmt"

	"github.com/UndertaIe/passwd/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt  string `json:"created_at"`
	ModifiedAt string `json:"modified_at"`
	DeletedAt  string `json:"deleted_at"`
	IsDeleted  bool   `json:"is_deleted"`
}

func NewDBEngine(dbSetting *config.DatabaseSetting) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", dbSetting.UserName, dbSetting.DBType, dbSetting.Host, dbSetting.DBName, dbSetting.Charset, dbSetting.ParseTime)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db.Debug(), nil
}
