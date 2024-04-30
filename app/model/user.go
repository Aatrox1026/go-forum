package model

import "gorm.io/gorm"

const (
	ROLE_ADMIN int64 = 1
)

type User struct {
	ID     int64 `gorm:"primaryKey;autoIncrement:false"`
	Name   string
	Passwd string
	Role   int64
	gorm.Model
}
