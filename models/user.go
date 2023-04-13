package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Email        string        `gorm:"type:varchar(255);unique_index;not null" validate:"required,email"`
	Username     string        `gorm:"type:varchar(255);unique_index;not null" validate:"required"`
	Password     string        `gorm:"type:varchar(255);not null" validate:"required,min=6"`
	Age          uint8         `gorm:"not null" validate:"required,gte=9"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_medias"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	u.Password = helpers.HashPassword([]byte(u.Password))
	err = nil
	return
}