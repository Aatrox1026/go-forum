package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	ROLE_ADMINISTRATOR int64 = 1
	ROLE_MANAGER       int64 = 2
	ROLE_NORMAL        int64 = 3
)

type User struct {
	ID        int64          `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password,omitempty"`
	Role      int64          `json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (user *User) Safe() (safe *User) {
	safe = new(User)
	*safe = *user
	safe.Password = ""
	return safe
}
