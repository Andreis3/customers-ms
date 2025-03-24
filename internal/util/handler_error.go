package util

import (
	"net/http"

	"github.com/andreis3/users-ms/internal/domain/validator"
)

type ErrorHandler struct {
	Status      int
	Origin      string
	ClientError []string
	LogError    []string
}

func (eh *ErrorHandler) InvalidCustomerAndAddres(validate *validator.Validator) *ErrorHandler {
	return &ErrorHandler{
		Status:      http.StatusBadRequest,
		Origin:      "CustomerProfile.Validate",
		ClientError: validate.Errors(),
		LogError:    validate.Errors(),
	}
}
