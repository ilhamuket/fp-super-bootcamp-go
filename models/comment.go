package models

import "time"

// Comment represents the comment entity
type Comment struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
	UserID    uint       `json:"user_id"`
	User      User       `gorm:"foreignKey:UserID" json:"user"`
	NewsID    uint       `json:"news_id"`
	Text      string     `json:"text" binding:"required"`
}

type CommentInputSwagger struct {
	NewsID uint   `json:"news_id" binding:"required"`
	Text   string `json:"text" binding:"required"`
}
