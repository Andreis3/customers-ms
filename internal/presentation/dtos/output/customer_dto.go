package output

import (
	"github.com/andreis3/customers-ms/internal/domain/entity"
	"github.com/andreis3/customers-ms/internal/util"
)

type CreatedCustomerDTO struct {
	Name        string          `json:"name"`
	Email       string          `json:"email"`
	DateOfBirth util.CustomDate `json:"date_of_birth"`
}

func CustomerOutputMapper(customer entity.Customer) CreatedCustomerDTO {
	return CreatedCustomerDTO{
		Email: customer.Email(),
		Name:  customer.FullName(),
		DateOfBirth: util.CustomDate{
			Time: customer.DateOfBirth(),
		},
	}
}
