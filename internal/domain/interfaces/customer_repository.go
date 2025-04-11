package interfaces

import (
	"context"

	"github.com/andreis3/users-ms/internal/domain/entity"
	"github.com/andreis3/users-ms/internal/domain/errors"
)

type CustomerRepository interface {
	InsertCustomer(ctx context.Context, data entity.Customer) (*entity.Customer, *errors.AppErrors)
}
