package repository

import "github.com/ikti-its/khanza-api/internal/modules/auth/internal/entity"

type AuthRepository interface {
	FindByEmail(email string) (entity.Auth, error)
}
