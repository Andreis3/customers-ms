package uow

import (
	"context"

	"github.com/andreis3/users-ms/internal/domain/apperrors"
	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/commons/infraerrors"
	"github.com/andreis3/users-ms/internal/infra/repositories/postgres/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	InternalServerError = "internal server error"
)

type UnitOfWork struct {
	DB         *pgxpool.Pool
	TX         pgx.Tx
	prometheus interfaces.Prometheus
}

func NewUnitOfWork(db *pgxpool.Pool, prometheus interfaces.Prometheus) *UnitOfWork {
	return &UnitOfWork{
		DB:         db,
		prometheus: prometheus,
	}
}

// Do handles transaction lifecycle safely.
func (u *UnitOfWork) Do(fn func(uow interfaces.UnitOfWork) *apperrors.AppErrors) *apperrors.AppErrors {
	ctx := context.Background()
	if u.TX != nil {
		return infraerrors.ErrorTransactionAlreadyExists()
	}

	tx, err := u.DB.Begin(ctx)
	if err != nil {
		return infraerrors.ErrorOpeningTransaction(err)
	}
	u.TX = tx
	defer func() { u.TX = nil }()

	if err := fn(u); err != nil {
		rollbackErr := u.TX.Rollback(ctx)
		if rollbackErr != nil {
			return infraerrors.ErrorExecuteRollback(rollbackErr)
		}
		return err
	}

	if err := u.TX.Commit(ctx); err != nil {
		return infraerrors.ErrorCommitOrRollback(err)
	}

	return nil
}

// --- Repository Accessors (Always fresh instances tied to current TX) --- //
func (u *UnitOfWork) CustomerRepository() interfaces.CustomerRepository {
	return repository.NewCustomerRepository(u.TX, u.prometheus)
}

func (u *UnitOfWork) AddressRepository() interfaces.AddressRepository {
	return repository.NewAddressRepository(u.TX, u.prometheus)
}
