package command

import (
	"context"

	"github.com/andreis3/customers-ms/internal/domain/aggregate"
	"github.com/andreis3/customers-ms/internal/domain/entity"
	"github.com/andreis3/customers-ms/internal/domain/errors"
)

type CreateCustomer interface {
	Execute(ctx context.Context, input aggregate.CustomerProfile) (*entity.Customer, *errors.Error)
}
