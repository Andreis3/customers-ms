package repository

import (
	"context"
	"time"

	"github.com/andreis3/users-ms/internal/domain/entity"
	"github.com/andreis3/users-ms/internal/domain/errors"
	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/adapters/db/postegres"
	"github.com/andreis3/users-ms/internal/infra/adapters/observability"
	infra_errors "github.com/andreis3/users-ms/internal/infra/commons/errors"
	"github.com/andreis3/users-ms/internal/infra/repositories/postgres/model"
)

type AddressRepository struct {
	DB      *postegres.Queries
	metrics interfaces.Prometheus
	model.Address
}

func NewAddressRepository(metrics interfaces.Prometheus) *AddressRepository {
	return &AddressRepository{
		metrics: metrics,
	}
}

func (c *AddressRepository) InsertBatchAddress(ctx context.Context, customerID int64, addresses []entity.Address) (*[]entity.Address, *errors.AppErrors) {
	ctx, span := observability.Tracer.Start(ctx, "AddressRepository.InsertBatchAddress")
	start := time.Now()
	defer func() {
		end := time.Since(start)
		c.metrics.ObserveInstructionDBDuration("postgres", "addresses", "insert", float64(end.Milliseconds()))
		span.End()
	}()
	query := `
	INSERT INTO addresses 
	(customer_id, street, number, complement, city, state, postal_code, country, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
	RETURNING id`

	addressesResult := make([]entity.Address, 0, len(addresses))

	for _, address := range addresses {

		model := c.FromModel(address)
		var id int64

		err := c.DB.QueryRow(ctx, query,
			customerID,
			model.Street,
			model.Number,
			model.Complement,
			model.City,
			model.State,
			model.PostalCode,
			model.Country,
			model.CreatedAt,
			model.UpdatedAt,
		).Scan(&id)
		if err != nil {
			return nil, infra_errors.ErrorCreatedBatchAddress(err)
		}

		addressCopy := address
		addressCopy.ID = id
		addressesResult = append(addressesResult, addressCopy)
	}

	return &addressesResult, nil
}
