package command

import (
	"context"

	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
)

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	Token     string
	ExpiresAt int64
}

type Login interface {
	Execute(ctx context.Context, input LoginInput) (*LoginOutput, *apperror.Error)
}
