package entities

import (
	"time"

	"gorm.io/gorm"
)

type Live struct {
	ID        uint           `json:"id"`
	Name      string         `gorm:"size:256;not null" json:"name"`
	Date      time.Time      `json:"date"` // TODO: 型変わるかも
	LiveHouse LiveHouse      `json:"live_house"`
	Bands     []Band         `json:"bands"`
	Price     uint           `json:"price"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
