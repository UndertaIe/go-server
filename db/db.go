package global

import "github.com/jinzhu/gorm"

var (
	support_databases := [...]string{"sqlite3", "mysql", "mongodb"}
	default_database := support_databases[0]
	DBEngine *gorm.DB
)
