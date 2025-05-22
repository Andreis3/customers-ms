package repository

import (
	"context"

	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
	"github.com/andreis3/customers-ms/internal/domain/entity/customer"
)

type CustomerRepository interface {
	InsertCustomer(ctx context.Context, data customer.Customer) (*customer.Customer, *apperror.Error)
	FindCustomerByEmail(ctx context.Context, email string) (*customer.Customer, *apperror.Error)
}
