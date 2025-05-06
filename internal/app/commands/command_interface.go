package commands

import (
	"context"

	"github.com/andreis3/customers-ms/internal/domain/aggregate"
	"github.com/andreis3/customers-ms/internal/domain/apperrors"
	"github.com/andreis3/customers-ms/internal/domain/entity/customer"
)

type ICreateCustomer interface {
	Execute(ctx context.Context, input aggregate.CustomerProfile) (*customer.Customer, *apperrors.AppErrors)
}
