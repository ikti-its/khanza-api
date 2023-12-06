package service

import (
	"github.com/fathoor/simkes-api/core/exception"
	"github.com/fathoor/simkes-api/core/helper"
	"github.com/fathoor/simkes-api/module/akun/akun/repository"
	"github.com/fathoor/simkes-api/module/auth/model"
	"github.com/fathoor/simkes-api/module/auth/validation"
	"time"
)

type authServiceImpl struct {
	repository.AkunRepository
}

func (service *authServiceImpl) Login(request *model.AuthRequest) (model.AuthResponse, error) {
	valid := validation.ValidateAuthRequest(request)
	if valid != nil {
		panic(exception.BadRequestError{
			Message: "Invalid request data",
		})
	}

	akun, err := service.AkunRepository.FindByNIP(request.NIP)
	if err != nil {
		panic(exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	if !helper.DecryptPassword(akun.Password, request.Password) {
		panic(exception.UnauthorizedError{
			Message: "Invalid password",
		})
	}

	token, err := helper.GenerateJWT(akun.NIP, akun.RoleID)
	exception.PanicIfError(err)

	response := model.AuthResponse{
		Token:   token,
		Type:    "Bearer",
		Expired: time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05"),
	}

	return response, err
}

func ProvideAuthService(repository *repository.AkunRepository) AuthService {
	return &authServiceImpl{*repository}
}
