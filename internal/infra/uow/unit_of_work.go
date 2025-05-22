package uow

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	irepository "github.com/andreis3/customers-ms/internal/domain/interfaces/repository"
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
func (u *UnitOfWork) Do(ctx context.Context, fn func(uow uow.UnitOfWork) *apperror.Error) *apperror.Error {
	ctx, child := observability.Tracer.Start(ctx, "UnitOfWork.Do")
	defer child.End()

	if u.TX != nil {
		child.RecordError(apperror.ErrorTransactionAlreadyExists())
		return apperror.ErrorTransactionAlreadyExists()
	}

	tx, err := u.DB.BeginTx(ctx, pgx.TxOptions{
		BeginQuery:  "BEGIN",
		CommitQuery: "COMMIT",
		AccessMode:  pgx.ReadWrite,
	})
	if err != nil {
		child.RecordError(apperror.ErrorOpeningTransaction(err))
		return apperror.ErrorOpeningTransaction(err)
	}

	u.TX = tx
	defer func() { u.TX = nil }()

	if err := fn(u); err != nil {
		rollbackErr := u.TX.Rollback(ctx)
		if rollbackErr != nil {
			child.RecordError(apperror.ErrorExecuteRollback(rollbackErr))
			return apperror.ErrorExecuteRollback(rollbackErr)
		}
		child.RecordError(err)
		return err
	}

	if err := u.TX.Commit(ctx); err != nil {
		child.RecordError(apperror.ErrorCommitOrRollback(err))
		return apperror.ErrorCommitOrRollback(err)
	}

	return nil
}

// --- Repository Accessors (Always fresh instances tied to current TX) --- //
func (u *UnitOfWork) CustomerRepository() irepository.CustomerRepository {
	return repository.NewCustomerRepository(u.TX, u.prometheus)
}

func (u *UnitOfWork) AddressRepository() irepository.AddressRepository {
	return repository.NewAddressRepository(u.TX, u.prometheus)
}
