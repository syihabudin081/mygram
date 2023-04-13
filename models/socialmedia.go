package models

type SocialMedia struct {
	GormModel
	Name string `gorm:"type:varchar(255);not null" validate:"required" json:"name"`
	SocialMedia_URL string `gorm:"type:varchar(255);not null" validate:"required" json:"social_media_url"`
	UserID uint
	User *User
}