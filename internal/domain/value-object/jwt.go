package valueobject

import "time"

type TokenClaims struct {
	CustomerID int64
	FullName   string
	Email      string
	ExpiresAt  time.Time
}
