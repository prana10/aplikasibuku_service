package domain

import (
	domainOTP "service-api/src/main/app/otp/domain"
	domainTransaction "service-api/src/main/app/transactions/domain"
	domainWishlist "service-api/src/main/app/wishlist/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string                          `gorm:"column:name;not null;type:varchar(255)" json:"name"`
	Username    string                          `gorm:"column:username;not null;type:varchar(255);unique" json:"username"`
	Password    string                          `gorm:"column:password;not null;type:varchar(255)"`
	Email       string                          `gorm:"column:name;not null;type:varchar(255)" json:"email"`
	PhoneNumber string                          `gorm:"column:phone_number;not null;type:varchar(255)" json:"phone_number"`
	Address     string                          `gorm:"column:address;not null;type:TEXT" json:"address"`
	Otp         []domainOTP.Otp                 `gorm:"foreignKey:UserID" json:"otp"`
	Wishlist    []domainWishlist.Wishlist       `gorm:"foreignKey:UserID" json:"-"`
	Transaction []domainTransaction.Transaction `gorm:"foreignKey:TransactionID" json:"transaction"`
}

func (User) TableName() string {
	return "users"
}
