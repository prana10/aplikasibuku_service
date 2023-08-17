package paymentmethod

import (
	domainPaymentMethod "service-api/src/main/app/payment_method/domain"
)

type PaymentMethodFormatter struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func FormatPaymentMethod(paymentMethod domainPaymentMethod.PaymentMethod) PaymentMethodFormatter {
	var paymentMethodFormatter PaymentMethodFormatter

	paymentMethodFormatter.ID = paymentMethod.ID
	paymentMethodFormatter.Name = paymentMethod.Name

	return paymentMethodFormatter
}

func FormatPaymentMethods(paymentMethods []domainPaymentMethod.PaymentMethod) []PaymentMethodFormatter {
	var paymentMethodFormatters []PaymentMethodFormatter

	for _, paymentMethod := range paymentMethods {
		paymentMethodFormatter := FormatPaymentMethod(paymentMethod)
		paymentMethodFormatters = append(paymentMethodFormatters, paymentMethodFormatter)
	}

	return paymentMethodFormatters
}
