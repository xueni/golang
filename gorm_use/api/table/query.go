package table

import "gorm.io/gorm"

func HasTable(db *gorm.DB, name string) bool {
	return db.Migrator().HasTable(name)
}
