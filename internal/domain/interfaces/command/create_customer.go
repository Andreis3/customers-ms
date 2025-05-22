package command

import (
	"context"

	"github.com/andreis3/customers-ms/internal/domain/aggregate"
	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
	"github.com/andreis3/customers-ms/internal/domain/entity/customer"
)

type CreateCustomer interface {
	Execute(ctx context.Context, input aggregate.CustomerProfile) (*customer.Customer, *apperror.Error)
}
