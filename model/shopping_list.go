package model

import "gorm.io/gorm"

//ShoppingList has Title
type ShoppingList struct {
	gorm.Model
	Title string `gorm:"not null"`
	Items []Item
	Users []User `gorm:"many2many:user_shopping_lists;association_foreignkey:UserID"`
}
