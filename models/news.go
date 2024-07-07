package models

import "time"

// News represents the news entity
type News struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
	UserID    uint       `json:"user_id"`
	Title     string     `json:"title" binding:"required"`
	Content   string     `json:"content" binding:"required"`
	Comments  []Comment  `gorm:"foreignkey:NewsID" json:"comments" preload:true`
}

type NewsInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
