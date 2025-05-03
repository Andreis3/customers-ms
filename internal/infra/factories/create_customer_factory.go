package factories

import (
	"github.com/andreis3/users-ms/internal/app/commands"
	"github.com/andreis3/users-ms/internal/domain/interfaces"
)

type createCustomerFactory struct {
	uow   interfaces.UnitOfWork
	crypt interfaces.Bcrypt
	log   interfaces.Logger
}

type ICreateCustomerFactory interface {
	Build() commands.ICreateCustomer
}

func NewCreateCustomerFactory(
	uow interfaces.UnitOfWork,
	crypto interfaces.Bcrypt,
	log interfaces.Logger,
) ICreateCustomerFactory {
	return &createCustomerFactory{uow: uow, crypt: crypto, log: log}
}

func (f *createCustomerFactory) Build() commands.ICreateCustomer {
	return commands.NewCreateCustomer(f.uow, f.crypt, f.log)
}
