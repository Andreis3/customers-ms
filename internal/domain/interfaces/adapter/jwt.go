package adapter

import (
	"github.com/andreis3/customers-ms/internal/domain/entity/customer"
	"github.com/andreis3/customers-ms/internal/domain/errors"
	valueobject "github.com/andreis3/customers-ms/internal/domain/value-object"
)

type TokenService interface {
	CreateToken(customer customer.Customer) (*valueobject.TokenClaims, *errors.Error)
	ValidateToken(tokenString string) (*valueobject.TokenClaims, *errors.Error)
	RefreshToken(tokenString string) (*valueobject.TokenClaims, *errors.Error)
}
