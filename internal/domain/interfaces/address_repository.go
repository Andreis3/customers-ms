package interfaces

import (
	"context"

	"github.com/andreis3/users-ms/internal/domain/entity"
	"github.com/andreis3/users-ms/internal/domain/errors"
)

type AddressRepository interface {
	InsertBatchAddress(ctx context.Context, customerID int64, addresses []entity.Address) (*[]entity.Address, *errors.AppErrors)
}
