package aggregate

import (
	"fmt"

	"github.com/andreis3/users-ms/internal/domain/entity"
	"github.com/andreis3/users-ms/internal/domain/validator"
	"github.com/andreis3/users-ms/internal/util"
)

type CustomerProfile struct {
	Customer  entity.Customer
	Addresses []entity.Address
	util.ErrorHandler
}

func NewUserProfile(custome entity.Customer, addresses []entity.Address) *CustomerProfile {
	userProfile := &CustomerProfile{
		Customer:  custome,
		Addresses: addresses,
	}
	return userProfile
}

func (u *CustomerProfile) Validate() *util.ErrorHandler {
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

	return u.InvalidCustomerAndAddres(mainValidator)
}
