package uow

import (
	"context"

	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
	"github.com/andreis3/customers-ms/internal/domain/interfaces"
	"github.com/andreis3/customers-ms/internal/infra/repositories/postgres/repository"
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
func (u *UnitOfWork) Do(fn func(uow interfaces.UnitOfWork) *apperror.Error) *apperror.Error {
	ctx := context.Background()
	if u.TX != nil {
		return apperror.ErrorTransactionAlreadyExists()
	}

	tx, err := u.DB.Begin(ctx)
	if err != nil {
		return apperror.ErrorOpeningTransaction(err)
	}
	u.TX = tx
	defer func() { u.TX = nil }()

	if err := fn(u); err != nil {
		rollbackErr := u.TX.Rollback(ctx)
		if rollbackErr != nil {
			return apperror.ErrorExecuteRollback(rollbackErr)
		}
		return err
	}

	if err := u.TX.Commit(ctx); err != nil {
		return apperror.ErrorCommitOrRollback(err)
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
