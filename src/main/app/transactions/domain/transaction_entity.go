package domain

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Code               uint `gorm:"unique;column:code;not null" json:"transaction_id"`
	UserID             uint `gorm:"column:user_id;not null" json:"user_id"`
	BookID             uint `gorm:"column:book_id;not null" json:"book_id"`
	IsPaymentCompleted bool `gorm:"column:is_payment_completed;not null" json:"is_payment_completed"`
	PaymentMethodID    uint `gorm:"column:payment_method_id;not null" json:"payment_method_id"`
}

func (t *Transaction) TableName() string {
	return "transactions"
}
