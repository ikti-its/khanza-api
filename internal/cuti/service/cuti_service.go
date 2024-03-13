package service

import "github.com/fathoor/simkes-api/internal/cuti/model"

type CutiService interface {
	Create(request *model.CutiCreateRequest) model.CutiResponse
	GetAll() []model.CutiResponse
	GetByNIP(nip string) []model.CutiResponse
	GetByID(id string) model.CutiResponse
	Update(id string, request *model.CutiUpdateRequest) model.CutiResponse
	UpdateStatus(id string, request *model.CutiUpdateRequest) model.CutiResponse
	Delete(id string)
}
