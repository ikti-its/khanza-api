package service

import "github.com/fathoor/simkes-api/module/auth/model"

type AuthService interface {
	Login(request *model.AuthRequest) (model.AuthResponse, error)
}
