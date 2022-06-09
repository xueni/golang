package models

import (
	"golang/gorm_use"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `gorm:"size:255"`
	Value string `gorm:"size:255;check:value <> 'jinzhu'"`
}

func (user *User) TableName() string {
	return gorm_use.TBALE_USER
}

type Page struct {
	Limit  int
	Offset int
}

type Result struct {
	Name  string
	Email string
}

type Email struct {
	gorm.Model
	UserID int    `gorm:"not null;comment:用户id"`
	Email  string `gorm:"size:255"`
}
