package entity

import (
	"fmt"
	"slices"
	"time"

	"github.com/andreis3/customers-ms/internal/domain/validator"
	valueobject "github.com/andreis3/customers-ms/internal/domain/value-object"
)

type Customer struct {
	id          int64
	email       valueobject.Email `json:"email"`
	password    valueobject.Password
	firstName   string `json:"first_name"`
	lastName    string
	cpf         valueobject.CPF
	dateOfBirth time.Time
	addresses   []Address
	createdAt   time.Time
	updatedAt   time.Time
}

type CustomerProps struct {
	id          int64
	email       string
	password    string
	firstName   string
	lastName    string
	cpf         string
	dateOfBirth time.Time
	addresses   []Address
	createdAt   time.Time
	updatedAt   time.Time
}

func BuilderCustomer() *CustomerProps {
	return &CustomerProps{}
}

func (b *CustomerProps) Build() Customer {
	customer := Customer{
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

func (c *Customer) AssignID(id int64) {
	c.id = id
}

func (c *Customer) ID() int64 {
	return c.id
}

func (c *Customer) AssignHashedPassword(hashed string) {
	c.password = valueobject.NewEncryptedPassword(hashed)
}

func (c *Customer) FullName() string {
	return fmt.Sprintf("%s %s", c.firstName, c.lastName)
}

func (c *Customer) Email() string {
	return c.email.String()
}

func (c *Customer) Password() string {
	return c.password.String()
}

func (c *Customer) FirstName() string {
	return c.firstName
}

func (c *Customer) LastName() string {
	return c.lastName
}

func (c *Customer) CPF() string {
	return c.cpf.String()
}

func (c *Customer) DateOfBirth() time.Time {
	return c.dateOfBirth
}

func (c *Customer) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Customer) UpdatedAt() time.Time {
	return c.updatedAt
}

func (c *Customer) Addresses() []Address {
	return slices.Clone(c.addresses)
}

func (c *Customer) Validate() *validator.Validator {
	v := validator.New()
	v.Assert(validator.NotBlank(c.firstName), "first_name", validator.ErrNotBlank)
	v.Assert(validator.NotBlank(c.lastName), "last_name", validator.ErrNotBlank)
	v.Assert(!c.dateOfBirth.IsZero(), "date_of_birth", validator.ErrNotBlank)

	v.Merge(c.email.Validate())
	v.Merge(c.cpf.Validate())
	v.Merge(c.password.Validate())
	return v
}

func (b *CustomerProps) WithID(id int64) *CustomerProps {
	b.id = id
	return b
}

func (b *CustomerProps) WithEmail(email string) *CustomerProps {
	b.email = email
	return b
}

func (b *CustomerProps) WithPassword(password string) *CustomerProps {
	b.password = password
	return b
}

func (b *CustomerProps) WithFirstName(firstName string) *CustomerProps {
	b.firstName = firstName
	return b
}

func (b *CustomerProps) WithLastName(lastName string) *CustomerProps {
	b.lastName = lastName
	return b
}

func (b *CustomerProps) WithCPF(cpf string) *CustomerProps {
	b.cpf = cpf
	return b
}

func (b *CustomerProps) WithDateOfBirth(dateOfBirth time.Time) *CustomerProps {
	b.dateOfBirth = dateOfBirth
	return b
}

func (b *CustomerProps) WithAddresses(addresses []Address) *CustomerProps {
	b.addresses = addresses
	return b
}

func (b *CustomerProps) WithCreatedAt(createdAt time.Time) *CustomerProps {
	b.createdAt = createdAt
	return b
}

func (b *CustomerProps) WithUpdatedAt(updatedAt time.Time) *CustomerProps {
	b.updatedAt = updatedAt
	return b
}
