package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	NewsID uint   `gorm:"not null"`
	Text   string `gorm:"not null"`
}
