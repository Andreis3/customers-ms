package aggregate

import "github.com/andreis3/users-ms/internal/domain/entity"

type UserProfile struct {
	User    entity.User
	Address entity.Address
}
