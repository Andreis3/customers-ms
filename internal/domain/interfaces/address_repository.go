package interfaces

import (
	"context"

	"github.com/andreis3/users-ms/internal/domain/apperrors"
	"github.com/andreis3/users-ms/internal/domain/entity/address"
)

type AddressRepository interface {
	InsertBatchAddress(ctx context.Context, customerID int64, addresses []address.Address) (*[]address.Address, *apperrors.AppErrors)
}
