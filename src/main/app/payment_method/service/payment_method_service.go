package service

import (
	domainPaymentMethod "service-api/src/main/app/payment_method/domain"
	inputPaymetnMethod "service-api/src/main/util/input/payment_method"
)

type PaymentMethodService interface {
	InsertPaymentMethod(paymentMethod inputPaymetnMethod.PaymentMethodInput) (domainPaymentMethod.PaymentMethod, error)
	GetAllPaymentMethod() ([]domainPaymentMethod.PaymentMethod, error)
	GetPaymentMethodById(id int) (domainPaymentMethod.PaymentMethod, error)
	UpdatePaymentMethod(paymentMethod inputPaymetnMethod.PaymentMethodInput) (*domainPaymentMethod.PaymentMethod, error)
	DeletePaymentMethod(id int) error
}
