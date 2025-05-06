package usecase

import (
	"fmt"
	"time"

	"github.com/ikti-its/khanza-api/internal/modules/reseppulang/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/reseppulang/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/reseppulang/internal/repository"
)

type PermintaanResepPulangUseCase struct {
	Repository repository.PermintaanResepPulangRepository
}

func NewPermintaanResepPulangUseCase(repo repository.PermintaanResepPulangRepository) *PermintaanResepPulangUseCase {
	return &PermintaanResepPulangUseCase{Repository: repo}
}

func (u *PermintaanResepPulangUseCase) Create(request *model.PermintaanResepPulangRequest) (model.PermintaanResepPulangResponse, error) {
	// Parse tgl_permintaan (nullable)
	var tglPermintaan *time.Time
	if request.TglPermintaan != "" {
		t, err := time.Parse("2006-01-02", request.TglPermintaan)
		if err != nil {
			return model.PermintaanResepPulangResponse{}, fmt.Errorf("invalid tgl_permintaan format: %v", err)
		}
		tglPermintaan = &t
	}

	jam, err := time.Parse("15:04:05", request.Jam)
	if err != nil {
		return model.PermintaanResepPulangResponse{}, fmt.Errorf("invalid jam format: %v", err)
	}

	tglValidasi, err := time.Parse("2006-01-02", request.TglValidasi)
	if err != nil {
		return model.PermintaanResepPulangResponse{}, fmt.Errorf("invalid tgl_validasi format: %v", err)
	}

	jamValidasi, err := time.Parse("15:04:05", request.JamValidasi)
	if err != nil {
		return model.PermintaanResepPulangResponse{}, fmt.Errorf("invalid jam_validasi format: %v", err)
	}

	data := entity.PermintaanResepPulang{
		NoPermintaan:  request.NoPermintaan,
		TglPermintaan: tglPermintaan,
		Jam:           jam,
		NoRawat:       request.NoRawat,
		KdDokter:      request.KdDokter,
		Status:        request.Status,
		TglValidasi:   tglValidasi,
		JamValidasi:   jamValidasi,
	}

	if err := u.Repository.Insert(&data); err != nil {
		return model.PermintaanResepPulangResponse{}, fmt.Errorf("failed to insert permintaan resep pulang: %v", err)
	}

	return model.PermintaanResepPulangResponse{
		NoPermintaan:  data.NoPermintaan,
		TglPermintaan: formatDatePtr(data.TglPermintaan),
		Jam:           data.Jam.Format("15:04:05"),
		NoRawat:       data.NoRawat,
		KdDokter:      data.KdDokter,
		Status:        data.Status,
		TglValidasi:   data.TglValidasi.Format("2006-01-02"),
		JamValidasi:   data.JamValidasi.Format("15:04:05"),
	}, nil
}

func (u *PermintaanResepPulangUseCase) GetAll() ([]model.PermintaanResepPulangResponse, error) {
	data, err := u.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	var result []model.PermintaanResepPulangResponse
	for _, p := range data {
		result = append(result, model.PermintaanResepPulangResponse{
			NoPermintaan:  p.NoPermintaan,
			TglPermintaan: formatDatePtr(p.TglPermintaan),
			Jam:           p.Jam.Format("15:04:05"),
			NoRawat:       p.NoRawat,
			KdDokter:      p.KdDokter,
			Status:        p.Status,
			TglValidasi:   p.TglValidasi.Format("2006-01-02"),
			JamValidasi:   p.JamValidasi.Format("15:04:05"),
		})
	}
	return result, nil
}

func (u *PermintaanResepPulangUseCase) GetByNoRawat(noRawat string) ([]model.PermintaanResepPulangResponse, error) {
	data, err := u.Repository.FindByNoRawat(noRawat)
	if err != nil {
		return nil, err
	}

	var result []model.PermintaanResepPulangResponse
	for _, p := range data {
		result = append(result, model.PermintaanResepPulangResponse{
			NoPermintaan:  p.NoPermintaan,
			TglPermintaan: formatDatePtr(p.TglPermintaan),
			Jam:           p.Jam.Format("15:04:05"),
			NoRawat:       p.NoRawat,
			KdDokter:      p.KdDokter,
			Status:        p.Status,
			TglValidasi:   p.TglValidasi.Format("2006-01-02"),
			JamValidasi:   p.JamValidasi.Format("15:04:05"),
		})
	}
	return result, nil
}

