package commands

import (
	"context"
	"log/slog"

	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
	"github.com/andreis3/customers-ms/internal/domain/interfaces"
	"github.com/andreis3/customers-ms/internal/infra/adapters/observability"
)

type AuthenticateCustomer struct {
	log                interfaces.Logger
	customerRepository interfaces.CustomerRepository
	authService        interfaces.Auth
	bcrypt             interfaces.Bcrypt
}

type AuthenticateCustomerInput struct {
	Email    string
	Password string
}

type AuthenticateCustomerOutput struct {
	Token     string
	ExpiresAt int64
}

func NewAuthenticateCustomer(
	log interfaces.Logger,
	customerRepository interfaces.CustomerRepository,
	authService interfaces.Auth,
	bcrypt interfaces.Bcrypt,
) *AuthenticateCustomer {
	return &AuthenticateCustomer{
		log:                log,
		customerRepository: customerRepository,
		authService:        authService,
		bcrypt:             bcrypt,
	}
}

func (a *AuthenticateCustomer) Execute(ctx context.Context, input AuthenticateCustomerInput) (*AuthenticateCustomerOutput, *apperror.Error) {
	a.log.InfoText("Received input to authenticate customer",
		slog.String("email", input.Email),
		slog.String("password", input.Password))

	ctx, child := observability.Tracer.Start(ctx, "AuthenticateCustomer.Execute")

	defer child.End()
	traceID := child.SpanContext().TraceID().String()

	customer, err := a.customerRepository.FindCustomerByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, apperror.ErrorInvalidCredentials()
	}

	isValid := a.bcrypt.CompareHash(input.Password, customer.Password())
	if !isValid {
		return nil, apperror.ErrorInvalidCredentials()
	}

	token, err := a.authService.GenerateToken(*customer)

	if err != nil {
		return nil, err
	}

	a.log.InfoJSON("end request",
		slog.String("trace_id", traceID),
		slog.String("token", token.Token))

	output := &AuthenticateCustomerOutput{
		Token:     token.Token,
		ExpiresAt: token.ExpiresAt.Unix(),
	}
	
	return output, nil
}
