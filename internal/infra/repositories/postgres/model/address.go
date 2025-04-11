package model

import (
	"time"

	"github.com/andreis3/users-ms/internal/domain/entity"
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

func (a Address) ToEntity() entity.Address {
	return entity.Address{
		ID:         pointers.ToInt64(a.ID),
		CustomerID: pointers.ToInt64(a.CustomerID),
		Street:     pointers.ToString(a.Street),
		Number:     pointers.ToString(a.Number),
		Complement: pointers.ToString(a.Complement),
		City:       pointers.ToString(a.City),
		State:      pointers.ToString(a.State),
		PostalCode: pointers.ToString(a.PostalCode),
		Country:    pointers.ToString(a.Country),
		CreatedAt:  pointers.ToTime(a.CreatedAt),
		UpdatedAt:  pointers.ToTime(a.UpdatedAt),
	}
}

func (a Address) FromModel(entity entity.Address) *Address {
	return &Address{
		ID:         &entity.ID,
		CustomerID: &entity.CustomerID,
		Street:     &entity.Street,
		Number:     &entity.Number,
		Complement: &entity.Complement,
		City:       &entity.City,
		State:      &entity.State,
		PostalCode: &entity.PostalCode,
		Country:    &entity.Country,
		CreatedAt:  &entity.CreatedAt,
		UpdatedAt:  &entity.UpdatedAt,
	}
}
