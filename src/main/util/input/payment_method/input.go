package paymentmethod

type PaymentMethodInput struct {
	Name string `json:"name" binding:"required"`
}
