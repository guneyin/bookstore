package util

import (
	"github.com/guneyin/bookstore/repo/user"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&user.User{})
}
