package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Define UserRole type and constants
type UserRole string

const (
	AdminRole UserRole = "admin"
	UsersRole UserRole = "user"
	GuestRole UserRole = "guest"
)

// User struct
type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username  string    `gorm:"type:varchar(100);not null" json:"username"`
	Email     string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	Role      UserRole  `gorm:"type:varchar(20);default:'user';not null" json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}

// BeforeCreate hook will generate a UUID before inserting a new record
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.ID = uuid.New()
	return
}
