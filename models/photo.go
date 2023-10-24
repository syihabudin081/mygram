package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// represent model for photo
type Photo struct {
	GormModel
	
	Photo_URL	string ` json:"photo_url" form:"photo_url"`
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