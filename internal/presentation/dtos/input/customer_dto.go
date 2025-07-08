package input

import (
	"github.com/andreis3/customers-ms/internal/domain/aggregate"
	"github.com/andreis3/customers-ms/internal/domain/entity"
	"github.com/andreis3/customers-ms/internal/util"
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
	customer := entity.BuilderCustomer().
		WithEmail(c.Email).
		WithPassword(c.Password).
		WithFirstName(c.FirstName).
		WithLastName(c.LastName).
		WithCPF(c.CPF).
		WithDateOfBirth(c.DateOfBirth.Time).
		Build()

	var addresses []entity.Address

	if len(c.Addresses) > 0 {
		for _, a := range c.Addresses {
			addresses = append(addresses, a.MapperToAggregate())
		}
	}
	return *aggregate.NewCustomerProfile(customer, addresses)
}
