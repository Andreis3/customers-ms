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

func NewEmail(email string) *Email {
	return &Email{Address: email}
}

func (e *Email) Validate() {
	cleanedEmail := cleanEmail(e.Address)
	e.Validator.Assert(validator.NotBlank(cleanedEmail), "email", validator.NotBlankField)
	e.Validator.Assert(isValidEmail(cleanedEmail), "email", "invalid email format")
}

func cleanEmail(email string) string {
	return strings.TrimSpace(email)
}

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}
