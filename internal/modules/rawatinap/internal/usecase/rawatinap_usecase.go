package usecase

import (
	"fmt"
	"time"

	"github.com/ikti-its/khanza-api/internal/modules/rawatinap/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/rawatinap/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/rawatinap/internal/repository"
)

type RawatInapUseCase struct {
	Repository repository.RawatInapRepository
}

func NewRawatInapUseCase(repo repository.RawatInapRepository) *RawatInapUseCase {
	return &RawatInapUseCase{Repository: repo}
}

// Create a new rawat_inap record
func (u *RawatInapUseCase) Create(request *model.RawatInapRequest) (model.RawatInapResponse, error) {
	tglMasuk, err := time.Parse("2006-01-02", request.TanggalMasuk)
	if err != nil {
		return model.RawatInapResponse{}, fmt.Errorf("invalid tanggal_masuk format: %v", err)
	}
	jamMasuk, err := time.Parse("15:04:05", request.JamMasuk)
	if err != nil {
		return model.RawatInapResponse{}, fmt.Errorf("invalid jam_masuk format: %v", err)
	}
	var tanggalKeluar time.Time
	var jamKeluar time.Time

	if request.TanggalKeluar != "" {
		tanggalKeluar, err = time.Parse("2006-01-02", request.TanggalKeluar)
		if err != nil {
			return model.RawatInapResponse{}, fmt.Errorf("invalid tanggal_keluar format: %v", err)
		}
	}

	if request.JamKeluar != "" {
		jamKeluar, err = time.Parse("15:04:05", request.JamKeluar)
		if err != nil {
			return model.RawatInapResponse{}, fmt.Errorf("invalid jam_keluar format: %v", err)
		}
	}

	rawat := entity.RawatInap{
		NomorRawat:      request.NomorRawat,
		NomorRM:         request.NomorRM,
		NamaPasien:      request.NamaPasien,
		AlamatPasien:    request.AlamatPasien,
		PenanggungJawab: request.PenanggungJawab,
		HubunganPJ:      request.HubunganPJ,
		JenisBayar:      request.JenisBayar,
		Kamar:           request.Kamar,
		TarifKamar:      request.TarifKamar,
		DiagnosaAwal:    request.DiagnosaAwal,
		DiagnosaAkhir:   request.DiagnosaAkhir,
		TanggalMasuk:    tglMasuk,
		JamMasuk:        jamMasuk,
		TanggalKeluar:   tanggalKeluar,
		JamKeluar:       jamKeluar,
		TotalBiaya:      request.TotalBiaya,
		StatusPulang:    request.StatusPulang,
		LamaRanap:       request.LamaRanap,
		DokterPJ:        request.DokterPJ,
		StatusBayar:     request.StatusBayar,
	}

	err = u.Repository.Insert(&rawat)
	if err != nil {
		return model.RawatInapResponse{}, fmt.Errorf("failed to create rawat inap: %v", err)
	}

	return model.FromEntity(rawat), nil
}

func (u *RawatInapUseCase) GetAll() ([]model.RawatInapResponse, error) {
	list, err := u.Repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve rawat inap records: %v", err)
	}

	var result []model.RawatInapResponse
	for _, item := range list {
		result = append(result, model.FromEntity(item))
	}
	return result, nil
}

func (u *RawatInapUseCase) GetByNomorRawat(nomorRawat string) (model.RawatInapResponse, error) {
	rawat, err := u.Repository.FindByNomorRawat(nomorRawat)
	if err != nil {
		return model.RawatInapResponse{}, fmt.Errorf("rawat inap not found")
	}
	return model.FromEntity(rawat), nil
}

func (u *RawatInapUseCase) Update(nomorRawat string, request *model.RawatInapRequest) (model.RawatInapResponse, error) {
	rawat, err := u.Repository.FindByNomorRawat(nomorRawat)
	if err != nil {
		return model.RawatInapResponse{}, fmt.Errorf("rawat inap not found")
	}

	// Optional parsing
	if request.TanggalMasuk != "" {
		rawat.TanggalMasuk, _ = time.Parse("2006-01-02", request.TanggalMasuk)
	}
	if request.JamMasuk != "" {
		rawat.JamMasuk, _ = time.Parse("15:04:05", request.JamMasuk)
	}
	var tanggalKeluar time.Time
	var jamKeluar time.Time

	if request.TanggalKeluar != "" {
		tanggalKeluar, err = time.Parse("2006-01-02", request.TanggalKeluar)
		if err != nil {
			return model.RawatInapResponse{}, fmt.Errorf("invalid tanggal_keluar format: %v", err)
		}
	}

	if request.JamKeluar != "" {
		jamKeluar, err = time.Parse("15:04:05", request.JamKeluar)
		if err != nil {
			return model.RawatInapResponse{}, fmt.Errorf("invalid jam_keluar format: %v", err)
		}
	}

	rawat.NamaPasien = request.NamaPasien
	rawat.AlamatPasien = request.AlamatPasien
	rawat.PenanggungJawab = request.PenanggungJawab
	rawat.HubunganPJ = request.HubunganPJ
	rawat.JenisBayar = request.JenisBayar
	rawat.Kamar = request.Kamar
	rawat.TarifKamar = request.TarifKamar
	rawat.DiagnosaAwal = request.DiagnosaAwal
	rawat.DiagnosaAkhir = request.DiagnosaAkhir
	rawat.TanggalKeluar = tanggalKeluar
	rawat.JamKeluar = jamKeluar
	rawat.TotalBiaya = request.TotalBiaya
	rawat.StatusPulang = request.StatusPulang
	rawat.LamaRanap = request.LamaRanap
	rawat.DokterPJ = request.DokterPJ
	rawat.StatusBayar = request.StatusBayar

	err = u.Repository.Update(&rawat)
	if err != nil {
		return model.RawatInapResponse{}, fmt.Errorf("failed to update rawat inap: %v", err)
	}

	return model.FromEntity(rawat), nil
}

func (u *RawatInapUseCase) Delete(nomorRawat string) error {
	err := u.Repository.Delete(nomorRawat)
	if err != nil {
		return fmt.Errorf("failed to delete rawat inap: %v", err)
	}
	return nil
}
