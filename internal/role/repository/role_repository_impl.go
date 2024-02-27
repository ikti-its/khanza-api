package repository

import (
	"github.com/fathoor/simkes-api/internal/role/entity"
	"gorm.io/gorm"
)

type roleRepositoryImpl struct {
	*gorm.DB
}

func (repository *roleRepositoryImpl) Insert(role *entity.Role) error {
	return repository.DB.Create(&role).Error
}

func (repository *roleRepositoryImpl) FindAll() ([]entity.Role, error) {
	var roles []entity.Role
	err := repository.DB.Find(&roles).Error

	return roles, err
}

func (repository *roleRepositoryImpl) FindByRole(r string) (entity.Role, error) {
	var role entity.Role
	err := repository.DB.Take(&role, "nama = ?", r).Error

	return role, err
}

func (repository *roleRepositoryImpl) Update(role *entity.Role) error {
	return repository.DB.Save(&role).Error
}

func (repository *roleRepositoryImpl) Delete(role *entity.Role) error {
	return repository.DB.Delete(&role).Error
}

func NewRoleRepositoryProvider(db *gorm.DB) RoleRepository {
	return &roleRepositoryImpl{db}
}
