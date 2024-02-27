package repository

import (
	"github.com/fathoor/simkes-api/internal/role/entity"
)

type RoleRepository interface {
	Insert(role *entity.Role) error
	FindAll() ([]entity.Role, error)
	FindByRole(r string) (entity.Role, error)
	Update(role *entity.Role) error
	Delete(role *entity.Role) error
}
