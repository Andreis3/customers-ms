package app

import (
	"github.com/andreis3/customers-ms/internal/app/decorator"
	"github.com/andreis3/customers-ms/internal/app/queries"
	"github.com/andreis3/customers-ms/internal/app/services"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/query"
	"github.com/andreis3/customers-ms/internal/infra/adapters/db"
	"github.com/andreis3/customers-ms/internal/infra/repositories/repository"
)

func NewGetCustomerAddressesFactory(
	db *db.Postgres,
	cache *db.Redis,
	log adapter.Logger,
	tracer adapter.Tracer,
	metrics adapter.Prometheus,
) query.GetCustomerAddresses {
	customerRepository := repository.NewCustomerRepository(db, metrics, tracer)
	addressRepository := repository.NewAddressRepository(db, metrics, tracer)
	redisCache := repository.NewCache(cache.Client(), metrics, tracer)
	addressesDecorator := decorator.NewCachedAddressesRepository(addressRepository, redisCache)
	customerService := services.NewCustomerService(customerRepository)
	return queries.NewGetCustomerAddresses(log, addressesDecorator, customerService, tracer)
}
