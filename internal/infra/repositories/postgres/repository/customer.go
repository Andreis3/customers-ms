package repository

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"

	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
	"github.com/andreis3/customers-ms/internal/domain/entity/customer"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/infra/adapters/observability"
	"github.com/andreis3/customers-ms/internal/infra/repositories/postgres/model"
)

type CustomerRepository struct {
	DB      adapter.InstructionPostgres
	metrics adapter.Prometheus
	model.Customer
}

func NewCustomerRepository(
	db adapter.InstructionPostgres,
	metrics adapter.Prometheus,
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
	modelCustomer := c.FromModel(data)
	const query = `
	INSERT INTO customers 
	(email, password, first_name, last_name, cpf, date_of_birth, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	RETURNING id`

	var id int64

	err := c.DB.QueryRow(ctx, query,
		modelCustomer.Email,
		modelCustomer.Password,
		modelCustomer.FirstName,
		modelCustomer.LastName,
		modelCustomer.CPF,
		modelCustomer.DateOfBirth,
		modelCustomer.CreatedAT,
		modelCustomer.UpdatedAT).Scan(&id)
	if err != nil {
		return nil, apperror.ErrorSaveCustomer(err)
	}

	modelCustomer.ID = &id
	result := modelCustomer.ToEntity()

	span.SetAttributes(
		attribute.Int64("customer_id", id),
	)

	return &result, nil
}

func (c *CustomerRepository) FindCustomerByEmail(ctx context.Context, email string) (*customer.Customer, *apperror.Error) {
	ctx, span := observability.Tracer.Start(ctx, "CustomerRepository.FindCustomerByEmail")
	start := time.Now()

	defer func() {
		end := time.Since(start)
		c.metrics.ObserveInstructionDBDuration("postgres", "customers", "select", float64(end.Milliseconds()))
		span.End()
	}()

	const query = `
	SELECT id, email, password, first_name, last_name, cpf, date_of_birth, created_at, updated_at
	FROM customers
	WHERE email = $1`

	var modelCustomer model.Customer

	rows, err := c.DB.Query(ctx, query, email)
	if err != nil {
		return nil, apperror.ErrorFindCustomerByEmail(err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}
	err = rows.Scan(
		&modelCustomer.ID,
		&modelCustomer.Email,
		&modelCustomer.Password,
		&modelCustomer.FirstName,
		&modelCustomer.LastName,
		&modelCustomer.CPF,
		&modelCustomer.DateOfBirth,
		&modelCustomer.CreatedAT,
		&modelCustomer.UpdatedAT,
	)
	if err != nil {
		return nil, apperror.ErrorFindCustomerByEmail(err)
	}

	result := modelCustomer.ToEntity()

	span.SetAttributes(
		attribute.Int64("customer_id", *modelCustomer.ID),
	)

	return &result, nil
}
