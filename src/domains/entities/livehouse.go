package entities

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrTooLongLiveHouseName = errors.New("too long name")
	ErrTooLongLiveHouseMemo = errors.New("too long memo")
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

func (h *LiveHouse) ValidateName() error {
	if len(h.Name) > 256 {
		return ErrTooLongLiveName
	}
	return nil
}

func (h *LiveHouse) ValidateMemo() error {
	if len(h.Name) > 1024 {
		return ErrTooLongLiveMemo
	}
	return nil
}

// TODO: 電話番号validation
// const phonePattern string = "/0[0-9]{1,4}-[0-9]{1,6}(-[0-9]{0,5})?/"

// func (u *User) ValidatePhone() error {
// 	if len(u.Email) > 256 {
// 		return ErrTooLongEmail
// 	}
// 	_, err := mail.ParseAddress(u.Email) // RFC 5322に定義されたアドレスにパースできるか確認
// 	if err != nil {
// 		return ErrInvalidEmail
// 	}
// 	// 上記だとかなり不自然なアドレスも通ってしまうので厳しい正規表現を書いている。
// 	if !emailRegexp.MatchString(u.Email) {
// 		return ErrInvalidEmail
// 	}

// 	return nil
// }
