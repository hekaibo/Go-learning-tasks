package model

import "gorm.io/gorm"

type Gamer struct {
	gorm.Model
	Name  string `gorm:"type:varchar(20);not null"`
	Email string `gorm:"uniqueIndex;type:varchar(128);not null"`
}

func (Gamer) TableName() string {
	return "gamers"
}
