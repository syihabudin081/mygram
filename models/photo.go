package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title     string `gorm:"type:varchar(255);not null" validate:"required" json:"title"`
	Photo_URL string `gorm:"type:varchar(255);not null" validate:"required" json:"photo_url"`
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