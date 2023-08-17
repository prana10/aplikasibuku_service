package domain

import "gorm.io/gorm"

type PaymentMethod struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}

func (PaymentMethod) TableName() string {
	return "payment_methods"
}
