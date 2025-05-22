package command

import (
	"context"

	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
)

type AuthenticateCustomerInput struct {
	Email    string
	Password string
}

type AuthenticateCustomerOutput struct {
	Token     string
	ExpiresAt int64
}

type AuthenticateCustomer interface {
	Execute(ctx context.Context, input AuthenticateCustomerInput) (*AuthenticateCustomerOutput, *apperror.Error)
}
