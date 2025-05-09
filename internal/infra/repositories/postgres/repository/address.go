package repository

import (
	"context"
	"time"

	"github.com/andreis3/users-ms/internal/domain/apperrors"
	"github.com/andreis3/users-ms/internal/domain/entity/address"
	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/adapters/observability"
	"github.com/andreis3/users-ms/internal/infra/repositories/postgres/model"
	"github.com/jackc/pgx/v5"
)

type AddressRepository struct {
	DB      interfaces.InstructionPostgres
	metrics interfaces.Prometheus
	model.Address
}

func NewAddressRepository(
	db interfaces.InstructionPostgres,
	metrics interfaces.Prometheus,
) *AddressRepository {
	return &AddressRepository{
		DB:      db,
		metrics: metrics,
	}
}

func (c *AddressRepository) InsertBatchAddress(ctx context.Context, customerID int64, addresses []address.Address) (*[]address.Address, *apperrors.AppErrors) {
	ctx, span := observability.Tracer.Start(ctx, "AddressRepository.InsertBatchAddress")
	start := time.Now()
	defer func() {
		end := time.Since(start)
		c.metrics.ObserveInstructionDBDuration("postgres", "addresses", "insert", float64(end.Milliseconds()))
		span.End()
	}()

	batch := &pgx.Batch{}

	query := `-- name: InsertAddress :one
	INSERT INTO addresses 
	(customer_id, street, number, complement, city, state, postal_code, country, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
	RETURNING id`

	for _, address := range addresses {
		model := c.FromModel(address)
		batch.Queue(query,
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
		)
	}

	br := c.DB.SendBatch(ctx, batch)
	defer br.Close()

	addressesResult := make([]address.Address, 0, len(addresses))

	for _, address := range addresses {

		var id int64

		_ = br.QueryRow().Scan(&id)

		addressCopy := address
		addressCopy.AssignID(id)

		addressesResult = append(addressesResult, addressCopy)
	}

	return &addressesResult, nil
}
