package crypto

import "golang.org/x/crypto/bcrypt"

type Crypto struct{}

func NewBcrypt() *Crypto {
	return &Crypto{}
}

func (b *Crypto) Hash(data string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	return string(bytes), err
}

func (b *Crypto) CompareHash(data string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data))
	return err == nil
}
