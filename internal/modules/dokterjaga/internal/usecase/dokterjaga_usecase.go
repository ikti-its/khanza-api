package usecase

import (
	"fmt"
	"time"

	"github.com/ikti-its/khanza-api/internal/modules/dokterjaga/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/dokterjaga/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/dokterjaga/internal/repository"
)

type DokterJagaUseCase struct {
	Repository repository.DokterJagaRepository
}

func NewDokterJagaUseCase(repo repository.DokterJagaRepository) *DokterJagaUseCase {
	return &DokterJagaUseCase{Repository: repo}
}

// Create a new dokter_jaga entry
func (u *DokterJagaUseCase) Create(request *model.DokterJagaRequest) (model.DokterJagaResponse, error) {
	// Parse the hari_kerja string to time.Time
	hariKerja, err := time.Parse("2006-01-02", request.HariKerja)
	if err != nil {
		return model.DokterJagaResponse{}, fmt.Errorf("invalid date format for hari_kerja, use YYYY-MM-DD: %v", err)
	}

	dokter := entity.DokterJaga{
		KodeDokter: request.KodeDokter,
		NamaDokter: request.NamaDokter,
		HariKerja:  hariKerja,
		JamMulai:   request.JamMulai,
		JamSelesai: request.JamSelesai,
		Poliklinik: request.Poliklinik,
		Status:     request.Status,
	}

	err = u.Repository.Insert(&dokter)
	if err != nil {
		return model.DokterJagaResponse{}, fmt.Errorf("failed to create dokter jaga: %v", err)
	}

	return model.DokterJagaResponse{
		KodeDokter: dokter.KodeDokter,
		NamaDokter: dokter.NamaDokter,
		HariKerja:  request.HariKerja, // still return string here for client
		JamMulai:   dokter.JamMulai,
		JamSelesai: dokter.JamSelesai,
		Poliklinik: dokter.Poliklinik,
		Status:     dokter.Status,
	}, nil
}

// Retrieve all dokter_jaga records
func (u *DokterJagaUseCase) GetAll() ([]model.DokterJagaResponse, error) {
	list, err := u.Repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve dokter jaga: %v", err)
	}

	var response []model.DokterJagaResponse
	for _, d := range list {
		response = append(response, model.DokterJagaResponse{
			KodeDokter: d.KodeDokter,
			NamaDokter: d.NamaDokter,
			HariKerja:  d.HariKerja.Format("2006-01-02"),
			JamMulai:   d.JamMulai,
			JamSelesai: d.JamSelesai,
			Poliklinik: d.Poliklinik,
			Status:     d.Status,
		})
	}
	return response, nil
}

// Get by kode_dokter
func (u *DokterJagaUseCase) GetByKodeDokter(kode string) ([]model.DokterJagaResponse, error) {
	list, err := u.Repository.FindByKodeDokter(kode)
	if err != nil {
		return nil, fmt.Errorf("dokter jaga not found")
	}

	var response []model.DokterJagaResponse
	for _, d := range list {
		response = append(response, model.DokterJagaResponse{
			KodeDokter: d.KodeDokter,
			NamaDokter: d.NamaDokter,
			HariKerja:  d.HariKerja.Format("2006-01-02"),
			JamMulai:   d.JamMulai,
			JamSelesai: d.JamSelesai,
			Poliklinik: d.Poliklinik,
			Status:     d.Status,
		})
	}
	return response, nil
}

// Update a record
func (u *DokterJagaUseCase) Update(request *model.DokterJagaRequest) (model.DokterJagaResponse, error) {
	// Ensure record exists
	records, err := u.Repository.FindByKodeDokter(request.KodeDokter)
	if err != nil || len(records) == 0 {
		return model.DokterJagaResponse{}, fmt.Errorf("dokter jaga not found")
	}

	// 🛠️ Parse the hari_kerja string
	hariKerja, err := time.Parse("2006-01-02", request.HariKerja)
	if err != nil {
		return model.DokterJagaResponse{}, fmt.Errorf("invalid format for hari_kerja: %v", err)
	}

	// Update the matched one
	dokter := entity.DokterJaga{
		KodeDokter: request.KodeDokter,
		NamaDokter: request.NamaDokter,
		HariKerja:  hariKerja,
		JamMulai:   request.JamMulai,
		JamSelesai: request.JamSelesai,
		Poliklinik: request.Poliklinik,
		Status:     request.Status,
	}

	err = u.Repository.Update(&dokter)
	if err != nil {
		return model.DokterJagaResponse{}, fmt.Errorf("failed to update dokter jaga: %v", err)
	}

	return model.DokterJagaResponse{
		KodeDokter: dokter.KodeDokter,
		NamaDokter: dokter.NamaDokter,
		HariKerja:  request.HariKerja, // returning string as expected by client
		JamMulai:   dokter.JamMulai,
		JamSelesai: dokter.JamSelesai,
		Poliklinik: dokter.Poliklinik,
		Status:     dokter.Status,
	}, nil
}

// Delete a dokter jaga shift by kode_dokter and hari_kerja
func (u *DokterJagaUseCase) Delete(kodeDokter, hariKerja string) error {
	return u.Repository.Delete(kodeDokter, hariKerja)
}

// Change status of a shift
func (u *DokterJagaUseCase) UpdateStatus(kodeDokter, hariKerja, status string) error {
	return u.Repository.UpdateStatus(kodeDokter, hariKerja, status)
}

// Get all with a certain status
func (u *DokterJagaUseCase) GetByStatus(status string) ([]model.DokterJagaResponse, error) {
	list, err := u.Repository.FindByStatus(status)
	if err != nil {
		return nil, fmt.Errorf("failed to find dokter jaga with status: %v", err)
	}

	var response []model.DokterJagaResponse
	for _, d := range list {
		response = append(response, model.DokterJagaResponse{
			KodeDokter: d.KodeDokter,
			NamaDokter: d.NamaDokter,
			HariKerja:  d.HariKerja.Format("2006-01-02"),
			JamMulai:   d.JamMulai,
			JamSelesai: d.JamSelesai,
			Poliklinik: d.Poliklinik,
			Status:     d.Status,
		})
	}
	return response, nil
}
