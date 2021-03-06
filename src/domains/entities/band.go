package entities

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrTooLongBandName = errors.New("too long name")
)

type Band struct {
	ID        uint           `json:"id"`
	Name      string         `gorm:"size:256;not null" json:"name"`
	LiveID    uint           `gorm:"size:256;not null" json:"live_id"`
	Live      Live           `json:"live"`
	Members   []User         `json:"members"`
	Order     uint           `json:"order"`
	PDFkey    string         `json:"pdf_key"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (b *Band) ValidateName() error {
	if len(b.Name) > 256 {
		return ErrTooLongBandName
	}
	return nil
}
