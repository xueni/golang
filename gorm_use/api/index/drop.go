package index

import "gorm.io/gorm"

func DropIndex(db *gorm.DB, name string, data interface{}) error {
	return db.Migrator().DropIndex(data, name)
}
