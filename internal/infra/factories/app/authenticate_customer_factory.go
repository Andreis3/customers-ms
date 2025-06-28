package app

import (
	"github.com/andreis3/customers-ms/internal/app/commands"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/command"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/commons"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/postgres"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/service"
)

type AuthenticateCustomerFactory interface {
	Build() command.Login
}

type authenticateCustomerFactory struct {
	log                commons.Logger
	customerRepository postgres.CustomerRepository
	authService        service.Auth
	bcrypt             adapter.Bcrypt
}

func NewAuthenticateCustomerFactory(
	log commons.Logger,
	customerRepository postgres.CustomerRepository,
	authService service.Auth,
	bcrypt adapter.Bcrypt,
) AuthenticateCustomerFactory {
	return &authenticateCustomerFactory{
		log:                log,
		customerRepository: customerRepository,
		authService:        authService,
		bcrypt:             bcrypt,
	}
}

func (f *authenticateCustomerFactory) Build() command.Login {
	return commands.NewAuthenticateCustomer(f.log, f.customerRepository, f.authService, f.bcrypt)
}
