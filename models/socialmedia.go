package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name            string `gorm:"type:varchar(255);not null" validate:"required" json:"name"`
	SocialMedia_URL string `gorm:"type:varchar(255);not null" validate:"required" json:"social_media_url"`
	UserID          uint
	User            *User
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(s)
	if err != nil {
		return err
	}

	err = nil
	return
}