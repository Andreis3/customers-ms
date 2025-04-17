package interfaces

import (
	"context"

	"github.com/andreis3/users-ms/internal/domain/entity/customer"
	"github.com/andreis3/users-ms/internal/domain/errors"
)

type CustomerRepository interface {
	InsertCustomer(ctx context.Context, data *customer.Customer) (*customer.Customer, *errors.AppErrors)
}
