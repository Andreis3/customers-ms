package address

import (
	"time"
)

type AddressBuilder struct {
	id         int64
	customerID int64
	street     string
	number     string
	complement string
	city       string
	state      string
	postalCode string
	country    string
	createdAt  time.Time
	updatedAt  time.Time
}

func NewBuilder() *AddressBuilder {
	return &AddressBuilder{}
}

func (b *AddressBuilder) Build() Address {
	return Address{
		id:         b.id,
		customerID: b.customerID,
		street:     b.street,
		number:     b.number,
		complement: b.complement,
		city:       b.city,
		state:      b.state,
		postalCode: b.postalCode,
		country:    b.country,
		createdAt:  b.createdAt,
		updatedAt:  b.updatedAt,
	}
}

func (b *AddressBuilder) WithID(id int64) *AddressBuilder {
	b.id = id
	return b
}

func (b *AddressBuilder) WithCustomerID(customerID int64) *AddressBuilder {
	b.customerID = customerID
	return b
}

func (b *AddressBuilder) WithStreet(street string) *AddressBuilder {
	b.street = street
	return b
}

func (b *AddressBuilder) WithNumber(number string) *AddressBuilder {
	b.number = number
	return b
}

func (b *AddressBuilder) WithComplement(complement string) *AddressBuilder {
	b.complement = complement
	return b
}

func (b *AddressBuilder) WithCity(city string) *AddressBuilder {
	b.city = city
	return b
}

func (b *AddressBuilder) WithState(state string) *AddressBuilder {
	b.state = state
	return b
}

func (b *AddressBuilder) WithPostalCode(postalCode string) *AddressBuilder {
	b.postalCode = postalCode
	return b
}

func (b *AddressBuilder) WithCountry(country string) *AddressBuilder {
	b.country = country
	return b
}

func (b *AddressBuilder) WithCreatedAt(createdAt time.Time) *AddressBuilder {
	b.createdAt = createdAt
	return b
}

func (b *AddressBuilder) WithUpdatedAt(updatedAt time.Time) *AddressBuilder {
	b.updatedAt = updatedAt
	return b
}
