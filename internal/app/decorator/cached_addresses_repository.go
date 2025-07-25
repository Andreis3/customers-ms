package decorator

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/andreis3/customers-ms/internal/domain/entity"
	"github.com/andreis3/customers-ms/internal/domain/errors"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/postgres"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/redis"
	"github.com/andreis3/customers-ms/internal/infra/repositories/criteria"
)

type CachedAddressesRepository struct {
	inner postgres.AddressRepository
	cache redis.Cache
}

func NewCachedAddressesRepository(inner postgres.AddressRepository, cache redis.Cache) postgres.AddressesSearchRepository {
	return &CachedAddressesRepository{
		inner: inner,
		cache: cache,
	}
}

func (c *CachedAddressesRepository) SearchAddresses(ctx context.Context, params criteria.AddressSearchCriteria) (*[]entity.Address, *errors.Error) {
	cachedKey, err := c.generateCacheKey(params)
	if err != nil {
		return nil, errors.ErrorGenerateCacheKey(err)
	}

	var cached []entity.Address

	found, err := c.cache.Get(ctx, cachedKey, &cached)
	if err != nil {
		return nil, errors.ErrorGetCache(err)
	}

	if found {
		return &cached, nil
	}

	addresses, err := c.inner.SearchAddresses(ctx, params)
	if err != nil {
		return nil, errors.ErrorSearchAddresses(err)
	}

	const ttl = 5
	err = c.cache.Set(ctx, cachedKey, addresses, ttl)
	if err != nil {
		return nil, errors.ErrorSetCache(err)
	}
	return addresses, nil
}

func (c *CachedAddressesRepository) generateCacheKey(params criteria.AddressSearchCriteria) (string, *errors.Error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", errors.ErrorGenerateCacheKey(err)
	}
	sum := sha256.Sum256(bytes)
	return "customers:search:" + hex.EncodeToString(sum[:]), nil
}
