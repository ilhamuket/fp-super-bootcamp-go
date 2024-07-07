package models

// UserRole represents the user-role relationship
type UserRole struct {
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
}
