package model

import "gorm.io/gorm"

type Notice struct {
	gorm.Model
	Title       string `dao:"type:varchar(128);not null"`
	Content     string `dao:"type:text;not null"`
	PublishUser string `dao:"type:text;not null"`
}

func (Notice) TableName() string {
	return "notices"
}
