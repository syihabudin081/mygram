package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// represent model for social media
type SocialMedia struct {
	GormModel
	Name	string `gorm:"not null" json:"name" form:"name" valid:"required~ Your Name is required"`
	SocialMedia_URL	string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~ Your URL is required"`
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