package service

import (
	"github.com/andreis3/customers-ms/internal/domain/entity/customer"
	"github.com/andreis3/customers-ms/internal/domain/error"
	valueobject "github.com/andreis3/customers-ms/internal/domain/value-object"
)

type Auth interface {
	GenerateToken(customer customer.Customer) (*valueobject.TokenClaims, *error.Error)
	ValidateToken(tokenString string) (*valueobject.TokenClaims, *error.Error)
	RefreshToken(tokenString string) (*valueobject.TokenClaims, *error.Error)
}
