package gorm

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mission/task4/gorm/model"
	"os"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	err := godotenv.Load()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Print(err)
	}

	DB.AutoMigrate(&model.User{})
	return DB
}
