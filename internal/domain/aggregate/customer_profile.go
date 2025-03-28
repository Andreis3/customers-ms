package aggregate

import (
	"fmt"

	"github.com/andreis3/users-ms/internal/domain/entity"
	"github.com/andreis3/users-ms/internal/domain/errors"
	"github.com/andreis3/users-ms/internal/domain/validator"
)

type CustomerProfile struct {
	Customer  entity.Customer
	Addresses []entity.Address
}

func NewUserProfile(custome entity.Customer, addresses []entity.Address) *CustomerProfile {
	userProfile := &CustomerProfile{
		Customer:  custome,
		Addresses: addresses,
	}
	return userProfile
}

func (u *CustomerProfile) Validate() *errors.DomainErrors {
	mainValidator := validator.NewValidator()

	validateUser := u.Customer.Validate()
	mainValidator.Merge(validateUser)

	for i, address := range u.Addresses {
		addresValidator := address.Validate()

		for key, message := range addresValidator.FieldErrors {
			prefixedKey := fmt.Sprintf("address[%d].%s", i, key)
			mainValidator.AddFieldError(prefixedKey, message)
		}
	}

	if !mainValidator.HasErrors() {
		return nil
	}

	return errors.InvalidCustomerError(mainValidator)
}
