package postgres

import (
	"github.com/fathoor/simkes-api/internal/modules/auth/internal/entity"
	"github.com/fathoor/simkes-api/internal/modules/auth/internal/repository"
	"gorm.io/gorm"
)

type authRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) repository.AuthRepository {
	return &authRepositoryImpl{db}
}

func (r *authRepositoryImpl) FindByEmail(email string) (entity.Auth, error) {
	var auth entity.Auth

	err := r.DB.Table("akun").Select("id, email, password, role").Where("email = ?", email).First(&auth).Error

	return auth, err
}
