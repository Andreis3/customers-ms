package valueobject

import (
	"regexp"
	"strings"

	"github.com/andreis3/users-ms/internal/domain/validator"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type Email struct {
	Address   string
	Validator validator.Validator
}

// NewEmail cria uma nova instância do Email
func NewEmail(email string) *Email {
	return &Email{Address: email}
}

// Validate executa as validações no email
func (e *Email) Validate() {
	cleanedEmail := cleanEmail(e.Address)
	e.Validator.Assert(validator.NotBlank(cleanedEmail), "email", validator.NotBlankField)
	e.Validator.Assert(isValidEmail(cleanedEmail), "email", "invalid email format")
}

// cleanEmail remove espaços em branco desnecessários do e-mail
func cleanEmail(email string) string {
	return strings.TrimSpace(email)
}

// isValidEmail verifica se o e-mail está no formato correto
func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}
