package service

import "github.com/fathoor/simkes-api/internal/kehadiran/model"

type KehadiranService interface {
	CheckIn(request *model.KehadiranRequest) model.KehadiranResponse
	CheckOut(request *model.KehadiranRequest) model.KehadiranResponse
	GetAll() []model.KehadiranResponse
	GetByNIP(nip string) []model.KehadiranResponse
	GetByID(id string) model.KehadiranResponse
	Update(id string, request *model.KehadiranUpdateRequest) model.KehadiranResponse
	Delete(id string)
}
