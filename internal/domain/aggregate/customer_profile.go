package aggregate

import (
	"fmt"

	"github.com/andreis3/customers-ms/internal/domain/entity"
	"github.com/andreis3/customers-ms/internal/domain/errors"
	"github.com/andreis3/customers-ms/internal/domain/validator"
)

type CustomerProfile struct {
	Customer  entity.Customer
	Addresses []entity.Address
}

func NewCustomerProfile(custome entity.Customer, addresses []entity.Address) *CustomerProfile {
	userProfile := &CustomerProfile{
		Customer:  custome,
		Addresses: addresses,
	}
	return userProfile
}

func (u *CustomerProfile) Validate() *errors.Error {
	mainValidator := validator.New()

	validateCustomer := u.Customer.Validate()
	mainValidator.Merge(validateCustomer)

	for i, address := range u.Addresses {
		addresValidator := address.Validate()

		for key, messages := range addresValidator.FieldErrors {
			prefixedKey := fmt.Sprintf("addresses[%d].%s", i, key)
			for _, msg := range messages {
				mainValidator.AddFieldError(prefixedKey, msg)
			}
		}
	}

	if !mainValidator.HasErrors() {
		return nil
	}

	return errors.InvalidCustomerError(mainValidator)
}
