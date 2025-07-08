package commands

import (
	"context"
	"log/slog"

	"github.com/andreis3/customers-ms/internal/domain/aggregate"
	"github.com/andreis3/customers-ms/internal/domain/entity"
	"github.com/andreis3/customers-ms/internal/domain/errors"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/commons"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/postgres"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/service"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/uow"
)

type CreateCustomerCommand struct {
	uow                func(ctx context.Context) uow.UnitOfWork
	bcrypt             adapter.Bcrypt
	log                commons.Logger
	customerRepository postgres.CustomerRepository
	addressRepository  postgres.AddressRepository
	customerService    service.CustomerService
	tracer             adapter.Tracer
}

func NewCreateCustomer(
	uow func(ctx context.Context) uow.UnitOfWork,
	bcrypt adapter.Bcrypt,
	log commons.Logger,
	customerService service.CustomerService,
	customerRepository postgres.CustomerRepository,
	addressRepository postgres.AddressRepository,
	tracer adapter.Tracer,
) *CreateCustomerCommand {
	return &CreateCustomerCommand{
		uow:                uow,
		bcrypt:             bcrypt,
		log:                log,
		customerService:    customerService,
		customerRepository: customerRepository,
		addressRepository:  addressRepository,
		tracer:             tracer,
	}
}

func (c *CreateCustomerCommand) Execute(ctx context.Context, input aggregate.CustomerProfile) (*entity.Customer, *errors.Error) {
	ctx, span := c.tracer.Start(ctx, "CreatedCustomer.Execute")
	defer span.End()
	traceID := span.SpanContext().TraceID()
	c.log.InfoJSON("Received input to create CreateCustomer",
		slog.String("trace_id", traceID),
		slog.Any("input", input))

	err := input.Validate()
	if err != nil {
		span.RecordError(err)
		c.log.ErrorJSON("Failed validate input to create CreateCustomer",
			slog.String("trace_id", traceID),
			slog.Any("error", err.Errors))
		return nil, err
	}

	customerAlreadyExists := c.customerService.ExistCustomerByEmail(ctx, input.Customer.Email())
	if customerAlreadyExists {
		span.RecordError(errors.ErrCustomerAlreadyExists())
		c.log.ErrorJSON("Customer already exists",
			slog.String("trace_id", traceID),
			slog.Any("error", errors.ErrCustomerAlreadyExists))
		return nil, errors.ErrCustomerAlreadyExists()
	}

	var customerResult *entity.Customer
	uowInstance := c.uow(ctx)

	errUow := uowInstance.Do(ctx, func(ctxUow context.Context) *errors.Error {
		hash, err := c.bcrypt.Hash(input.Customer.Password())
		if err != nil {
			span.RecordError(err)
			c.log.ErrorJSON("Failed hash password",
				slog.String("trace_id", traceID),
				slog.Any("error", err))
			return err
		}

		input.Customer.AssignHashedPassword(hash)

		customerResult, err = c.customerRepository.InsertCustomer(ctxUow, input.Customer)
		if err != nil {
			span.RecordError(err)
			c.log.ErrorJSON("Failed insert customerResult",
				slog.String("trace_id", traceID),
				slog.Any("error", err))
			return err
		}

		if len(input.Addresses) > 0 {
			_, err = c.addressRepository.InsertBatchAddress(ctxUow, customerResult.ID(), input.Addresses)
			if err != nil {
				span.RecordError(err)
				c.log.ErrorJSON("Failed insert address",
					slog.String("trace_id", traceID),
					slog.Any("error", err))
				return err
			}
		}
		return nil
	})

	if errUow != nil {
		span.RecordError(errUow)
		c.log.ErrorJSON("Failed insert customerResult",
			slog.String("trace_id", traceID),
			slog.Any("error", errUow))
		return nil, errUow
	}

	return customerResult, nil
}
