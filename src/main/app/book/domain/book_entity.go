package domain

import (
	domainGenre "service-api/src/main/app/genre/domain"
	domainWishlist "service-api/src/main/app/wishlist/domain"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string                    `gorm:"column:name;not null" json:"name"`
	Author      string                    `gorm:"column:author;not null" json:"author"`
	Price       float64                   `gorm:"column:price;not null" json:"price"`
	Isbn        int                       `gorm:"column:isbn;not null" json:"isbn"`
	Description string                    `gorm:"column:description;not null'type:TEXT" json:"description"`
	Genre       []domainGenre.Genre       `gorm:"many2many:book_genres;" json:"genre"`
	Wishlist    []domainWishlist.Wishlist `gorm:"foreignKey:BookID" json:"-"`
}

func (Book) TableName() string {
	return "books"
}
