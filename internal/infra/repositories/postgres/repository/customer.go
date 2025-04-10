package repository

import (
	"context"
	"time"

	"github.com/andreis3/users-ms/internal/domain/entity"
	"github.com/andreis3/users-ms/internal/domain/errors"
	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/adapters/db/postegres"
	infra_errors "github.com/andreis3/users-ms/internal/infra/commons/errors"
	"github.com/andreis3/users-ms/internal/infra/repositories/postgres/model"
)

type CustomerRepository struct {
	DB      *postegres.Queries
	metrics interfaces.Prometheus
}

func NewCustomerRepository(metrics interfaces.Prometheus) *CustomerRepository {
	return &CustomerRepository{
		metrics: metrics,
	}
}

func (c *CustomerRepository) InsertCustomer(ctx context.Context, data entity.Customer) (*entity.Customer, *errors.AppErrors) {
	start := time.Now()
	defer func() {
		end := time.Since(start)
		c.metrics.ObserveInstructionDBDuration("postgres", "customers", "insert", float64(end.Milliseconds()))
	}()
	customer := model.EntityToModel(data)
	const query = `
	INSERT INTO customers 
	(email, password, first_name, last_name, cpf, date_of_birth, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	RETURNING id`

	var id int64

	err := c.DB.QueryRow(ctx, query,
		customer.Email,
		customer.Password,
		customer.FirstName,
		customer.LastName,
		customer.CPF,
		customer.DateOfBirth,
		customer.CreatedAT,
		customer.UpdatedAT).Scan(&id)
	if err != nil {
		return nil, infra_errors.ErrorSaveCustomer(err)
	}

	result := data
	result.ID = id

	return &result, nil
}
