package model

import (
	"time"

	"github.com/andreis3/users-ms/internal/domain/entity/address"
	"github.com/andreis3/users-ms/internal/infra/commons/pointers"
)

type Address struct {
	ID         *int64
	CustomerID *int64
	Street     *string
	Number     *string
	Complement *string
	City       *string
	State      *string
	PostalCode *string
	Country    *string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

func (a Address) ToEntity() *address.Address {
	return address.NewBuilder().
		WithID(pointers.ToInt64(a.ID)).
		WithCustomerID(pointers.ToInt64(a.CustomerID)).
		WithStreet(pointers.ToString(a.Street)).
		WithNumber(pointers.ToString(a.Number)).
		WithComplement(pointers.ToString(a.Complement)).
		WithCity(pointers.ToString(a.City)).
		WithState(pointers.ToString(a.State)).
		WithPostalCode(pointers.ToString(a.PostalCode)).
		WithCountry(pointers.ToString(a.Country)).
		WithCreatedAt(pointers.ToTime(a.CreatedAt)).
		WithUpdatedAt(pointers.ToTime(a.UpdatedAt)).
		Build()
}

func (a Address) FromModel(entity address.Address) *Address {
	return &Address{
		CustomerID: pointers.ToInt64Pointer(entity.CustomerID()),
		Street:     pointers.ToStringPointer(entity.Street()),
		Number:     pointers.ToStringPointer(entity.Number()),
		Complement: pointers.ToStringPointer(entity.Complement()),
		City:       pointers.ToStringPointer(entity.City()),
		State:      pointers.ToStringPointer(entity.State()),
		PostalCode: pointers.ToStringPointer(entity.PostalCode()),
		Country:    pointers.ToStringPointer(entity.Country()),
		CreatedAt:  pointers.ToTimePointer(entity.CreatedAt()),
		UpdatedAt:  pointers.ToTimePointer(entity.UpdatedAt()),
	}
}
