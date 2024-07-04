package models

import "github.com/jinzhu/gorm"

type News struct {
	gorm.Model
	UserID   uint   `gorm:"not null"`
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	Comments []Comment
}
