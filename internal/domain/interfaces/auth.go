package interfaces

import (
	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
	"github.com/andreis3/customers-ms/internal/domain/entity/customer"
	valueobject "github.com/andreis3/customers-ms/internal/domain/value-object"
)

type Auth interface {
	GenerateToken(customer customer.Customer) (string, *apperror.Error)
	ValidateToken(tokenString string) (*valueobject.TokenClaims, *apperror.Error)
	RefreshToken(tokenString string) (string, *apperror.Error)
}
