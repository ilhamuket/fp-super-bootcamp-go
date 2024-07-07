package models

import "time"

// Role represents the role entity
type Role struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
	Name      string     `json:"name" binding:"required"`
	Users     []User     `gorm:"many2many:user_roles" json:"users"`
}
