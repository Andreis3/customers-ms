package uow

import (
	"context"

	"github.com/andreis3/users-ms/internal/domain/apperrors"
	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/commons/infraerrors"
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

func (u *UnitOfWork) Do(callback func(uow interfaces.UnitOfWork) *apperrors.AppErrors) *apperrors.AppErrors {
	ctx := context.Background()
	if u.TX != nil {
		return infraerrors.ErrorTransactionAlreadyExists()
	}

	tx, err := u.DB.Begin(ctx)
	if err != nil {
		return infraerrors.ErrorOpeningTransaction(err)
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

func (u *UnitOfWork) Rollback() *apperrors.AppErrors {
	if u.TX == nil {
		return infraerrors.ErrorRollBackTransactionEmpty()
	}

	defer func() {
		u.TX = nil
	}()

	err := u.TX.Rollback(context.Background())
	if err != nil {
		return infraerrors.ErrorExecuteRollback(err)
	}
	return nil
}

func (u *UnitOfWork) CommitOrRollback() *apperrors.AppErrors {
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
		return infraerrors.ErrorCommitOrRollback(err)
	}

	return nil
}
