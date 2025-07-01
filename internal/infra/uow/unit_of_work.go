package uow

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/andreis3/customers-ms/internal/domain/errors"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/postgres"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/uow"
	"github.com/andreis3/customers-ms/internal/infra/adapters/observability"
	"github.com/andreis3/customers-ms/internal/infra/repositories/postgres/repository"
)

type UnitOfWork struct {
	DB         *pgxpool.Pool
	TX         pgx.Tx
	prometheus adapter.Prometheus
}

func NewUnitOfWork(db *pgxpool.Pool, prometheus adapter.Prometheus) *UnitOfWork {
	return &UnitOfWork{
		DB:         db,
		prometheus: prometheus,
	}
}

// Do handles transaction lifecycle safely.
func (u *UnitOfWork) Do(ctx context.Context, fn func(uow uow.UnitOfWork) *errors.Error) *errors.Error {
	ctx, child := observability.Tracer.Start(ctx, "UnitOfWork.Do")
	defer child.End()

	if u.TX != nil {
		child.RecordError(errors.ErrorTransactionAlreadyExists())
		return errors.ErrorTransactionAlreadyExists()
	}

	tx, err := u.DB.BeginTx(ctx, pgx.TxOptions{
		BeginQuery:  "BEGIN",
		CommitQuery: "COMMIT",
		AccessMode:  pgx.ReadWrite,
	})
	if err != nil {
		child.RecordError(errors.ErrorOpeningTransaction(err))
		return errors.ErrorOpeningTransaction(err)
	}

	u.TX = tx
	defer func() { u.TX = nil }()

	if err := fn(u); err != nil {
		rollbackErr := u.TX.Rollback(ctx)
		if rollbackErr != nil {
			child.RecordError(errors.ErrorExecuteRollback(rollbackErr))
			return errors.ErrorExecuteRollback(rollbackErr)
		}
		child.RecordError(err)
		return err
	}

	if err := u.TX.Commit(ctx); err != nil {
		child.RecordError(errors.ErrorCommitOrRollback(err))
		return errors.ErrorCommitOrRollback(err)
	}

	return nil
}

// --- Repository Accessors (Always fresh instances tied to current TX) --- //
func (u *UnitOfWork) CustomerRepository() postgres.CustomerRepository {
	return repository.NewCustomerRepository(u.TX, u.prometheus)
}

func (u *UnitOfWork) AddressRepository() postgres.AddressRepository {
	return repository.NewAddressRepository(u.TX, u.prometheus)
}
