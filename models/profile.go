package models

import "github.com/jinzhu/gorm"

type Profile struct {
	gorm.Model
	UserID  uint `gorm:"not null"`
	Bio     string
	Picture string
}
