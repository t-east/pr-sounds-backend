package http_tests

import (
	"github.com/t-east/pr-sounds-backend/domains/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// テスト用のDBを取得
func LoadTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&entities.User{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
