package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Roles    []Role `gorm:"many2many:user_roles"`
	Profile  Profile
	News     []News
	Comments []Comment
}
