package models

import (
	"time"
)

// User represents the user entity
type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
	Username  string     `gorm:"unique;not null" json:"username"`
	Password  string     `gorm:"not null" json:"password"`
	Email     string     `gorm:"unique;not null" json:"email"`
	Roles     []Role     `gorm:"many2many:user_roles" json:"roles"`
	Profile   Profile    `json:"profile"`
	News      []News     `json:"news"`
	Comments  []Comment  `json:"comments"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

// LoginInput adalah struktur input untuk login pengguna
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ProfileInput struct {
	Bio     string `json:"bio" binding:"required"`
	Picture string `json:"picture" binding:"required"`
}

type ChangePasswordInput struct {
	Password string `json:"password" binding:"required"`
}
