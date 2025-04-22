package input

import (
	"github.com/andreis3/users-ms/internal/domain/aggregate"
	"github.com/andreis3/users-ms/internal/domain/entity/address"
	"github.com/andreis3/users-ms/internal/domain/entity/customer"
	"github.com/andreis3/users-ms/internal/util"
)

type CreatedCustomerDTO struct {
	Email       string          `json:"email"`
	Password    string          `json:"password"`
	FirstName   string          `json:"first_name"`
	LastName    string          `json:"last_name"`
	CPF         string          `json:"cpf"`
	DateOfBirth util.CustomDate `json:"date_of_birth"`
	Addresses   []AddressDTO    `json:"addresses"`
}

func (c *CreatedCustomerDTO) MapperToAggregate() aggregate.CustomerProfile {
	customer := customer.NewBuilder().
		WithEmail(c.Email).
		WithPassword(c.Password).
		WithFirstName(c.FirstName).
		WithLastName(c.LastName).
		WithCPF(c.CPF).
		WithDateOfBirth(c.DateOfBirth.Time).
		Build()

	var addresses []address.Address

	if len(c.Addresses) > 0 {
		for _, a := range c.Addresses {
			addresses = append(addresses, a.MapperToAggregate())
		}
	}
	return *aggregate.NewCustomerProfile(customer, addresses)
}
