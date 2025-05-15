package commands

import (
	"context"

	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
	"github.com/andreis3/customers-ms/internal/domain/interfaces"
)

type AuthenticateCustomer struct {
	customerRepository interfaces.CustomerRepository
	authService        interfaces.Auth
	bcrypt             interfaces.Bcrypt
}

type AuthenticateCustomerInput struct {
	Email    string
	Password string
}

func NewAuthenticateCustomer(
	customerRepository interfaces.CustomerRepository,
	authService interfaces.Auth,
	bcrypt interfaces.Bcrypt,
) *AuthenticateCustomer {
	return &AuthenticateCustomer{
		customerRepository: customerRepository,
		authService:        authService,
		bcrypt:             bcrypt,
	}
}

func (a *AuthenticateCustomer) Execute(ctx context.Context, input AuthenticateCustomerInput) (string, *apperror.Error) {
	customer, err := a.customerRepository.FindCustomerByEmail(ctx, input.Email)
	if err != nil {
		return "", err
	}

	isValid := a.bcrypt.CompareHash(customer.Password(), input.Password)
	if !isValid {
		return "", err
	}

	token, err := a.authService.GenerateToken(*customer)
	if err != nil {
		return "", err
	}

	return token, nil
}
