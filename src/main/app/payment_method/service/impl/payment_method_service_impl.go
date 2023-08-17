package impl

import (
	domainPaymentMethod "service-api/src/main/app/payment_method/domain"
	repositoryPaymentMethod "service-api/src/main/app/payment_method/repository"
	inputPaymetnMethod "service-api/src/main/util/input/payment_method"
)

type paymentMethodService struct {
	paymentMethodRepo repositoryPaymentMethod.PaymentMethodRepository
}

func NewPaymentMethodService(repo repositoryPaymentMethod.PaymentMethodRepository) *paymentMethodService {
	return &paymentMethodService{
		paymentMethodRepo: repo,
	}
}

func (service *paymentMethodService) GetAllPaymentMethod() ([]domainPaymentMethod.PaymentMethod, error) {
	paymentMethods, err := service.paymentMethodRepo.GetAllPaymentMethod()
	if err != nil {
		return paymentMethods, err
	}

	return paymentMethods, nil
}

// create payment method
func (service *paymentMethodService) InsertPaymentMethod(inputPaymentMethod inputPaymetnMethod.PaymentMethodInput) (domainPaymentMethod.PaymentMethod, error) {
	newPaymentMethod := domainPaymentMethod.PaymentMethod{}
	newPaymentMethod.Name = inputPaymentMethod.Name

	paymentMethod, err := service.paymentMethodRepo.CreatePaymentMethod(newPaymentMethod)
	if err != nil {
		return paymentMethod, err
	}

	return paymentMethod, nil
}

// get payment method by id
func (service *paymentMethodService) GetPaymentMethodByID(paymentMethodID uint) (domainPaymentMethod.PaymentMethod, error) {
	paymentMethod, err := service.paymentMethodRepo.GetPaymentMethodByID(paymentMethodID)
	if err != nil {
		return paymentMethod, err
	}

	return paymentMethod, nil
}

// update payment method by id
func (service *paymentMethodService) UpdatePaymentMethodByID(paymentMethodID uint, paymentMethod domainPaymentMethod.PaymentMethod) (domainPaymentMethod.PaymentMethod, error) {
	paymentMethod, err := service.paymentMethodRepo.UpdatePaymentMethodByID(paymentMethodID, paymentMethod)
	if err != nil {
		return paymentMethod, err
	}

	return paymentMethod, nil
}

// delete payment method by id
func (service *paymentMethodService) DeletePaymentMethod(paymentMethodID uint) error {
	err := service.paymentMethodRepo.DeletePaymentMethodByID(paymentMethodID)
	if err != nil {
		return err
	}

	return nil
}
