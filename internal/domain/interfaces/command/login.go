package command

import (
	"context"

	"github.com/andreis3/customers-ms/internal/domain/errors"
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
	Execute(ctx context.Context, input LoginInput) (*LoginOutput, *errors.Error)
}
