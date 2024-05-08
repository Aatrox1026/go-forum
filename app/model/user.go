package model

import "gorm.io/gorm"

const (
	ROLE_ADMINISTRATOR int64 = 1
	ROLE_MANAGER       int64 = 2
	ROLE_NORMAL        int64 = 3
)

type User struct {
	ID       int64 `gorm:"primaryKey;autoIncrement:false"`
	Name     string
	Email    string
	Password string
	Role     int64
	gorm.Model
}
