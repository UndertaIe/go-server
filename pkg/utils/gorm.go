package utils

import (
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func IsExistsRecord(err error) (bool, error) {
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func IsDupEntryError(err error) bool {
	e, ok := err.(*mysql.MySQLError)
	return ok && e.Number == 1062
}
