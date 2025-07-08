package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/andreis3/customers-ms/internal/domain/entity"
	"github.com/andreis3/customers-ms/internal/domain/errors"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/infra/adapters/db/postegres"
	"github.com/andreis3/customers-ms/internal/infra/repositories/postgres/model"
)

type AddressRepository struct {
	DB      adapter.InstructionPostgres
	metrics adapter.Prometheus
	model.Address
	tracer adapter.Tracer
}

func NewAddressRepository(
	db adapter.InstructionPostgres,
	metrics adapter.Prometheus,
	tracer adapter.Tracer,
) *AddressRepository {
	return &AddressRepository{
		DB:      db,
		metrics: metrics,
		tracer:  tracer,
	}
}

func (c *AddressRepository) InsertBatchAddress(ctx context.Context, customerID int64, addresses []entity.Address) (*[]entity.Address, *errors.Error) {
	ctx, span := c.tracer.Start(ctx, "AddressRepository.InsertBatchAddress")
	start := time.Now()

	defer func() {
		end := time.Since(start)
		c.metrics.ObserveInstructionDBDuration("postgres", "addresses", "insert", float64(end.Milliseconds()))
		span.End()
	}()

	batch := &pgx.Batch{}

	query := `
	INSERT INTO addresses 
	(customer_id, street, number, complement, city, state, postal_code, country, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
	RETURNING id`

	for _, address := range addresses {
		modelAddress := c.FromModel(address)
		batch.Queue(query,
			customerID,
			modelAddress.Street,
			modelAddress.Number,
			modelAddress.Complement,
			modelAddress.City,
			modelAddress.State,
			modelAddress.PostalCode,
			modelAddress.Country,
			modelAddress.CreatedAt,
			modelAddress.UpdatedAt,
		)
	}

	db := c.resolveDB(ctx)

	br := db.SendBatch(ctx, batch)
	defer br.Close()

	addressesResult := make([]entity.Address, 0, len(addresses))

	for _, address := range addresses {

		var id int64

		_ = br.QueryRow().Scan(&id)

		addressCopy := address
		addressCopy.AssignID(id)

		addressesResult = append(addressesResult, addressCopy)
	}

	return &addressesResult, nil
}

func (c *AddressRepository) resolveDB(ctx context.Context) adapter.InstructionPostgres {
	if tx, ok := postegres.TxFromContext(ctx); ok {
		return tx
	}
	return c.DB
}
