package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Name       string `gorm:"index"`
	Username   string `gorm:"uniqueIndex"`
	Email      string `gorm:"uniqueIndex"`
	Address    string
	Phone      string
}

type UserList []User
