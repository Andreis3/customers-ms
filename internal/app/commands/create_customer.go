package commands

import (
	"context"
	"log/slog"

	"github.com/andreis3/customers-ms/internal/domain/aggregate"
	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
	"github.com/andreis3/customers-ms/internal/domain/entity/customer"
	"github.com/andreis3/customers-ms/internal/domain/interfaces"
	"github.com/andreis3/customers-ms/internal/infra/adapters/observability"
)

type CreateCustomerCommand struct {
	uow             interfaces.UnitOfWork
	bcrypt          interfaces.Bcrypt
	log             interfaces.Logger
	customerService interfaces.CustomerService
}

func NewCreateCustomer(
	uow interfaces.UnitOfWork,
	bcrypt interfaces.Bcrypt,
	log interfaces.Logger,
	customerService interfaces.CustomerService,
) *CreateCustomerCommand {
	return &CreateCustomerCommand{
		uow:             uow,
		bcrypt:          bcrypt,
		log:             log,
		customerService: customerService,
	}
}

func (c *CreateCustomerCommand) Execute(ctx context.Context, input aggregate.CustomerProfile) (*customer.Customer, *apperror.Error) {
	ctx, child := observability.Tracer.Start(ctx, "CreatedCustomer.Execute")
	defer child.End()
	traceID := child.SpanContext().TraceID().String()
	c.log.InfoText("Received input to create customer",
		slog.String("trace_id", traceID),
		slog.Any("input", input))

	err := input.Validate()
	if err != nil {
		child.RecordError(err)
		c.log.ErrorJSON("Failed validate input to create customer",
			slog.String("trace_id", traceID),
			slog.Any("error", err.Errors))
		return nil, err
	}

	customerAlreadyExists := c.customerService.ExistCustomerByEmail(ctx, input.Customer.Email())
	if customerAlreadyExists {
		child.RecordError(apperror.ErrCustomerAlreadyExists())
		c.log.ErrorJSON("Customer already exists",
			slog.String("trace_id", traceID),
			slog.Any("error", apperror.ErrCustomerAlreadyExists))
		return nil, apperror.ErrCustomerAlreadyExists()
	}

	var customer *customer.Customer

	errUow := c.uow.Do(func(uow interfaces.UnitOfWork) *apperror.Error {
		hash, err := c.bcrypt.Hash(input.Customer.Password())
		if err != nil {
			child.RecordError(err)
			c.log.ErrorJSON("Failed hash password",
				slog.String("trace_id", traceID),
				slog.Any("error", err))
			return err
		}
		input.Customer.AssignHashedPassword(hash)

		customerRepository := uow.CustomerRepository()

		customer, err = customerRepository.InsertCustomer(ctx, input.Customer)
		if err != nil {
			child.RecordError(err)
			c.log.ErrorJSON("Failed insert customer",
				slog.String("trace_id", traceID),
				slog.Any("error", err))
			return err
		}

		if len(input.Addresses) > 0 {
			addressRepository := uow.AddressRepository()
			_, err = addressRepository.InsertBatchAddress(ctx, customer.ID(), input.Addresses)
			if err != nil {
				child.RecordError(err)
				c.log.ErrorJSON("Failed insert address",
					slog.String("trace_id", traceID),
					slog.Any("error", err))
				return err
			}
		}
		return nil
	})

	if errUow != nil {
		child.RecordError(errUow)
		c.log.ErrorJSON("Failed insert customer",
			slog.String("trace_id", traceID),
			slog.Any("error", errUow))
		return nil, errUow
	}

	return customer, nil
}
