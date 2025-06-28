package model

import (
	"time"

	"github.com/andreis3/customers-ms/internal/domain/entity/customer"
	"github.com/andreis3/customers-ms/internal/infra/commons/pointers"
)

type Customer struct {
	ID          *int64
	Email       *string
	Password    *string
	FirstName   *string
	LastName    *string
	CPF         *string
	DateOfBirth *time.Time
	CreatedAT   *time.Time
	UpdatedAT   *time.Time
}

func (c Customer) ToEntity() customer.Customer {
	return customer.NewBuilder().
		WithID(pointers.ToInt64(c.ID)).
		WithEmail(pointers.ToString(c.Email)).
		WithPassword(pointers.ToString(c.Password)).
		WithFirstName(pointers.ToString(c.FirstName)).
		WithLastName(pointers.ToString(c.LastName)).
		WithCPF(pointers.ToString(c.CPF)).
		WithDateOfBirth(pointers.ToTime(c.DateOfBirth)).
		WithCreatedAt(pointers.ToTime(c.CreatedAT)).
		WithUpdatedAt(pointers.ToTime(c.UpdatedAT)).
		Build()
}

func (c Customer) FromModel(customer customer.Customer) *Customer {
	dateNow := time.Now()
	return &Customer{
		Email:       pointers.ToStringPointer(customer.Email()),
		Password:    pointers.ToStringPointer(customer.Password()),
		FirstName:   pointers.ToStringPointer(customer.FirstName()),
		LastName:    pointers.ToStringPointer(customer.LastName()),
		CPF:         pointers.ToStringPointer(customer.CPF()),
		DateOfBirth: pointers.ToTimePointer(customer.DateOfBirth()),
		CreatedAT:   pointers.ToTimePointer(dateNow),
		UpdatedAT:   pointers.ToTimePointer(dateNow),
	}
}
