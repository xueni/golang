package data

import "gorm.io/gorm"

//可单个插入或者批量
func Insert(db *gorm.DB, data interface{}) (dbNew *gorm.DB, err error) {
	dbNew = db.Create(data)
	err = dbNew.Error
	return
}

func SelectInsert(db *gorm.DB, data interface{}) (dbNew *gorm.DB, err error) {
	dbNew = db.Select("name", "value").Create(data)
	err = dbNew.Error
	return
}

func OmitInsert(db *gorm.DB, data interface{}) (dbNew *gorm.DB, err error) {
	dbNew = db.Omit("name").Create(data)
	err = dbNew.Error
	return
}

func ModelInsert(db *gorm.DB, data interface{}) error {
	return db.Model(data).Create(data).Error
}
