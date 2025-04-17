package interfaces

import (
	"github.com/andreis3/users-ms/internal/domain/apperrors"
)

type Bcrypt interface {
	Hash(data string) (string, *apperrors.AppErrors)
	CompareHash(hash, data string) bool
}
