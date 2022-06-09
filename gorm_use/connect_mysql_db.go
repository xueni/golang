package gorm_use

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:111111@tcp(127.0.0.1:3306)/test?charset=utf8mb3&parseTime=True&loc=Local"

func ConnectDB() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func ConnectConfigDB() (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{
		DriverName:                    "",
		ServerVersion:                 "",
		DSN:                           dsn,
		Conn:                          nil,
		SkipInitializeWithVersion:     false,
		DefaultStringSize:             0,
		DefaultDatetimePrecision:      nil,
		DisableDatetimePrecision:      false,
		DontSupportRenameIndex:        false,
		DontSupportRenameColumn:       false,
		DontSupportForShareClause:     false,
		DontSupportNullAsDefaultValue: false,
	}))
}
