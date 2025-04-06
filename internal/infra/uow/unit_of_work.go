package uow

import (
	"context"

	"github.com/andreis3/users-ms/internal/app/interfaces"
	"github.com/andreis3/users-ms/internal/domain/errors"
	infra_errors "github.com/andreis3/users-ms/internal/infra/commons/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	InternalServerError = "internal server error"
)

type UnitOfWork struct {
	DB           *pgxpool.Pool
	TX           pgx.Tx
	Repositories map[string]interfaces.RepositoryFactory
}

func NewUnitOfWork(db *pgxpool.Pool) *UnitOfWork {
	return &UnitOfWork{
		DB:           db,
		Repositories: make(map[string]interfaces.RepositoryFactory),
	}
}

func (uow *UnitOfWork) Register(name string, callback interfaces.RepositoryFactory) {
	uow.Repositories[name] = callback
}

func (u *UnitOfWork) GetRepository(name string) any {
	ctx := context.Background()
	if u.TX == nil {
		tx, err := u.DB.Begin(ctx)
		if err != nil {
			return nil
		}
		u.TX = tx
	}

	return u.Repositories[name](u.TX)
}

func (u *UnitOfWork) Do(callback func(uow interfaces.UnitOfWork) *errors.AppErrors) *errors.AppErrors {
	ctx := context.Background()
	if u.TX != nil {
		return infra_errors.ErrorTransactionAlreadyExists()
	}

	tx, err := u.DB.Begin(ctx)
	if err != nil {
		return infra_errors.ErrorOpeningTransaction(err)
	}
	u.TX = tx

	if err := callback(u); err != nil {
		if errRB := u.Rollback(); errRB != nil {
			return errRB
		}
		return err
	}

	return u.CommitOrRollback()
}

func (u *UnitOfWork) Rollback() *errors.AppErrors {
	if u.TX == nil {
		return infra_errors.ErrorRollBackTransactionEmpty()
	}

	defer func() {
		u.TX = nil
	}()

	err := u.TX.Rollback(context.Background())
	if err != nil {
		return infra_errors.ErrorExecuteRollback(err)
	}
	return nil
}

func (u *UnitOfWork) CommitOrRollback() *errors.AppErrors {
	ctx := context.Background()
	defer func() {
		u.TX = nil
	}()

	if u.TX == nil {
		return nil
	}

	if err := u.TX.Commit(ctx); err != nil {
		if errRB := u.Rollback(); errRB != nil {
			return errRB
		}
		return infra_errors.ErrorCommitOrRollback(err)
	}

	return nil
}
