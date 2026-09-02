package auth

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleSuperAdmin Role = "super_admin"
	RoleAdmin      Role = "admin"
	RoleVolunteer  Role = "volunteer"
)

// User represents an admin or volunteer in the system
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	Email     string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"` // Omit password from JSON
	Role               Role           `gorm:"type:varchar(20);not null" json:"role"`
	RegistrationNumber string         `gorm:"type:varchar(50)" json:"registration_number"`
	ScholarType        string         `gorm:"type:varchar(20)" json:"scholar_type"` // e.g., "hosteller" or "day_scholar"
	YearOfStudy        int            `json:"year_of_study"`
	IsActive           bool           `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// LoginRequest represents the payload for authentication
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse represents the response after successful login
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
