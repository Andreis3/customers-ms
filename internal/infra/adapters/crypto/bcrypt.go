package crypto

import (
	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
	"golang.org/x/crypto/bcrypt"
)

type BcryptCrypto struct{}

func NewBcrypt() *BcryptCrypto {
	return &BcryptCrypto{}
}

func (b *BcryptCrypto) Hash(data string) (string, *apperror.Error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(data), 5)
	if err != nil {
		return "", apperror.ErrorHashPassword(err)
	}
	return string(bytes), nil
}

func (b *BcryptCrypto) CompareHash(data string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(data)) == nil
}
