package database

import (
	"context"
	"github.com/guneyin/bookstore/entity"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const dbPath = "data/data.db"

var (
	dbOnce sync.Once
	DB     *gorm.DB
)

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
	var errDB error

	dbOnce.Do(func() {
		checkDBPath()

		db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
			NowFunc: func() time.Time { return time.Now().Local() },
			Logger:  logger.Default.LogMode(logger.Info),
		})
		errDB = err

		DB = db

		slog.Info("[DATABASE]::CONNECTED")
	})
	if errDB != nil {
		return errDB
	}

	return DB.AutoMigrate(
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
