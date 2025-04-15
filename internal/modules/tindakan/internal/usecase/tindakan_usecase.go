package usecase

import (
	"fmt"
	"time"

	"github.com/ikti-its/khanza-api/internal/modules/tindakan/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/tindakan/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/tindakan/internal/repository"
)

type TindakanUseCase struct {
	Repository repository.TindakanRepository
}

func ptrStr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func ptrInt(i int64) *int64 {
	if i == 0 {
		return nil
	}
	return &i
}

func derefStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func derefInt(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

func NewTindakanUseCase(repo repository.TindakanRepository) *TindakanUseCase {
	return &TindakanUseCase{Repository: repo}
}

// Create a new tindakan entry
func (u *TindakanUseCase) Create(request *model.TindakanRequest) (model.TindakanResponse, error) {
	exists, err := u.Repository.CheckDokterExists(request.KodeDokter)
	if err != nil {
		return model.TindakanResponse{}, fmt.Errorf("database error: %v", err)
	}
	if !exists {
		return model.TindakanResponse{}, fmt.Errorf("dokter with kode_dokter '%s' does not exist", request.KodeDokter)
	}

	tgl, err := time.Parse("2006-01-02", request.TanggalRawat)
	if err != nil {
		return model.TindakanResponse{}, fmt.Errorf("invalid tanggal_rawat format: %v", err)
	}
	jam, err := time.Parse("15:04:05", request.JamRawat)
	if err != nil {
		return model.TindakanResponse{}, fmt.Errorf("invalid jam_rawat format: %v", err)
	}

	tindakan := entity.Tindakan{
		NomorRawat:   request.NomorRawat,
		NomorRM:      request.NomorRM,
		NamaPasien:   request.NamaPasien,
		Tindakan:     ptrStr(request.Tindakan),
		KodeDokter:   ptrStr(request.KodeDokter),
		NamaDokter:   ptrStr(request.NamaDokter),
		NIP:          ptrStr(request.NIP),
		NamaPetugas:  ptrStr(request.NamaPetugas),
		TanggalRawat: tgl,
		JamRawat:     jam,
		Biaya:        ptrInt(int64(request.Biaya)),
	}

	err = u.Repository.Insert(&tindakan)
	if err != nil {
		return model.TindakanResponse{}, fmt.Errorf("failed to insert tindakan: %v", err)
	}

	return model.TindakanResponse{
		NomorRawat:   tindakan.NomorRawat,
		NomorRM:      tindakan.NomorRM,
		NamaPasien:   tindakan.NamaPasien,
		Tindakan:     derefStr(tindakan.Tindakan),
		KodeDokter:   derefStr(tindakan.KodeDokter),
		NamaDokter:   derefStr(tindakan.NamaDokter),
		NIP:          derefStr(tindakan.NIP),
		NamaPetugas:  derefStr(tindakan.NamaPetugas),
		TanggalRawat: tgl.Format("2006-01-02"),
		JamRawat:     jam.Format("15:04:05"),
		Biaya:        float64(derefInt(tindakan.Biaya)),
	}, nil
}

func (u *TindakanUseCase) GetAll() ([]model.TindakanResponse, error) {
	list, err := u.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	var response []model.TindakanResponse
	for _, t := range list {
		response = append(response, model.TindakanResponse{
			NomorRawat:   t.NomorRawat,
			NomorRM:      t.NomorRM,
			NamaPasien:   t.NamaPasien,
			Tindakan:     derefStr(t.Tindakan),
			KodeDokter:   derefStr(t.KodeDokter),
			NamaDokter:   derefStr(t.NamaDokter),
			NIP:          derefStr(t.NIP),
			NamaPetugas:  derefStr(t.NamaPetugas),
			TanggalRawat: t.TanggalRawat.Format("2006-01-02"),
			JamRawat:     t.JamRawat.Format("15:04:05"),
			Biaya:        float64(derefInt(t.Biaya)),
		})
	}
	return response, nil
}

func (u *TindakanUseCase) GetByNomorRawat(nomorRawat string) ([]model.TindakanResponse, error) {
	records, err := u.Repository.FindByNomorRawat(nomorRawat)
	if err != nil {
		return nil, err
	}

	var result []model.TindakanResponse
	for _, t := range records {
		result = append(result, model.TindakanResponse{
			NomorRawat:   t.NomorRawat,
			NomorRM:      t.NomorRM,
			NamaPasien:   t.NamaPasien,
			Tindakan:     derefStr(t.Tindakan),
			KodeDokter:   derefStr(t.KodeDokter),
			NamaDokter:   derefStr(t.NamaDokter),
			NIP:          derefStr(t.NIP),
			NamaPetugas:  derefStr(t.NamaPetugas),
			TanggalRawat: t.TanggalRawat.Format("2006-01-02"),
			JamRawat:     t.JamRawat.Format("15:04:05"),
			Biaya:        float64(derefInt(t.Biaya)),
		})
	}

	return result, nil
}

func (u *TindakanUseCase) Update(nomorRawat string, request *model.TindakanRequest) (model.TindakanResponse, error) {
	records, err := u.Repository.FindByNomorRawat(nomorRawat)
	if err != nil || len(records) == 0 {
		return model.TindakanResponse{}, fmt.Errorf("tindakan not found")
	}

	existing := &records[0] // just take the first one if updating one
	if err != nil {
		return model.TindakanResponse{}, fmt.Errorf("tindakan not found")
	}

	tgl, err := time.Parse("2006-01-02", request.TanggalRawat)
	if err != nil {
		return model.TindakanResponse{}, fmt.Errorf("invalid tanggal_rawat format: %v", err)
	}
	jam, err := time.Parse("15:04:05", request.JamRawat)
	if err != nil {
		return model.TindakanResponse{}, fmt.Errorf("invalid jam_rawat format: %v", err)
	}

	existing.NamaPasien = request.NamaPasien
	existing.Tindakan = ptrStr(request.Tindakan)
	existing.KodeDokter = ptrStr(request.KodeDokter)
	existing.NamaDokter = ptrStr(request.NamaDokter)
	existing.NIP = ptrStr(request.NIP)
	existing.NamaPetugas = ptrStr(request.NamaPetugas)
	existing.TanggalRawat = tgl
	existing.JamRawat = jam
	existing.Biaya = ptrInt(int64(request.Biaya))

	err = u.Repository.Update(existing)
	if err != nil {
		return model.TindakanResponse{}, fmt.Errorf("update failed: %v", err)
	}

	return model.TindakanResponse{
		NomorRawat:   existing.NomorRawat,
		NomorRM:      existing.NomorRM,
		NamaPasien:   existing.NamaPasien,
		Tindakan:     derefStr(existing.Tindakan),
		KodeDokter:   derefStr(existing.KodeDokter),
		NamaDokter:   derefStr(existing.NamaDokter),
		NIP:          derefStr(existing.NIP),
		NamaPetugas:  derefStr(existing.NamaPetugas),
		TanggalRawat: existing.TanggalRawat.Format("2006-01-02"),
		JamRawat:     existing.JamRawat.Format("15:04:05"),
		Biaya:        float64(derefInt(existing.Biaya)),
	}, nil
}

// Delete tindakan
func (u *TindakanUseCase) Delete(nomorRawat, jamRawat string) error {
	return u.Repository.Delete(nomorRawat, jamRawat)
}
