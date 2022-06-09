package table

import "gorm.io/gorm"

func DropTable(db *gorm.DB, name string) error {
	return db.Migrator().DropTable(name)
}
