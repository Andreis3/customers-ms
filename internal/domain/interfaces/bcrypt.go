package interfaces

import apperror "github.com/andreis3/customers-ms/internal/domain/app-error"

type Bcrypt interface {
	Hash(data string) (string, *apperror.Error)
	CompareHash(hash, data string) bool
}
