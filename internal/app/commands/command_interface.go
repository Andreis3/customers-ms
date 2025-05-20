package commands

import (
	"context"

	"github.com/andreis3/customers-ms/internal/domain/aggregate"
	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
	"github.com/andreis3/customers-ms/internal/domain/entity/customer"
)

type ICreateCustomer interface {
	Execute(ctx context.Context, input aggregate.CustomerProfile) (*customer.Customer, *apperror.Error)
}

type IAuthenticateCustomer interface {
	Execute(ctx context.Context, input AuthenticateCustomerInput) (*AuthenticateCustomerOutput, *apperror.Error)
}
