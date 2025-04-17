package aggregate

import (
	"fmt"

	"github.com/andreis3/users-ms/internal/domain/apperrors"
	"github.com/andreis3/users-ms/internal/domain/entity/address"
	"github.com/andreis3/users-ms/internal/domain/entity/customer"
	"github.com/andreis3/users-ms/internal/domain/validator"
)

type CustomerProfile struct {
	Customer  customer.Customer
	Addresses []address.Address
}

func NewUserProfile(custome customer.Customer, addresses []address.Address) *CustomerProfile {
	userProfile := &CustomerProfile{
		Customer:  custome,
		Addresses: addresses,
	}
	return userProfile
}

func (u *CustomerProfile) Validate() *apperrors.AppErrors {
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

	return apperrors.InvalidCustomerError(mainValidator)
}
