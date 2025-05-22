package app

import (
	"github.com/andreis3/customers-ms/internal/app/commands"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/command"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/commons"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/repository"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/service"
)

type IAuthenticateCustomerFactory interface {
	Build() command.AuthenticateCustomer
}

type authenticateCustomerFactory struct {
	log                commons.Logger
	customerRepository repository.CustomerRepository
	authService        service.Auth
	bcrypt             adapter.Bcrypt
}

func NewAuthenticateCustomerFactory(
	log commons.Logger,
	customerRepository repository.CustomerRepository,
	authService service.Auth,
	bcrypt adapter.Bcrypt,
) IAuthenticateCustomerFactory {
	return &authenticateCustomerFactory{
		log:                log,
		customerRepository: customerRepository,
		authService:        authService,
		bcrypt:             bcrypt,
	}
}

func (f *authenticateCustomerFactory) Build() command.AuthenticateCustomer {
	return commands.NewAuthenticateCustomer(f.log, f.customerRepository, f.authService, f.bcrypt)
}
