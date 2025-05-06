//go:build wireinject
// +build wireinject

package providers

import (
	"github.com/andreis3/customers-ms/internal/domain/interfaces"
	"github.com/andreis3/customers-ms/internal/presentation/http/handlers/customer"
	"github.com/google/wire"
)

func NewCreateCustomerHandler(
	log interfaces.Logger,
	prometheus interfaces.Prometheus,
	crypto interfaces.Bcrypt,
	uow interfaces.UnitOfWork,
) customer.CreateCustomerHandler {
	return customer.NewCreateCustomerHandler(log, prometheus, crypto, uow)
}

var CreateCustomerHandlerSet = wire.NewSet(NewCreateCustomerHandler)