func (u *PermintaanResepPulangUseCase) GetByNoPermintaan(noPermintaan string) (model.PermintaanResepPulangResponse, error) {
	data, err := u.Repository.FindByNoPermintaan(noPermintaan)
	if err != nil {
		return model.PermintaanResepPulangResponse{}, fmt.Errorf("data not found")
	}

	return model.PermintaanResepPulangResponse{
		NoPermintaan:  data.NoPermintaan,
		TglPermintaan: formatDatePtr(data.TglPermintaan),
		Jam:           data.Jam.Format("15:04:05"),
		NoRawat:       data.NoRawat,
		KdDokter:      data.KdDokter,
		Status:        data.Status,
		TglValidasi:   data.TglValidasi.Format("2006-01-02"),
		JamValidasi:   data.JamValidasi.Format("15:04:05"),
	}, nil
}

func (u *PermintaanResepPulangUseCase) Update(noPermintaan string, request *model.PermintaanResepPulangRequest) (model.PermintaanResepPulangResponse, error) {
	data, err := u.Repository.FindByNoPermintaan(noPermintaan)
	if err != nil {
		return model.PermintaanResepPulangResponse{}, fmt.Errorf("data not found")
	}

	// Update fields
	if request.TglPermintaan != "" {
		t, err := time.Parse("2006-01-02", request.TglPermintaan)
		if err != nil {
			return model.PermintaanResepPulangResponse{}, fmt.Errorf("invalid tgl_permintaan format: %v", err)
		}
		data.TglPermintaan = &t
	}

	jam, err := time.Parse("15:04:05", request.Jam)
	if err != nil {
		return model.PermintaanResepPulangResponse{}, fmt.Errorf("invalid jam format: %v", err)
	}
	tglValidasi, err := time.Parse("2006-01-02", request.TglValidasi)
	if err != nil {
		return model.PermintaanResepPulangResponse{}, fmt.Errorf("invalid tgl_validasi format: %v", err)
	}
	jamValidasi, err := time.Parse("15:04:05", request.JamValidasi)
	if err != nil {
		return model.PermintaanResepPulangResponse{}, fmt.Errorf("invalid jam_validasi format: %v", err)
	}

	data.Jam = jam
	data.NoRawat = request.NoRawat
	data.KdDokter = request.KdDokter
	data.Status = request.Status
	data.TglValidasi = tglValidasi
	data.JamValidasi = jamValidasi

	if err := u.Repository.Update(data); err != nil {
		return model.PermintaanResepPulangResponse{}, fmt.Errorf("update failed: %v", err)
	}

	return model.PermintaanResepPulangResponse{
		NoPermintaan:  data.NoPermintaan,
		TglPermintaan: formatDatePtr(data.TglPermintaan),
		Jam:           data.Jam.Format("15:04:05"),
		NoRawat:       data.NoRawat,
		KdDokter:      data.KdDokter,
		Status:        data.Status,
		TglValidasi:   data.TglValidasi.Format("2006-01-02"),
		JamValidasi:   data.JamValidasi.Format("15:04:05"),
	}, nil
}

func (u *PermintaanResepPulangUseCase) Delete(noPermintaan string) error {
	return u.Repository.Delete(noPermintaan)
}

// Utilities
func formatDatePtr(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02")
}

// TODO: implement your own logic for this
func generateNoPermintaan() string {
	return fmt.Sprintf("PRP%s", time.Now().Format("20060102150405"))
}

func (u *PermintaanResepPulangUseCase) UpdateStatus(noPermintaan string, status string) (model.PermintaanResepPulangResponse, error) {
	fmt.Println("🚀 [DEBUG] UpdateStatus UseCase. no_permintaan:", noPermintaan)

	data, err := u.Repository.FindByNoPermintaan(noPermintaan)
	if err != nil {
		return model.PermintaanResepPulangResponse{}, fmt.Errorf("error while fetching data: %w", err)
	}
	if data == nil {
		return model.PermintaanResepPulangResponse{}, fmt.Errorf("data permintaan tidak ditemukan")
	}

	fmt.Println("✅ [DEBUG] Data ditemukan:", data)

	// Update status
	data.Status = status

	if err := u.Repository.Update(data); err != nil {
		return model.PermintaanResepPulangResponse{}, fmt.Errorf("failed to update status: %w", err)
	}

	return model.PermintaanResepPulangResponse{
		NoPermintaan:  data.NoPermintaan,
		TglPermintaan: formatDatePtr(data.TglPermintaan),
		Jam:           data.Jam.Format("15:04:05"),
		NoRawat:       data.NoRawat,
		KdDokter:      data.KdDokter,
		Status:        data.Status,
		TglValidasi:   data.TglValidasi.Format("2006-01-02"),
		JamValidasi:   data.JamValidasi.Format("15:04:05"),
	}, nil
}
