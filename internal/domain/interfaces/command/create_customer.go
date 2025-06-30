package command

import (
	"context"

	"github.com/andreis3/customers-ms/internal/domain/aggregate"
	"github.com/andreis3/customers-ms/internal/domain/entity/customer"
	"github.com/andreis3/customers-ms/internal/domain/error"
)

type CreateCustomer interface {
	Execute(ctx context.Context, input aggregate.CustomerProfile) (*customer.Customer, *error.Error)
}
