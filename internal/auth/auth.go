package auth

import (
	"github.com/invzhi/outward/internal/sqlc"
)

func CurrentUserID() (sqlc.ID, error) {
	// TODO
	return sqlc.ID(1), nil
}
