package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title	string `gorm:"not null" json:"title" form:"title" valid:"required~ Your Title is required"`
	Photo_URL	string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~ Your Photo URL is required"`
	Caption   string `gorm:"type:varchar(255);" validate:"required" json:"caption"`
	UserID    uint
	User      *User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	err = nil
	return
}