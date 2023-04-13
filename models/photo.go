package models

type Photo struct {
	Title string `gorm:"type:varchar(255);not null" validate:"required" json:"title"`
	Photo_URL string `gorm:"type:varchar(255);not null" validate:"required" json:"photo_url"`
	Caption string `gorm:"type:varchar(255);" validate:"required" json:"caption"`
	PhotoID uint
	UserID uint
	Photo *Photo
	User *User
}