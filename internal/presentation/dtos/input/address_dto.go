package input

import "github.com/andreis3/customers-ms/internal/domain/entity/address"

type AddressDTO struct {
	Street     string `json:"street"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

func (a *AddressDTO) MapperToAggregate() address.Address {
	return address.NewBuilder().
		WithStreet(a.Street).
		WithNumber(a.Number).
		WithComplement(a.Complement).
		WithCity(a.City).
		WithState(a.State).
		WithCountry(a.Country).
		WithPostalCode(a.PostalCode).
		Build()
}
