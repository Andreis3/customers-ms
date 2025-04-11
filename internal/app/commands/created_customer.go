package commands

import (
	"context"

	"github.com/andreis3/users-ms/internal/domain/aggregate"
	"github.com/andreis3/users-ms/internal/domain/entity"
	"github.com/andreis3/users-ms/internal/domain/errors"
	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/adapters/observability"
	"github.com/andreis3/users-ms/internal/infra/factories"
)

type CreatedCustomerCommand struct {
	uow interfaces.UnitOfWork
}

func NewCreatedCustomer(uow interfaces.UnitOfWork) *CreatedCustomerCommand {
	return &CreatedCustomerCommand{
		uow: uow,
	}
}

func (c *CreatedCustomerCommand) Execute(ctx context.Context, data aggregate.CustomerProfile) (*entity.Customer, *errors.AppErrors) {
	ctx, span := observability.Tracer.Start(ctx, "CreatedCustomer.Execute")
	defer span.End()
	err := data.Validate()
	if err != nil {
		return nil, err
	}

	customerResult := &entity.Customer{}

	errUow := c.uow.Do(func(unitOfWork interfaces.UnitOfWork) *errors.AppErrors {
		repo, err := factories.LoadCustomerFactory(unitOfWork)
		if err != nil {
			span.RecordError(err)
			return err
		}

		customer, err := repo.CustomerRepo.InsertCustomer(ctx, data.Customer)
		if err != nil {
			return err
		}
		if len(data.Addresses) > 0 {
			_, err = repo.AddressRepo.InsertBatchAddress(ctx, customer.ID, data.Addresses)
			if err != nil {
				span.RecordError(err)
				return err
			}
		}
		customerResult = customer
		return nil
	})
	if errUow != nil {
		span.RecordError(errUow)
		return nil, errUow
	}
	return customerResult, nil
}
