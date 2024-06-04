package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(20);not null"`
	Email string `gorm:"uniqueIndex;type:varchar(128);not null"`
}

func (User) TableName() string {
	return "users"
}
