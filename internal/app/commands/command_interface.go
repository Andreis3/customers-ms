package commands

import (
	"context"

	"github.com/andreis3/users-ms/internal/domain/aggregate"
	"github.com/andreis3/users-ms/internal/domain/apperrors"
	"github.com/andreis3/users-ms/internal/domain/entity/customer"
)

type ICreateCustomer interface {
	Execute(ctx context.Context, input aggregate.CustomerProfile) (*customer.Customer, *apperrors.AppErrors)
}
