package model

import "gorm.io/gorm"

type Gamer struct {
	gorm.Model
	Name  string `dao:"type:varchar(20);not null"`
	Email string `dao:"uniqueIndex;type:varchar(128);not null"`
}

func (Gamer) TableName() string {
	return "gamers"
}
