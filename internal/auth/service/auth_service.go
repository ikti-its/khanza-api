package service

import "github.com/fathoor/simkes-api/internal/auth/model"

type AuthService interface {
	Login(request *model.AuthRequest) model.AuthResponse
}
