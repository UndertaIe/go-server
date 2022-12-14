package database

import (
	"fmt"
	"time"

	"github.com/UndertaIe/go-server/config"
	"github.com/UndertaIe/go-server/global"
	"github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"
)

type BaseModel struct {
	CreatedAt  string `json:"created_at" gorm:"column:created_at"`
	ModifiedAt string `json:"modified_at" gorm:"column:modified_at"`
	IsDeleted  bool   `json:"is_deleted" gorm:"column:is_deleted"`
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
	if config.IsDebug(global.ServerSettings.RunMode) {
		db = db.Debug()
	}
	if global.TracingSettings.Enabled {
		db.Use(tracing.NewPlugin())
	}
	db.Callback().Create().Replace("gorm:before_create", createCallback) //替换gorm的全局钩子函数
	db.Callback().Update().Replace("gorm:before_update", updateCallback)
	return db, nil
}

func createCallback(db *gorm.DB) {
	now := time.Now().Format("2006-01-02 15:04:05")
	db.Statement.SetColumn("created_at", now)
	db.Statement.SetColumn("modified_at", now)
}

func updateCallback(db *gorm.DB) {
	now := time.Now().Format("2006-01-02 15:04:05")
	db.Statement.SetColumn("modified_at", now)
}

func UselLogger(log *logrus.Logger) {
	logger.Default = logger.New(global.Logger, logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})

	logger.Discard = logger.New(global.Logger, logger.Config{})
}
