package customer

import (
	"time"

	"github.com/andreis3/users-ms/internal/domain/entity/address"
	valueobject "github.com/andreis3/users-ms/internal/domain/value-object"
)

type CustomerBuilder struct {
	id          int64
	email       string
	password    string
	firstName   string
	lastName    string
	cpf         string
	dateOfBirth time.Time
	addresses   []address.Address
	createdAt   time.Time
	updatedAt   time.Time
}

func NewBuilder() *CustomerBuilder {
	return &CustomerBuilder{}
}

func (b *CustomerBuilder) Build() *Customer {
	customer := &Customer{
		id:          b.id,
		email:       valueobject.NewEmail(b.email),
		password:    valueobject.NewPassword(b.password),
		firstName:   b.firstName,
		lastName:    b.lastName,
		cpf:         valueobject.NewCPF(b.cpf),
		dateOfBirth: b.dateOfBirth,
		addresses:   b.addresses,
		createdAt:   b.createdAt,
		updatedAt:   b.updatedAt,
	}

	return customer
}

func (b *CustomerBuilder) WithID(id int64) *CustomerBuilder {
	b.id = id
	return b
}

func (b *CustomerBuilder) WithEmail(email string) *CustomerBuilder {
	b.email = email
	return b
}

func (b *CustomerBuilder) WithPassword(password string) *CustomerBuilder {
	b.password = password
	return b
}

func (b *CustomerBuilder) WithFirstName(firstName string) *CustomerBuilder {
	b.firstName = firstName
	return b
}

func (b *CustomerBuilder) WithLastName(lastName string) *CustomerBuilder {
	b.lastName = lastName
	return b
}

func (b *CustomerBuilder) WithCPF(cpf string) *CustomerBuilder {
	b.cpf = cpf
	return b
}

func (b *CustomerBuilder) WithDateOfBirth(dateOfBirth time.Time) *CustomerBuilder {
	b.dateOfBirth = dateOfBirth
	return b
}

func (b *CustomerBuilder) WithAddresses(addresses []address.Address) *CustomerBuilder {
	b.addresses = addresses
	return b
}

func (b *CustomerBuilder) WithCreatedAt(createdAt time.Time) *CustomerBuilder {
	b.createdAt = createdAt
	return b
}

func (b *CustomerBuilder) WithUpdatedAt(updatedAt time.Time) *CustomerBuilder {
	b.updatedAt = updatedAt
	return b
}
