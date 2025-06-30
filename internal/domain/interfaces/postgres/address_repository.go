package postgres

import (
	"context"

	"github.com/andreis3/customers-ms/internal/domain/entity/address"
	"github.com/andreis3/customers-ms/internal/domain/error"
)

type AddressRepository interface {
	InsertBatchAddress(ctx context.Context, customerID int64, addresses []address.Address) (*[]address.Address, *error.Error)
}
