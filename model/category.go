package model

import "gorm.io/gorm"

// Category has Items
type Category struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Items       []Item
}
