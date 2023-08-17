package paymentmethod

import (
	"net/http"
	servicePaymentMethod "service-api/src/main/app/payment_method/service"
	infra "service-api/src/main/infra"
	formatPaymentMethod "service-api/src/main/util/formatter/payment_method"
	inputPaymetnMethod "service-api/src/main/util/input/payment_method"

	"github.com/gin-gonic/gin"
)

type paymentMethodController struct {
	paymentMethodService servicePaymentMethod.PaymentMethodService
}

func NewPaymentMethodController(service servicePaymentMethod.PaymentMethodService) *paymentMethodController {
	return &paymentMethodController{
		paymentMethodService: service,
	}
}

// get all payment method
func (controller *paymentMethodController) GetAllPaymentMethod(context *gin.Context) {
	paymentMethods, err := controller.paymentMethodService.GetAllPaymentMethod()
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", http.StatusBadRequest, nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := formatPaymentMethod.FormatPaymentMethods(paymentMethods)
	response := infra.NewResponseAPI("success", "success", http.StatusOK, formatter)
	context.JSON(http.StatusOK, response)
}

// create payment method
func (controller *paymentMethodController) InsertPaymentMethod(context *gin.Context) {
	var input inputPaymetnMethod.PaymentMethodInput
	err := context.ShouldBindJSON(&input)
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", http.StatusBadRequest, nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	paymentMethod, err := controller.paymentMethodService.InsertPaymentMethod(input)
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", http.StatusBadRequest, nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := formatPaymentMethod.FormatPaymentMethod(paymentMethod)
	response := infra.NewResponseAPI("success", "success", http.StatusOK, formatter)
	context.JSON(http.StatusOK, response)
}
