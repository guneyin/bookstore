package util

import (
	"github.com/guneyin/bookstore/entity"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
		&entity.Book{},
	)
}
