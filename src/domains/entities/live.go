package entities

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrTooLongLiveName = errors.New("too long name")
	ErrTooLongLiveMemo = errors.New("too long memo")
	ErrInvalidSchedule = errors.New("invalid schdule")
)

type Live struct {
	ID        uint           `json:"id"`
	Name      string         `gorm:"size:256;not null" json:"name"`
	Memo      string         `gorm:"size:1024" json:"memo"`
	StartAt   time.Time      `json:"start_at"` // TODO: 型変わるかも
	EndAt     time.Time      `json:"end_at"`   // TODO: 型変わるかも
	LiveHouse LiveHouse      `json:"live_house"`
	Bands     []Band         `json:"bands"`
	Price     uint           `json:"price"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (l *Live) ValidateName() error {
	if len(l.Name) > 256 {
		return ErrTooLongLiveName
	}
	return nil
}

func (l *Live) ValidateSchedule() error {
	if l.StartAt.After(l.EndAt) {
		return ErrInvalidSchedule
	}
	return nil
}

func (l *Live) ValidateMemo() error {
	if len(l.Memo) > 1024 {
		return ErrTooLongLiveMemo
	}
	return nil
}
