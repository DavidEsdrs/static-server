package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int
	Username string
}

type Contact struct {
	gorm.Model
	Phone  string
	UserId int
	User   User
}
