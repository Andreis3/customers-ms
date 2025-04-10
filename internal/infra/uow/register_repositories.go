package uow

import (
	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/adapters/db/postegres"
	"github.com/andreis3/users-ms/internal/infra/repositories/postgres/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	CustomerRepository = "customer_repository"
)

func NewRegisterRepositories(pool *pgxpool.Pool, metrics interfaces.Prometheus) {
	uow := NewUnitOfWork(pool)
	uow.Register(CustomerRepository, func(tx any) any {
		repo := repository.NewCustomerRepository(metrics)
		repo.DB = postegres.New(tx.(pgx.Tx))
		return repo
	})
}
