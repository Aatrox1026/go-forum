package model

import "gorm.io/gorm"

const (
	ROLE_ADMINISTRATOR int64 = 1
	ROLE_MANAGER       int64 = 2
	ROLE_NORMAL        int64 = 3
)

type User struct {
	ID       int64  `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     int64  `json:"role"`
	gorm.Model
}
