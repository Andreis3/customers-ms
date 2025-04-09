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

func EntityToModel(entity entity.Customer) *Customer {
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

func ModelToEntity(model Customer) entity.Customer {
	return entity.Customer{
		ID:          pointers.ToInt64(model.ID),
		Email:       pointers.ToString(model.Email),
		Password:    pointers.ToString(model.Password),
		FirstName:   pointers.ToString(model.FirstName),
		LastName:    pointers.ToString(model.LastName),
		CPF:         pointers.ToString(model.CPF),
		DateOfBirth: pointers.ToTime(model.DateOfBirth),
		CreatedAT:   pointers.ToTime(model.CreatedAT),
		UpdatedAT:   pointers.ToTime(model.UpdatedAT),
	}
}
