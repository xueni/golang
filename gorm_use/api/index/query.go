package index

import "gorm.io/gorm"

func HasIndex(db *gorm.DB, name string, data interface{}) bool {
	return db.Migrator().HasIndex(data, name)
}
