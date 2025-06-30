package crypto

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/andreis3/customers-ms/internal/domain/error"
)

type Bcrypt struct{}

func NewBcrypt() *Bcrypt {
	return &Bcrypt{}
}

func (b *Bcrypt) Hash(data string) (string, *error.Error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(data), 5)
	if err != nil {
		return "", error.ErrorHashPassword(err)
	}
	return string(bytes), nil
}

func (b *Bcrypt) CompareHash(data string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(data)) == nil
}
