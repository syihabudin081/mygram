package models

import "time"

// GormModel is a model that contains common columns for all tables.
type GormModel struct {
	ID		uint `gorm:"primary_key" json:"id" ` 
	CreatedAt	*time.Time `json:"created_at,omitempty"`
	UpdatedAt	*time.Time `json:"updated_at,omitempty"`
}
