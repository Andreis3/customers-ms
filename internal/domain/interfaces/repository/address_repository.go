package repository

import (
	"context"

	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
	"github.com/andreis3/customers-ms/internal/domain/entity/address"
)

type AddressRepository interface {
	InsertBatchAddress(ctx context.Context, customerID int64, addresses []address.Address) (*[]address.Address, *apperror.Error)
}
