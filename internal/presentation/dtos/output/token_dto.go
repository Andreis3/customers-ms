package output

import (
	"github.com/andreis3/customers-ms/internal/domain/interfaces/command"
)

type TokenDTO struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

func TokenOutputMapper(token *command.LoginOutput) TokenDTO {
	return TokenDTO{
		Token:     token.Token,
		ExpiresAt: token.ExpiresAt,
	}
}
