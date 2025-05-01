package factories

import (
	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/adapters/observability"
	"github.com/andreis3/users-ms/internal/infra/repositories/postgres/repository"
)

type CreateCustomer struct {
	Customer interfaces.CustomerRepository
	Address  interfaces.AddressRepository
}

func LoadCustomerFactory(db interfaces.InstructionPostgres) *CreateCustomer {
	prometheus := observability.NewPrometheus()

	customerRepo := repository.NewCustomerRepository(db, prometheus)

	addressRepo := repository.NewAddressRepository(db, prometheus)
	return &CreateCustomer{
		Customer: customerRepo,
		Address:  addressRepo,
	}
}
