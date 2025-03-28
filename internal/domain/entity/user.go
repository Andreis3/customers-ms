package entity

import (
	"github.com/andreis3/users-ms/internal/domain/validator"
)

type User struct {
	ID        string
	Email     string
	Password  string
	Validator validator.Validator
}

func (u *User) Validate() map[string]string {
	u.Validator.Assert(validator.NotBlank(u.Email), "email", validator.NotBlankField)
	u.Validator.Assert(validator.NotBlank(u.Password), "password", validator.NotBlankField)
	return u.Validator.FieldErrors
}
