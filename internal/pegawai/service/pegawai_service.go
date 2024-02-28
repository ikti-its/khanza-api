package service

import "github.com/fathoor/simkes-api/internal/pegawai/model"

type PegawaiService interface {
	Create(request *model.PegawaiRequest) model.PegawaiResponse
	GetAll() []model.PegawaiResponse
	GetPage(page, size int) model.PegawaiPageResponse
	GetByNIP(nip string) model.PegawaiResponse
	Update(nip string, request *model.PegawaiRequest) model.PegawaiResponse
	Delete(nip string)
}
