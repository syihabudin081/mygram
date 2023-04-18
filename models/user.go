package models

import (
	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// represent model for user
type User struct {
	GormModel
	Email		string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~ Your email is not valid"`
	Username	string `gorm:"not null" json:"username" form:"username" valid:"required~ Your Username is required"`
	Password	string `gorm:"not null" json:"password" form:"password" valid:"required~ Your password is required,minstringlength(6)~ Your password must be at least 6 characters"`
	Age          uint8         `gorm:"not null" validate:"required,gte=9" json:"age"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_medias"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	u.Password = helpers.HashPassword(u.Password)
	err = nil
	return
}