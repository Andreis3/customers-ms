package entity

import (
	"time"

	"github.com/andreis3/users-ms/internal/domain/validator"
)

type Customer struct {
	ID          string
	Email       string
	Password    string
	FirstName   string
	LastName    string
	CPF         string
	DateOfBirth time.Time
	CreatedAT   time.Time
	UpdatedAT   time.Time
	Addresses   []Address
	validator.Validator
}

func CustomerBuilder() *Customer {
	return &Customer{}
}

func (c *Customer) SetID(id string) *Customer {
	c.ID = id
	return c
}

func (c *Customer) SetEmail(email string) *Customer {
	c.Email = email
	return c
}

func (c *Customer) SetPassword(password string) *Customer {
	c.Password = password
	return c
}

func (c *Customer) SetFirstName(firstName string) *Customer {
	c.FirstName = firstName
	return c
}

func (c *Customer) SetLasName(lastName string) *Customer {
	c.LastName = lastName
	return c
}

func (c *Customer) SetCPF(cpf string) *Customer {
	c.CPF = cpf
	return c
}

func (c *Customer) SetDateOfBirth(dateOfBirth time.Time) *Customer {
	c.DateOfBirth = dateOfBirth
	return c
}

func (c *Customer) SetAddresses(addresses []Address) *Customer {
	c.Addresses = addresses
	return c
}

func (c *Customer) SetCreatedAT(createAt time.Time) *Customer {
	c.CreatedAT = createAt
	return c
}

func (c *Customer) SetUpdatedAT(updatedAt time.Time) *Customer {
	c.UpdatedAT = updatedAt
	return c
}

func (c *Customer) Build() *Customer {
	return c
}

func (c *Customer) Validate() *validator.Validator {
	c.Assert(validator.NotBlank(c.Email), "email", validator.NotBlankField)
	c.Assert(validator.NotBlank(c.Password), "password", validator.NotBlankField)
	c.Assert(validator.NotBlank(c.FirstName), "first_name", validator.NotBlankField)
	c.Assert(validator.NotBlank(c.LastName), "last_name", validator.NotBlankField)
	c.Assert(validator.NotBlank(c.CPF), "last_name", validator.NotBlankField)
	c.Assert(validator.NotBlank(c.DateOfBirth.GoString()), "date_of_birth", validator.NotBlankField)

	return &c.Validator
}
