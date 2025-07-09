package commands

import (
	"context"
	"log/slog"

	"github.com/andreis3/customers-ms/internal/app/dto"
	"github.com/andreis3/customers-ms/internal/app/mapper"
	"github.com/andreis3/customers-ms/internal/domain/entity"
	"github.com/andreis3/customers-ms/internal/domain/errors"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/postgres"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/service"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/uow"
)

type CreateCustomerCommand struct {
	uow                func(ctx context.Context) uow.UnitOfWork
	bcrypt             adapter.Bcrypt
	log                adapter.Logger
	customerRepository postgres.CustomerRepository
	addressRepository  postgres.AddressRepository
	customerService    service.CustomerService
	tracer             adapter.Tracer
}

func NewCreateCustomer(
	uow func(ctx context.Context) uow.UnitOfWork,
	bcrypt adapter.Bcrypt,
	log adapter.Logger,
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

func (c *CreateCustomerCommand) Execute(ctx context.Context, input dto.CreateCustomerInput) (*entity.Customer, *errors.Error) {
	ctx, span := c.tracer.Start(ctx, "CreatedCustomer.Execute")
	customerProfile := mapper.ToCustomerProfile(input)
	defer span.End()
	traceID := span.SpanContext().TraceID()
	c.log.InfoJSON("Received input to create CreateCustomer",
		slog.String("trace_id", traceID),
		slog.Any("input", input))

	err := customerProfile.Validate()
	if err != nil {
		span.RecordError(err)
		c.log.ErrorJSON("Failed validate input to create CreateCustomer",
			slog.String("trace_id", traceID),
			slog.Any("error", err.Errors))
		return nil, err
	}

	//customerAlreadyExists := c.customerService.ExistCustomerByEmail(ctx, customerProfile.Customer.Email())
	//if customerAlreadyExists {
	//	span.RecordError(errors.ErrCustomerAlreadyExists())
	//	c.log.ErrorJSON("Customer already exists",
	//		slog.String("trace_id", traceID),
	//		slog.Any("error", errors.ErrCustomerAlreadyExists))
	//	return nil, errors.ErrCustomerAlreadyExists()
	//}

	var customerResult *entity.Customer
	uowInstance := c.uow(ctx)

	errUow := uowInstance.Do(ctx, func(ctxUow context.Context) *errors.Error {
		hash, err := c.bcrypt.Hash(customerProfile.Customer.Password())
		if err != nil {
			span.RecordError(err)
			c.log.ErrorJSON("Failed hash password",
				slog.String("trace_id", traceID),
				slog.Any("error", err))
			return err
		}

		customerProfile.Customer.AssignHashedPassword(hash)

		customerResult, err = c.customerRepository.InsertCustomer(ctxUow, customerProfile.Customer)
		if err != nil {
			span.RecordError(err)
			c.log.ErrorJSON("Failed insert customerResult",
				slog.String("trace_id", traceID),
				slog.Any("error", err))
			return err
		}

		if len(customerProfile.Addresses) > 0 {
			_, err = c.addressRepository.InsertBatchAddress(ctxUow, customerResult.ID(), customerProfile.Addresses)
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
