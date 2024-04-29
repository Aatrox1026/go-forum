package model

import "gorm.io/gorm"

const (
	ROLE_ADMIN int64 = 1
)

type User struct {
	gorm.Model
	Name   string
	Passwd string
	Role   int64
}
