package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint           `json:"id"`
	Name           string         `gorm:"size:256;not null" json:"name"`
	Email          string         `gorm:"size:256;unique;not null;index" json:"email"`
	Password       string         `gorm:"not null" json:"password"`
	HashedPassword string         `gorm:"not null" json:"-"`
	Icon           string         `json:"icon"`
	IsActive       bool           `gorm:"default:false;not null" json:"-"`
	IsSuperUser    bool           `gorm:"default:false;not null" json:"-"`
	GroupID        uint           `json:"group_id"`
	Group          Group          `json:"group"`
	Band           Band           `json:"band"`
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `json:"-"`
}
