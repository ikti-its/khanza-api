package service

import (
	"github.com/fathoor/simkes-api/internal/akun/repository"
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/app/helper"
	"github.com/fathoor/simkes-api/internal/auth/model"
	"github.com/fathoor/simkes-api/internal/auth/validation"
	"time"
)

type authServiceImpl struct {
	repository.AkunRepository
}

func (service *authServiceImpl) Login(request *model.AuthRequest) model.AuthResponse {
	if valid := validation.ValidateAuthRequest(request); valid != nil {
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

	token, err := helper.GenerateJWT(akun.NIP, akun.RoleNama)
	exception.PanicIfError(err)

	response := model.AuthResponse{
		Token:   token,
		Expired: time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05"),
	}

	return response
}

func NewAuthServiceProvider(repository *repository.AkunRepository) AuthService {
	return &authServiceImpl{*repository}
}
