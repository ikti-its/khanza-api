package service

import (
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/role/entity"
	"github.com/fathoor/simkes-api/internal/role/model"
	"github.com/fathoor/simkes-api/internal/role/repository"
	"github.com/fathoor/simkes-api/internal/role/validation"
)

type roleServiceImpl struct {
	repository.RoleRepository
}

func (service *roleServiceImpl) Create(request *model.RoleRequest) model.RoleResponse {
	if valid := validation.ValidateRoleRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	role := entity.Role{
		Nama: request.Nama,
	}

	if err := service.RoleRepository.Insert(&role); err != nil {
		exception.PanicIfError(err)
	}

	response := model.RoleResponse{
		Nama: role.Nama,
	}

	return response
}

func (service *roleServiceImpl) GetAll() []model.RoleResponse {
	roles, err := service.RoleRepository.FindAll()
	exception.PanicIfError(err)

	response := make([]model.RoleResponse, len(roles))
	for i, role := range roles {
		response[i] = model.RoleResponse{
			Nama: role.Nama,
		}
	}

	return response
}

func (service *roleServiceImpl) GetByRole(r string) model.RoleResponse {
	role, err := service.RoleRepository.FindByRole(r)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Role not found",
		})
	}

	response := model.RoleResponse{
		Nama: role.Nama,
	}

	return response
}

func (service *roleServiceImpl) Update(r string, request *model.RoleRequest) model.RoleResponse {
	if r == "Admin" {
		panic(exception.ForbiddenError{
			Message: "Role Admin is forbidden to be updated",
		})
	}

	if valid := validation.ValidateRoleRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	if _, err := service.RoleRepository.FindByRole(request.Nama); err == nil {
		panic(exception.BadRequestError{
			Message: "Role already exists",
		})
	}

	role, err := service.RoleRepository.FindByRole(r)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Role not found",
		})
	}

	role.Nama = request.Nama

	err = service.RoleRepository.Update(&role)
	exception.PanicIfError(err)

	response := model.RoleResponse{
		Nama: role.Nama,
	}

	return response
}

func (service *roleServiceImpl) Delete(r string) {
	if r == "Admin" {
		panic(exception.ForbiddenError{
			Message: "Role Admin is forbidden to be deleted",
		})
	}

	role, err := service.RoleRepository.FindByRole(r)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Role not found",
		})
	}

	if err := service.RoleRepository.Delete(&role); err != nil {
		exception.PanicIfError(err)
	}
}

func NewRoleServiceProvider(repository *repository.RoleRepository) RoleService {
	return &roleServiceImpl{*repository}
}
