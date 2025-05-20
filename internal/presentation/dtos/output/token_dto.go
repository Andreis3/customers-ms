package output

import "github.com/andreis3/customers-ms/internal/app/commands"

type TokenDTO struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

func TokenOutputMapper(token *commands.AuthenticateCustomerOutput) TokenDTO {
	return TokenDTO{
		Token:     token.Token,
		ExpiresAt: token.ExpiresAt,
	}
}
