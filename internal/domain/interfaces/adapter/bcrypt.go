package adapter

import "github.com/andreis3/customers-ms/internal/domain/error"

type Bcrypt interface {
	Hash(data string) (string, *error.Error)
	CompareHash(hash, data string) bool
}
