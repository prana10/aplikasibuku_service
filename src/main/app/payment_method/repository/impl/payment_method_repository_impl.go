package impl

import (
	domainPaymentMethod "service-api/src/main/app/payment_method/domain"

	"gorm.io/gorm"
)

type paymentMethodRepository struct {
	db *gorm.DB
}

func NewPaymentMethodRepository(db *gorm.DB) *paymentMethodRepository {
	return &paymentMethodRepository{db}
}

func (repository *paymentMethodRepository) CreatePaymentMethod(paymentMethod domainPaymentMethod.PaymentMethod) (domainPaymentMethod.PaymentMethod, error) {
	err := repository.db.Create(&paymentMethod).Error
	if err != nil {
		return paymentMethod, err
	}
	return paymentMethod, err
}

func (repository *paymentMethodRepository) GetAllPaymentMethod() ([]domainPaymentMethod.PaymentMethod, error) {
	var paymentMethods []domainPaymentMethod.PaymentMethod
	err := repository.db.Find(&paymentMethods).Error
	if err != nil {
		return paymentMethods, err
	}

	if len(paymentMethods) == 0 {
		return paymentMethods, gorm.ErrRecordNotFound
	}

	return paymentMethods, err
}

func (repository *paymentMethodRepository) GetPaymentMethodByID(paymentMethodID uint) (domainPaymentMethod.PaymentMethod, error) {
	var paymentMethod domainPaymentMethod.PaymentMethod
	err := repository.db.Where("id = ?", paymentMethodID).First(&paymentMethod).Error
	if err != nil {
		return paymentMethod, err
	}

	if paymentMethod.ID == 0 {
		return paymentMethod, gorm.ErrRecordNotFound
	}

	return paymentMethod, err
}

func (repository *paymentMethodRepository) UpdatePaymentMethodByID(paymentMethodID uint, paymentMethod domainPaymentMethod.PaymentMethod) (domainPaymentMethod.PaymentMethod, error) {
	err := repository.db.Model(&paymentMethod).Where("id = ?", paymentMethodID).Updates(domainPaymentMethod.PaymentMethod{Name: paymentMethod.Name}).Error
	if err != nil {
		return paymentMethod, err
	}

	return paymentMethod, err
}

func (repository *paymentMethodRepository) DeletePaymentMethodByID(paymentMethodID uint) (domainPaymentMethod.PaymentMethod, error) {
	var paymentMethod domainPaymentMethod.PaymentMethod
	err := repository.db.Where("id = ?", paymentMethodID).Delete(&paymentMethod).Error
	if err != nil {
		return paymentMethod, err
	}

	if paymentMethod.ID == 0 {
		return paymentMethod, gorm.ErrRecordNotFound
	}

	return paymentMethod, err
}
