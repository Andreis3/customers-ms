package app

import (
	"github.com/andreis3/customers-ms/internal/app/commands"
	"github.com/andreis3/customers-ms/internal/domain/interfaces"
)

type IAuthenticateCustomerFactory interface {
	Build() commands.IAuthenticateCustomer
}

type authenticateCustomerFactory struct {
	log                interfaces.Logger
	customerRepository interfaces.CustomerRepository
	authService        interfaces.Auth
	bcrypt             interfaces.Bcrypt
}

func NewAuthenticateCustomerFactory(
	log interfaces.Logger,
	customerRepository interfaces.CustomerRepository,
	authService interfaces.Auth,
	bcrypt interfaces.Bcrypt,
) IAuthenticateCustomerFactory {
	return &authenticateCustomerFactory{
		log:                log,
		customerRepository: customerRepository,
		authService:        authService,
		bcrypt:             bcrypt,
	}
}

func (f *authenticateCustomerFactory) Build() commands.IAuthenticateCustomer {
	return commands.NewAuthenticateCustomer(f.log, f.customerRepository, f.authService, f.bcrypt)
}
