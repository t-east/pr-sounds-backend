package entities

import (
	"time"

	"gorm.io/gorm"
)

type Group struct {
	ID        uint           `json:"id"`
	Name      string         `gorm:"size:256;not null" json:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
