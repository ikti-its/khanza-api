package service

import "github.com/fathoor/simkes-api/internal/jabatan/model"

type JabatanService interface {
	Create(request *model.JabatanRequest) model.JabatanResponse
	GetAll() []model.JabatanResponse
	GetByJabatan(j string) model.JabatanResponse
	Update(j string, request *model.JabatanRequest) model.JabatanResponse
	Delete(j string)
}
