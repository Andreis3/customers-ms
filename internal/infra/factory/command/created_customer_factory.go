package command

import (
	"github.com/andreis3/users-ms/internal/app/commands"
	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/adapters/crypto"
	"github.com/andreis3/users-ms/internal/infra/commons/logger"
)

func NewCreatedCustomerFactory(
	uow interfaces.UnitOfWork,
) commands.CreatedCustomerInterface {
	bcrypt := crypto.NewBcrypt()
	log := logger.NewLogger()
	return commands.NewCreatedCustomer(uow, bcrypt, log)
}
