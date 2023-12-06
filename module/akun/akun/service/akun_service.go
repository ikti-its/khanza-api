package service

import "github.com/fathoor/simkes-api/module/akun/akun/model"

type AkunService interface {
	Create(request *model.AkunRequest) error
	GetAll() ([]model.AkunResponse, error)
	GetByNIP(nip string) (model.AkunResponse, error)
	Update(nip string, request *model.AkunRequest) (model.AkunResponse, error)
	Delete(nip string) error
}
