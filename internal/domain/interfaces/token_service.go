package interfaces

import (
	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
	"github.com/andreis3/customers-ms/internal/domain/entity/customer"
	valueobject "github.com/andreis3/customers-ms/internal/domain/value-object"
)

type TokenService interface {
	CreateToken(customer customer.Customer) (*valueobject.TokenClaims, *apperror.Error)
	ValidateToken(tokenString string) (*valueobject.TokenClaims, *apperror.Error)
	RefreshToken(tokenString string) (*valueobject.TokenClaims, *apperror.Error)
}
