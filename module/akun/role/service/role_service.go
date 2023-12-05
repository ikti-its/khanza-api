package service

import "github.com/fathoor/simkes-api/module/akun/role/model"

type RoleService interface {
	Create(request *model.RoleRequest) error
	GetAll() ([]model.RoleResponse, error)
	GetByID(id int) (model.RoleResponse, error)
	Update(id int, request *model.RoleRequest) (model.RoleResponse, error)
	Delete(id int) error
}
