package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `dao:"type:varchar(20);not null"`
	Email string `dao:"uniqueIndex;type:varchar(128);not null"`
}

func (User) TableName() string {
	return "users"
}
