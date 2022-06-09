package index

import "gorm.io/gorm"

//目前未测试出有用，后续有待验证，需要所以可以在表结构中tag中写入
func CreateIndex(db *gorm.DB, name string, data interface{}) error {
	return db.Migrator().CreateIndex(data, name)
}
