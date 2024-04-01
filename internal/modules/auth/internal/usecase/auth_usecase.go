package usecase

import (
	"github.com/fathoor/simkes-api/internal/app/config"
	"github.com/fathoor/simkes-api/internal/app/exception"
	"github.com/fathoor/simkes-api/internal/app/helper"
	"github.com/fathoor/simkes-api/internal/modules/auth/internal/model"
	"github.com/fathoor/simkes-api/internal/modules/auth/internal/repository"
)

type AuthUseCase struct {
	Repository repository.AuthRepository
	Config     *config.Config
}

func NewAuthUseCase(repository *repository.AuthRepository, cfg *config.Config) *AuthUseCase {
	return &AuthUseCase{
		Repository: *repository,
		Config:     cfg,
	}
}

func (u *AuthUseCase) Login(request *model.AuthRequest) model.AuthResponse {
	akun, err := u.Repository.FindByEmail(request.Email)
	if err != nil {
		panic(&exception.NotFoundError{
			Message: "Akun not found",
		})
	}

	if !helper.DecryptPassword(akun.Password, request.Password) {
		panic(&exception.UnauthorizedError{
			Message: "Invalid password",
		})
	}

	token, err := helper.GenerateJWT(akun.Id, akun.Role, u.Config)
	exception.PanicIfError(err, "Failed to generate JWT")

	response := model.AuthResponse{
		Token: token,
	}

	return response
}
