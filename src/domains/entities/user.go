package entities

import (
	"errors"
	"net/mail"
	"regexp"
	"time"

	"gorm.io/gorm"
)

var (
	ErrInvalidEmail      = errors.New("invalid email")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrSamePassword      = errors.New("same password")
	ErrIncorrectPassword = errors.New("incorrect password")
	ErrTooLongEmail      = errors.New("email is too long")
	ErrEmailAlreadyExist = errors.New("this email already exists")
	ErrTooLongUserName       = errors.New("name is too long")
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

func (u *User) ValidateName() error {
	if len(u.Name) > 256 {
		return ErrTooLongUserName
	}
	return nil
}

const emailPattern string = `^[a-zA-Z0-9_+-.]+@[a-z0-9-.]+\.[a-z]+$`

var emailRegexp *regexp.Regexp = regexp.MustCompile(emailPattern)

func (u *User) ValidateEmail() error {
	if len(u.Email) > 256 {
		return ErrTooLongEmail
	}
	_, err := mail.ParseAddress(u.Email) // RFC 5322に定義されたアドレスにパースできるか確認
	if err != nil {
		return ErrInvalidEmail
	}
	// 上記だとかなり不自然なアドレスも通ってしまうので厳しい正規表現を書いている。
	if !emailRegexp.MatchString(u.Email) {
		return ErrInvalidEmail
	}

	return nil
}
