package domain

import "gorm.io/gorm"

type Wishlist struct {
	gorm.Model
	UserID uint
	BookID uint
}

func (Wishlist) TableName() string {
	return "wishlists"
}
