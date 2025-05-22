package commands

import (
	"context"
	"log/slog"

	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/command"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/commons"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/repository"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/service"
	"github.com/andreis3/customers-ms/internal/infra/adapters/observability"
)

type AuthenticateCustomer struct {
	log                commons.Logger
	customerRepository repository.CustomerRepository
	authService        service.Auth
	bcrypt             adapter.Bcrypt
}

func NewAuthenticateCustomer(
	log commons.Logger,
	customerRepository repository.CustomerRepository,
	authService service.Auth,
	bcrypt adapter.Bcrypt,
) *AuthenticateCustomer {
	return &AuthenticateCustomer{
		log:                log,
		customerRepository: customerRepository,
		authService:        authService,
		bcrypt:             bcrypt,
	}
}

func (a *AuthenticateCustomer) Execute(ctx context.Context, input command.AuthenticateCustomerInput) (*command.AuthenticateCustomerOutput, *apperror.Error) {
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

	output := &command.AuthenticateCustomerOutput{
		Token:     token.Token,
		ExpiresAt: token.ExpiresAt.Unix(),
	}

	return output, nil
}
