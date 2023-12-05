package service

import (
	"github.com/fathoor/simkes-api/core/exception"
	"github.com/fathoor/simkes-api/module/akun/role/entity"
	"github.com/fathoor/simkes-api/module/akun/role/model"
	"github.com/fathoor/simkes-api/module/akun/role/repository"
	"github.com/fathoor/simkes-api/module/akun/role/validation"
)

type roleServiceImpl struct {
	repository.RoleRepository
}

func (service *roleServiceImpl) Create(request *model.RoleRequest) error {
	valid := validation.ValidateRoleRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	role := entity.Role{
		Role: request.Role,
	}

	return service.RoleRepository.Insert(&role)
}

func (service *roleServiceImpl) GetAll() ([]model.RoleResponse, error) {
	roles, err := service.RoleRepository.FindAll()

	response := make([]model.RoleResponse, len(roles))
	for i, role := range roles {
		response[i] = model.RoleResponse{
			ID:   role.ID,
			Role: role.Role,
		}
	}

	return response, err
}

func (service *roleServiceImpl) GetByID(id int) (model.RoleResponse, error) {
	role, err := service.RoleRepository.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Role not found",
		})
	}

	response := model.RoleResponse{
		ID:   role.ID,
		Role: role.Role,
	}

	return response, err
}

func (service *roleServiceImpl) Update(id int, request *model.RoleRequest) (model.RoleResponse, error) {
	valid := validation.ValidateRoleRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	role, err := service.RoleRepository.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Role not found",
		})
	}

	if role != (entity.Role{}) {
		role.Role = request.Role
	}

	err = service.RoleRepository.Update(&role)

	response := model.RoleResponse{
		ID:   role.ID,
		Role: role.Role,
	}

	return response, err
}

func (service *roleServiceImpl) Delete(id int) error {
	role, err := service.RoleRepository.FindByID(id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Role not found",
		})
	}

	return service.RoleRepository.Delete(&role)
}

func ProvideRoleService(repository *repository.RoleRepository) RoleService {
	return &roleServiceImpl{*repository}
}
