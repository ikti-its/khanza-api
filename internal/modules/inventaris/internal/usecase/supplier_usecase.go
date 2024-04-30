package usecase

import (
	"github.com/ikti-its/khanza-api/internal/app/exception"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/repository"
)

type SupplierUseCase struct {
	Repository repository.SupplierRepository
}

func NewSupplierUseCase(repository *repository.SupplierRepository) *SupplierUseCase {
	return &SupplierUseCase{
		Repository: *repository,
	}
}

func (u *SupplierUseCase) Get() []model.SupplierResponse {
	supplier, err := u.Repository.Find()
	exception.PanicIfError(err, "Failed to get all supplier")

	response := make([]model.SupplierResponse, len(supplier))
	for i, supplier := range supplier {
		response[i] = model.SupplierResponse{
			Id:       supplier.Id,
			Nama:     supplier.Nama,
			Alamat:   supplier.Alamat,
			Telepon:  supplier.Telepon,
			Kota:     supplier.Kota,
			Bank:     supplier.Bank,
			Rekening: supplier.Rekening,
		}
	}

	return response
}
