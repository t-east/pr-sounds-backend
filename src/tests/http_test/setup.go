package http

import (
	"github.com/t-east/pr-sounds-backend/src/domains/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// テスト用のDBを取得
func LoadTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&entities.User{})
	return db, nil
}
