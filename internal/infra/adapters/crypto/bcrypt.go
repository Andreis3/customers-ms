package crypto

import (
	"github.com/andreis3/customers-ms/internal/domain/apperrors"
	"github.com/andreis3/customers-ms/internal/infra/commons/infraerrors"
	"golang.org/x/crypto/bcrypt"
)

type BcryptCrypto struct{}

func NewBcrypt() *BcryptCrypto {
	return &BcryptCrypto{}
}

func (b *BcryptCrypto) Hash(data string) (string, *apperrors.AppErrors) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(data), 5)
	if err != nil {
		return "", infraerrors.ErrorHashPassword(err)
	}
	return string(bytes), nil
}

func (b *BcryptCrypto) CompareHash(data string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(data)) == nil
}
