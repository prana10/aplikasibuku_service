package domain

import (
	domainTransaction "service-api/src/main/app/transactions/domain"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	Name        string                          `gorm:"type:varchar(255);not null" json:"name"`
	Transaction []domainTransaction.Transaction `gorm:"foreignKey:PaymentMethodID" json:"-"`
}

func (PaymentMethod) TableName() string {
	return "payment_methods"
}
