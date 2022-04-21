package entities

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrTooLongTaskName = errors.New("too long name")
	ErrTooLongTaskMemo = errors.New("too long memo")
	ErrDuplicateAff = errors.New("duplicate affiliations")
)

type Task struct {
	ID        uint           `json:"id"`
	Name      string         `gorm:"size:256;not null" json:"name"`
	Memo      string         `gorm:"size:1024" json:"memo"`
	UserID    uint           `json:"user_id"`
	GroupID   uint           `json:"group_id"`
	BandID    uint           `json:"band_id"`
	Deadline  time.Time      `json:"deadline"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (t *Task) ValidateName() error {
	if len(t.Name) > 256 {
		return ErrTooLongTaskName
	}
	return nil
}

func (t *Task) ValidateMemo() error {
	if len(t.Memo) > 1024 {
		return ErrTooLongTaskMemo
	}
	return nil
}

func (t *Task) ValidateBelongTo() error {
	if t.GroupID == 0 && t.BandID == 0 {
		return ErrDuplicateAff
	}
	return nil
}
