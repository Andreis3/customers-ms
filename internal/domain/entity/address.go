package entity

import (
	"time"

	"github.com/andreis3/users-ms/internal/domain/validator"
)

type Address struct {
	ID         string
	Street     string
	Number     string
	Complement string
	City       string
	State      string
	PostalCode string
	Country    string
	createAT   time.Time
	UpdatedAT  time.Time
	Validator  validator.Validator
}

func AddressBuilder() *Address {
	return &Address{}
}

func (a *Address) SetID(id string) *Address {
	a.ID = id
	return a
}

func (a *Address) SetStreet(street string) *Address {
	a.Street = street
	return a
}

func (a *Address) SetNumber(number string) *Address {
	a.Number = number
	return a
}

func (a *Address) SetComplement(complement string) *Address {
	a.Complement = complement
	return a
}

func (a *Address) SetCity(city string) *Address {
	a.City = city
	return a
}

func (a *Address) SetState(state string) *Address {
	a.State = state
	return a
}

func (a *Address) SetPostalCode(postalCode string) *Address {
	a.PostalCode = postalCode
	return a
}

func (a *Address) SetCountry(country string) *Address {
	a.Country = country
	return a
}

func (a *Address) SetCreatedAT(createdAt time.Time) *Address {
	a.createAT = createdAt
	return a
}

func (a *Address) SetUpdatedAT(updatedAt time.Time) *Address {
	a.UpdatedAT = updatedAt
	return a
}

func (a *Address) Build() *Address {
	return a
}

func (a *Address) Validate() *validator.Validator {
	a.Validator.CheckField(validator.NotBlank(a.Street), "street", validator.NotBlankField)
	a.Validator.CheckField(validator.NotBlank(a.Number), "number", validator.NotBlankField)
	a.Validator.CheckField(validator.NotBlank(a.City), "city", validator.NotBlankField)
	a.Validator.CheckField(validator.NotBlank(a.State), "state", validator.NotBlankField)
	a.Validator.CheckField(validator.NotBlank(a.PostalCode), "postal_code", validator.NotBlankField)
	a.Validator.CheckField(validator.NotBlank(a.Country), "country", validator.NotBlankField)
	return &a.Validator
}
