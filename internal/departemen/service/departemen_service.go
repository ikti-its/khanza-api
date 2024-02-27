package service

import "github.com/fathoor/simkes-api/internal/departemen/model"

type DepartemenService interface {
	Create(request *model.DepartemenRequest) model.DepartemenResponse
	GetAll() []model.DepartemenResponse
	GetByDepartemen(d string) model.DepartemenResponse
	Update(d string, request *model.DepartemenRequest) model.DepartemenResponse
	Delete(d string)
}
