package service

import (
	"github.com/fathoor/simkes-api/internal/akun/entity"
	"github.com/fathoor/simkes-api/internal/akun/model"
	"github.com/fathoor/simkes-api/internal/akun/repository"
	"github.com/fathoor/simkes-api/internal/akun/validation"
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/app/helper"
)

type akunServiceImpl struct {
	repository.AkunRepository
}

func (service *akunServiceImpl) Create(request *model.AkunRequest) model.AkunResponse {
	if valid := validation.ValidateAkunRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	encrypted, err := helper.EncryptPassword(request.Password)
	exception.PanicIfError(err)

	akun := entity.Akun{
		NIP:      request.NIP,
		Email:    request.Email,
		Password: string(encrypted),
		RoleNama: request.RoleNama,
	}

	if err := service.AkunRepository.Insert(&akun); err != nil {
		exception.PanicIfError(err)
	}

	response := model.AkunResponse{
		NIP:      akun.NIP,
		Email:    akun.Email,
		RoleNama: akun.RoleNama,
	}

	return response
}

func (service *akunServiceImpl) GetAll() []model.AkunResponse {
	akun, err := service.AkunRepository.FindAll()
	exception.PanicIfError(err)

	response := make([]model.AkunResponse, len(akun))
	for i, akun := range akun {
		response[i] = model.AkunResponse{
			NIP:      akun.NIP,
			Email:    akun.Email,
			RoleNama: akun.RoleNama,
		}
	}

	return response
}

func (service *akunServiceImpl) GetPage(page, size int) model.AkunPageResponse {
	akun, total, err := service.AkunRepository.FindPage(page, size)
	exception.PanicIfError(err)

	response := make([]model.AkunResponse, len(akun))
	for i, akun := range akun {
		response[i] = model.AkunResponse{
			NIP:      akun.NIP,
			Email:    akun.Email,
			RoleNama: akun.RoleNama,
		}
	}

	pagedResponse := model.AkunPageResponse{
		Akun:  response,
		Page:  page,
		Size:  size,
		Total: total,
	}

	return pagedResponse
}

func (service *akunServiceImpl) GetByNIP(nip string) model.AkunResponse {
	akun, err := service.AkunRepository.FindByNIP(nip)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	response := model.AkunResponse{
		NIP:      akun.NIP,
		Email:    akun.Email,
		RoleNama: akun.RoleNama,
	}

	return response
}

func (service *akunServiceImpl) Update(nip string, request *model.AkunRequest) model.AkunResponse {
	if valid := validation.ValidateAkunRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	akun, err := service.AkunRepository.FindByNIP(nip)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	if request.Password != "" {
		encrypted, err := helper.EncryptPassword(request.Password)
		exception.PanicIfError(err)

		akun.Password = string(encrypted)
	}

	akun.Email = request.Email

	response := model.AkunResponse{
		NIP:      akun.NIP,
		Email:    akun.Email,
		RoleNama: akun.RoleNama,
	}

	if err := service.AkunRepository.Update(&akun); err != nil {
		exception.PanicIfError(err)
	}

	return response
}

func (service *akunServiceImpl) UpdateAdmin(nip string, request *model.AkunRequest) model.AkunResponse {
	if valid := validation.ValidateAkunRequest(request); valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	akun, err := service.AkunRepository.FindByNIP(nip)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	if request.Password != "" {
		encrypted, err := helper.EncryptPassword(request.Password)
		exception.PanicIfError(err)

		akun.Password = string(encrypted)
	}

	if _, err := service.AkunRepository.FindByNIP(request.NIP); err == nil {
		panic(exception.BadRequestError{
			Message: "NIP already exists",
		})
	} else {
		akun.NIP = request.NIP
	}

	akun.Email = request.Email
	akun.RoleNama = request.RoleNama

	response := model.AkunResponse{
		NIP:      akun.NIP,
		Email:    akun.Email,
		RoleNama: akun.RoleNama,
	}

	if err := service.AkunRepository.Update(&akun); err != nil {
		exception.PanicIfError(err)
	}

	return response

}

func (service *akunServiceImpl) Delete(nip string) {
	akun, err := service.AkunRepository.FindByNIP(nip)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	if err := service.AkunRepository.Delete(&akun); err != nil {
		exception.PanicIfError(err)
	}
}

func NewAkunServiceProvider(repository *repository.AkunRepository) AkunService {
	return &akunServiceImpl{*repository}
}
