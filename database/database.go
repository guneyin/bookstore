package database

import (
	"context"
	"github.com/guneyin/bookstore/entity"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const dbPath = "data/data.db"

var DB *gorm.DB

func GetDB(ctx context.Context, debug ...bool) *gorm.DB {
	df := false
	db := DB.WithContext(ctx)

	if len(debug) > 0 {
		df = debug[0]
	}

	if df {
		return db.Debug()
	}

	return db
}

func Connect() error {
	checkDBPath()

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	slog.Info("[DATABASE]::CONNECTED")
	DB = db

	return db.AutoMigrate(
		&entity.User{},
		&entity.Book{},
		&entity.Cart{},
		&entity.Order{},
		&entity.OrderItem{},
	)
}

func checkDBPath() {
	dir := filepath.Dir(dbPath)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_ = os.MkdirAll(dir, os.ModePerm)
	}
}
