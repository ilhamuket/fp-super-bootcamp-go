package config

import (
	"final-project-golang-individu/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // import mysql driver
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/news_final?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Failed to connect database", err)
	}
	DB.LogMode(true)
	migrate()
}

func migrate() {
	DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Profile{}, &models.News{}, &models.Comment{}, &models.UserRole{})
	DB.Model(&models.Profile{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	DB.Model(&models.News{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	DB.Model(&models.Comment{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	DB.Model(&models.Comment{}).AddForeignKey("news_id", "news(id)", "CASCADE", "CASCADE")
}
