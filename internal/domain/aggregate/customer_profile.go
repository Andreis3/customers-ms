package aggregate

import (
	"fmt"

	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
	"github.com/andreis3/customers-ms/internal/domain/entity/address"
	"github.com/andreis3/customers-ms/internal/domain/entity/customer"
	"github.com/andreis3/customers-ms/internal/domain/validator"
)

type CustomerProfile struct {
	Customer  customer.Customer
	Addresses []address.Address
}

func NewCustomerProfile(custome customer.Customer, addresses []address.Address) *CustomerProfile {
	userProfile := &CustomerProfile{
		Customer:  custome,
		Addresses: addresses,
	}
	return userProfile
}

func (u *CustomerProfile) Validate() *apperror.Error {
	mainValidator := validator.New()

	validateUser := u.Customer.Validate()
	mainValidator.Merge(validateUser)

	for i, address := range u.Addresses {
		addresValidator := address.Validate()

		for key, messages := range addresValidator.FieldErrors {
			prefixedKey := fmt.Sprintf("address[%d].%s", i, key)
			for _, msg := range messages {
				mainValidator.AddFieldError(prefixedKey, msg)
			}
		}
	}

	if !mainValidator.HasErrors() {
		return nil
	}

	return apperror.InvalidCustomerError(mainValidator)
}
