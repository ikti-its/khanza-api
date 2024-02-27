package service

import (
	"github.com/fathoor/simkes-api/internal/akun/model"
)

type AkunService interface {
	Create(request *model.AkunRequest) model.AkunResponse
	GetAll() []model.AkunResponse
	GetPage(page, size int) model.AkunPageResponse
	GetByNIP(nip string) model.AkunResponse
	Update(nip string, request *model.AkunRequest) model.AkunResponse
	UpdateAdmin(nip string, request *model.AkunRequest) model.AkunResponse
	Delete(nip string)
}
