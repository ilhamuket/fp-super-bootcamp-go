package models

import "time"

// Profile represents the profile entity
type Profile struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
	UserID    uint       `json:"user_id"`
	Bio       string     `json:"bio"`
	Picture   string     `json:"picture"`
}
