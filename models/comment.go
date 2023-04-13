package models

type Comment struct {
	GormModel
	Message string `gorm:"type:varchar(255);not null" validate:"required"`
	UserID uint
	User *User
}