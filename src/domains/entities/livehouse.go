package entities

import (
	"time"

	"gorm.io/gorm"
)

type LiveHouse struct {
	ID        uint           `json:"id"`
	Name      string         `gorm:"size:256;not null" json:"name"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	Website   string         `json:"website"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
