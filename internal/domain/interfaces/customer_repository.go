package interfaces

import (
	"context"

	"github.com/andreis3/users-ms/internal/domain/apperrors"
	"github.com/andreis3/users-ms/internal/domain/entity/customer"
)

type CustomerRepository interface {
	InsertCustomer(ctx context.Context, data customer.Customer) (*customer.Customer, *apperrors.AppErrors)
}
