package repository

import (
	"context"
	"time"

	"github.com/andreis3/users-ms/internal/domain/entity/customer"
	"github.com/andreis3/users-ms/internal/domain/errors"
	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/adapters/db/postegres"
	"github.com/andreis3/users-ms/internal/infra/adapters/observability"
	infra_errors "github.com/andreis3/users-ms/internal/infra/commons/errors"
	"github.com/andreis3/users-ms/internal/infra/repositories/postgres/model"
	"go.opentelemetry.io/otel/attribute"
)

type CustomerRepository struct {
	DB      *postegres.Queries
	metrics interfaces.Prometheus
	model.Customer
}

func NewCustomerRepository(metrics interfaces.Prometheus) *CustomerRepository {
	return &CustomerRepository{
		metrics: metrics,
	}
}

func (c *CustomerRepository) InsertCustomer(ctx context.Context, data customer.Customer) (*customer.Customer, *errors.AppErrors) {
	ctx, span := observability.Tracer.Start(ctx, "CustomerRepository.InsertCustomer")
	start := time.Now()
	defer func() {
		end := time.Since(start)
		c.metrics.ObserveInstructionDBDuration("postgres", "customers", "insert", float64(end.Milliseconds()))
		span.End()
	}()
	model := c.FromModel(data)
	const query = `
	INSERT INTO customers 
	(email, password, first_name, last_name, cpf, date_of_birth, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	RETURNING id`

	var id int64

	err := c.DB.QueryRow(ctx, query,
		model.Email,
		model.Password,
		model.FirstName,
		model.LastName,
		model.CPF,
		model.DateOfBirth,
		model.CreatedAT,
		model.UpdatedAT).Scan(&id)
	if err != nil {
		return nil, infra_errors.ErrorSaveCustomer(err)
	}

	model.ID = &id
	result := model.ToEntity()

	span.SetAttributes(
		attribute.Int64("customer_id", id),
	)

	return result, nil
}
