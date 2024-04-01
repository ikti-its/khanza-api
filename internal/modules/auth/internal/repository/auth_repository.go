package repository

import "github.com/fathoor/simkes-api/internal/modules/auth/internal/entity"

type AuthRepository interface {
	FindByEmail(email string) (entity.Auth, error)
}
