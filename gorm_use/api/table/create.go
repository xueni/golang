package table

import "gorm.io/gorm"

func CreateTable(db *gorm.DB, data interface{}) error {
	return db.Migrator().CreateTable(data)
}
