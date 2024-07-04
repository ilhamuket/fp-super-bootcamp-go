package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Users []User `gorm:"many2many:user_roles;"`
}
