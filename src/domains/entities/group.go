package entities

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrInvalidName = errors.New("invalid name")
)

type GroupName string

var (
	Leader         GroupName = "leader"
	SubLeader      GroupName = "sub leader"
	Welcome        GroupName = "welcome"
	Treasurer      GroupName = "treasurer"
	LiveAdmin      GroupName = "live admin"
	CampAdmin      GroupName = "camp admin"
	PhotoAdmin     GroupName = "photo admin"
	VideoAdmin     GroupName = "video admin"
	PartyAdmin     GroupName = "party admin"
	SchollFesAdmin GroupName = "school fes admin"
)

type Group struct {
	ID        uint           `json:"id"`
	Name      GroupName      `gorm:"size:256;unique;not null" json:"name"`
	Memo      string         `gorm:"size:1024" json:"memo"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (g *Group) ValidateName() error {
	switch g.Name {
	case Leader, SubLeader, Welcome, Treasurer, LiveAdmin, CampAdmin, PhotoAdmin, VideoAdmin, PartyAdmin, SchollFesAdmin:
		return nil
	default:
		return ErrInvalidName
	}
}
