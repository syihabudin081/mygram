package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message	string `gorm:"not null" json:"message" form:"message" valid:"required~ Your Message is required"`
	UserID  uint
	PhotoID uint
	Photo   *Photo
	User    *User
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	err = nil
	return
}