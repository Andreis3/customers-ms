package model

import (
	"time"

	"github.com/andreis3/users-ms/internal/domain/entity"
	"github.com/andreis3/users-ms/internal/infra/commons/pointers"
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

func (c Customer) ToEntity() entity.Customer {
	return entity.Customer{
		ID:          pointers.ToInt64(c.ID),
		Email:       pointers.ToString(c.Email),
		Password:    pointers.ToString(c.Password),
		FirstName:   pointers.ToString(c.FirstName),
		LastName:    pointers.ToString(c.LastName),
		CPF:         pointers.ToString(c.CPF),
		DateOfBirth: pointers.ToTime(c.DateOfBirth),
		CreatedAT:   pointers.ToTime(c.CreatedAT),
		UpdatedAT:   pointers.ToTime(c.UpdatedAT),
	}
}

func (c Customer) FromModel(entity entity.Customer) *Customer {
	return &Customer{
		Email:       &entity.Email,
		Password:    &entity.Password,
		FirstName:   &entity.FirstName,
		LastName:    &entity.LastName,
		CPF:         &entity.CPF,
		DateOfBirth: &entity.DateOfBirth,
		CreatedAT:   &entity.CreatedAT,
		UpdatedAT:   &entity.UpdatedAT,
	}
}
