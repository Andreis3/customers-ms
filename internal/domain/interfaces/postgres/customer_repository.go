package postgres

import (
	"context"

	"github.com/andreis3/customers-ms/internal/domain/entity/customer"
	"github.com/andreis3/customers-ms/internal/domain/errors"
)

type CustomerRepository interface {
	InsertCustomer(ctx context.Context, data customer.Customer) (*customer.Customer, *errors.Error)
	FindCustomerByEmail(ctx context.Context, email string) (*customer.Customer, *errors.Error)
}
