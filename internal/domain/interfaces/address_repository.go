package interfaces

import (
	"context"

	"github.com/andreis3/users-ms/internal/domain/entity/address"
	"github.com/andreis3/users-ms/internal/domain/errors"
)

type AddressRepository interface {
	InsertBatchAddress(ctx context.Context, customerID int64, addresses []address.Address) (*[]address.Address, *errors.AppErrors)
}
