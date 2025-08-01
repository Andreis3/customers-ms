package service

import (
	"context"

	"github.com/andreis3/customers-ms/internal/domain/entity"
	"github.com/andreis3/customers-ms/internal/domain/errors"
	"github.com/andreis3/customers-ms/internal/domain/valueobject"
)

type Auth interface {
	GenerateToken(customer entity.Customer) (*valueobject.TokenClaims, *errors.Error)
	ValidateToken(tokenString string) (*valueobject.TokenClaims, *errors.Error)
	RefreshToken(tokenString string) (*valueobject.TokenClaims, *errors.Error)

	DecodeToken(ctx context.Context, tokenString string) (*valueobject.TokenClaims, *errors.Error)
}
