package services

import (
	"github.com/andreis3/customers-ms/internal/domain/entity"
	"github.com/andreis3/customers-ms/internal/domain/errors"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/valueobject"
)

type AuthService struct {
	tokenService adapter.TokenService
}

func NewAuthService(
	tokenService adapter.TokenService,
) *AuthService {
	return &AuthService{
		tokenService: tokenService,
	}
}

func (a *AuthService) GenerateToken(customer entity.Customer) (*valueobject.TokenClaims, *errors.Error) {
	return a.tokenService.CreateToken(customer)
}

func (a *AuthService) ValidateToken(tokenString string) (*valueobject.TokenClaims, *errors.Error) {
	return a.tokenService.ValidateToken(tokenString)
}

func (a *AuthService) RefreshToken(tokenString string) (*valueobject.TokenClaims, *errors.Error) {
	return a.tokenService.RefreshToken(tokenString)
}
