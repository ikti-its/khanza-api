package repository

import "github.com/fathoor/simkes-api/module/akun/role/entity"

type RoleRepository interface {
	Insert(role *entity.Role) error
	FindAll() ([]entity.Role, error)
	FindByID(id int) (entity.Role, error)
	Update(role *entity.Role) error
	Delete(role *entity.Role) error
}
