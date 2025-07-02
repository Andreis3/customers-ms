package uow

import (
	"context"

	"github.com/andreis3/customers-ms/internal/domain/errors"
)

type RepositoryFactory func(tx any) any

type UnitOfWork interface {
	Do(ctx context.Context, fn func(ctx context.Context) *errors.Error) *errors.Error
}
