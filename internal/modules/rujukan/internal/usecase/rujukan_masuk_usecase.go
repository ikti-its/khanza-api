package usecase

import (
	"fmt"
	"time"

	"github.com/ikti-its/khanza-api/internal/modules/rujukan/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/rujukan/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/rujukan/internal/repository"
)

type RujukanMasukUseCase struct {
	Repository repository.RujukanMasukRepository
}

func NewRujukanMasukUseCase(repo repository.RujukanMasukRepository) *RujukanMasukUseCase {
	return &RujukanMasukUseCase{Repository: repo}
}

// Create a new RujukanMasuk entry
func (u *RujukanMasukUseCase) Create(request *model.RujukanMasukRequest) (model.RujukanMasukResponse, error) {
	tanggalMasuk, err := time.Parse("2006-01-02", request.TanggalMasuk)
	if err != nil {
		return model.RujukanMasukResponse{}, fmt.Errorf("invalid tanggal_masuk format: %v", err)
	}

	tanggalKeluar, err := time.Parse("2006-01-02", request.TanggalKeluar)
	if err != nil {
		return model.RujukanMasukResponse{}, fmt.Errorf("invalid tanggal_keluar format: %v", err)
	}

	entityData := entity.RujukanMasuk{
		NomorRujuk:    request.NomorRujuk,
		Perujuk:       request.Perujuk,
		AlamatPerujuk: request.AlamatPerujuk,
		NomorRawat:    request.NomorRawat,
		NomorRM:       request.NomorRM,
		NamaPasien:    request.NamaPasien,
		Alamat:        request.Alamat,
		Umur:          request.Umur,
		TanggalMasuk:  tanggalMasuk,
		TanggalKeluar: tanggalKeluar,
		DiagnosaAwal:  request.DiagnosaAwal,
	}

	if err := u.Repository.Insert(&entityData); err != nil {
		return model.RujukanMasukResponse{}, fmt.Errorf("failed to insert: %v", err)
	}

	return model.RujukanMasukResponse{
		NomorRujuk:    entityData.NomorRujuk,
		Perujuk:       entityData.Perujuk,
		AlamatPerujuk: entityData.AlamatPerujuk,
		NomorRawat:    entityData.NomorRawat,
		NomorRM:       entityData.NomorRM,
		NamaPasien:    entityData.NamaPasien,
		Alamat:        entityData.Alamat,
		Umur:          entityData.Umur,
		TanggalMasuk:  entityData.TanggalMasuk.Format("2006-01-02"),
		TanggalKeluar: entityData.TanggalKeluar.Format("2006-01-02"),
		DiagnosaAwal:  entityData.DiagnosaAwal,
	}, nil
}

// Get all RujukanMasuk entries
func (u *RujukanMasukUseCase) GetAll() ([]model.RujukanMasukResponse, error) {
	list, err := u.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	var response []model.RujukanMasukResponse
	for _, r := range list {
		response = append(response, model.RujukanMasukResponse{
			NomorRujuk:    r.NomorRujuk,
			Perujuk:       r.Perujuk,
			AlamatPerujuk: r.AlamatPerujuk,
			NomorRawat:    r.NomorRawat,
			NomorRM:       r.NomorRM,
			NamaPasien:    r.NamaPasien,
			Alamat:        r.Alamat,
			Umur:          r.Umur,
			TanggalMasuk:  r.TanggalMasuk.Format("2006-01-02"),
			TanggalKeluar: r.TanggalKeluar.Format("2006-01-02"),
			DiagnosaAwal:  r.DiagnosaAwal,
		})
	}
	return response, nil
}

// Get by nomor_rawat
func (u *RujukanMasukUseCase) GetByNomorRawat(nomorRawat string) (model.RujukanMasukResponse, error) {
	r, err := u.Repository.FindByNomorRawat(nomorRawat)
	if err != nil {
		return model.RujukanMasukResponse{}, err
	}

	return model.RujukanMasukResponse{
		NomorRujuk:    r.NomorRujuk,
		Perujuk:       r.Perujuk,
		AlamatPerujuk: r.AlamatPerujuk,
		NomorRawat:    r.NomorRawat,
		NomorRM:       r.NomorRM,
		NamaPasien:    r.NamaPasien,
		Alamat:        r.Alamat,
		Umur:          r.Umur,
		TanggalMasuk:  r.TanggalMasuk.Format("2006-01-02"),
		TanggalKeluar: r.TanggalKeluar.Format("2006-01-02"),
		DiagnosaAwal:  r.DiagnosaAwal,
	}, nil
}

// Update by nomor_rawat
func (u *RujukanMasukUseCase) Update(nomorRawat string, request *model.RujukanMasukRequest) (model.RujukanMasukResponse, error) {
	r, err := u.Repository.FindByNomorRawat(nomorRawat)
	if err != nil {
		return model.RujukanMasukResponse{}, fmt.Errorf("data not found")
	}

	r.NomorRujuk = request.NomorRujuk
	r.Perujuk = request.Perujuk
	r.AlamatPerujuk = request.AlamatPerujuk
	r.NomorRM = request.NomorRM
	r.NamaPasien = request.NamaPasien
	r.Alamat = request.Alamat
	r.Umur = request.Umur
	r.DiagnosaAwal = request.DiagnosaAwal

	r.TanggalMasuk, err = time.Parse("2006-01-02", request.TanggalMasuk)
	if err != nil {
		return model.RujukanMasukResponse{}, fmt.Errorf("invalid tanggal_masuk format")
	}

	r.TanggalKeluar, err = time.Parse("2006-01-02", request.TanggalKeluar)
	if err != nil {
		return model.RujukanMasukResponse{}, fmt.Errorf("invalid tanggal_keluar format")
	}

	if err := u.Repository.Update(&r); err != nil {
		return model.RujukanMasukResponse{}, fmt.Errorf("failed to update")
	}

	return u.GetByNomorRawat(nomorRawat)
}

// Delete by nomor_rawat
func (u *RujukanMasukUseCase) Delete(nomorRawat string) error {
	return u.Repository.Delete(nomorRawat)
}
