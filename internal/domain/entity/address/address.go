package address

import (
	"time"

	"github.com/andreis3/customers-ms/internal/domain/validator"
)

type Address struct {
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

func (a *Address) AssignID(id int64) {
	a.id = id
}

func (a *Address) ID() int64 {
	return a.id
}

func (a *Address) CustomerID() int64 {
	return a.customerID
}

func (a *Address) Street() string {
	return a.street
}

func (a *Address) Number() string {
	return a.number
}

func (a *Address) Complement() string {
	return a.complement
}

func (a *Address) City() string {
	return a.city
}

func (a *Address) State() string {
	return a.state
}

func (a *Address) PostalCode() string {
	return a.postalCode
}

func (a *Address) Country() string {
	return a.country
}

func (a *Address) CreatedAt() time.Time {
	return a.createdAt
}

func (a *Address) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a *Address) Validate() *validator.Validator {
	v := validator.New()
	v.Assert(validator.NotBlank(a.street), "street", validator.ErrNotBlank)
	v.Assert(validator.NotBlank(a.number), "number", validator.ErrNotBlank)
	v.Assert(validator.NotBlank(a.city), "city", validator.ErrNotBlank)
	v.Assert(validator.NotBlank(a.state), "state", validator.ErrNotBlank)
	v.Assert(validator.NotBlank(a.postalCode), "postal_code", validator.ErrNotBlank)
	v.Assert(validator.NotBlank(a.country), "country", validator.ErrNotBlank)
	return v
}
