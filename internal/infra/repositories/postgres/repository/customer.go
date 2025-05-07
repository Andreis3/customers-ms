package repository

import (
	"context"
	"time"

	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
	"github.com/andreis3/customers-ms/internal/domain/entity/customer"
	"github.com/andreis3/customers-ms/internal/domain/interfaces"
	"github.com/andreis3/customers-ms/internal/infra/adapters/observability"
	"github.com/andreis3/customers-ms/internal/infra/repositories/postgres/model"
	"go.opentelemetry.io/otel/attribute"
)

type CustomerRepository struct {
	DB      interfaces.InstructionPostgres
	metrics interfaces.Prometheus
	model.Customer
}

func NewCustomerRepository(
	db interfaces.InstructionPostgres,
	metrics interfaces.Prometheus,
) *CustomerRepository {
	return &CustomerRepository{
		DB:      db,
		metrics: metrics,
	}
}

func (c *CustomerRepository) InsertCustomer(ctx context.Context, data customer.Customer) (*customer.Customer, *apperror.Error) {
	ctx, span := observability.Tracer.Start(ctx, "CustomerRepository.InsertCustomer")
	start := time.Now()
	defer func() {
		end := time.Since(start)
		c.metrics.ObserveInstructionDBDuration("postgres", "customers", "insert", float64(end.Milliseconds()))
		span.End()
	}()
	model := c.FromModel(data)
	const query = `-- name: InsertCustomer :one
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
		return nil, apperror.ErrorSaveCustomer(err)
	}

	model.ID = &id
	result := model.ToEntity()

	span.SetAttributes(
		attribute.Int64("customer_id", id),
	)

	return &result, nil
}
