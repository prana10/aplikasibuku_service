package repository

import (
	domainPaymentMethod "service-api/src/main/app/payment_method/domain"
)

type PaymentMethodRepository interface {
	CreatePaymentMethod(paymentMethod domainPaymentMethod.PaymentMethod) (domainPaymentMethod.PaymentMethod, error)
	GetAllPaymentMethod() ([]domainPaymentMethod.PaymentMethod, error)
	GetPaymentMethodByID(paymentMethodID uint) (domainPaymentMethod.PaymentMethod, error)
	UpdatePaymentMethodByID(paymentMethodID uint, paymentMethod domainPaymentMethod.PaymentMethod) (domainPaymentMethod.PaymentMethod, error)
	DeletePaymentMethodByID(paymentMethodID uint) error
}
