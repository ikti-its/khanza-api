package usecase

import (
	"fmt"
	"time"

	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/repository"
)

type CatatanObservasiRanapUseCase struct {
	Repository repository.CatatanObservasiRanapRepository
}

func NewCatatanObservasiRanapUseCase(repo repository.CatatanObservasiRanapRepository) *CatatanObservasiRanapUseCase {
	return &CatatanObservasiRanapUseCase{Repository: repo}
}

func (u *CatatanObservasiRanapUseCase) Create(request *model.CatatanObservasiRanapRequest) (*model.CatatanObservasiRanapResponse, error) {
	fmt.Println("DEBUG jam_rawat (request.Jam):", request.Jam)
	tgl, err := time.Parse("2006-01-02", request.Tanggal)
	if err != nil {
		return nil, fmt.Errorf("format tanggal salah: %v", err)
	}
	jam := request.Jam

	data := entity.CatatanObservasiRanap{
		NoRawat:      request.NoRawat,
		TglPerawatan: &tgl,
		JamRawat:     jam,
		GCS:          toPtr(request.GCS),
		TD:           request.TD,
		HR:           toPtr(request.HR),
		RR:           toPtr(request.RR),
		Suhu:         toPtr(request.Suhu),
		Spo2:         request.Spo2,
		NIP:          request.NIP,
	}

	if err := u.Repository.Insert(&data); err != nil {
		return nil, fmt.Errorf("gagal insert catatan observasi: %v", err)
	}

	return &model.CatatanObservasiRanapResponse{
		NoRawat: data.NoRawat,
		Tanggal: tgl.Format("2006-01-02"),
		Jam:     data.JamRawat,
		GCS:     getString(data.GCS),
		TD:      data.TD,
		HR:      getString(data.HR),
		RR:      getString(data.RR),
		Suhu:    getString(data.Suhu),
		Spo2:    data.Spo2,
		NIP:     data.NIP,
	}, nil
}

func (u *CatatanObservasiRanapUseCase) GetAll() ([]model.CatatanObservasiRanapResponse, error) {
	records, err := u.Repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data observasi: %v", err)
	}

	var responses []model.CatatanObservasiRanapResponse
	for _, r := range records {
		responses = append(responses, model.CatatanObservasiRanapResponse{
			NoRawat: r.NoRawat,
			Tanggal: formatDate(r.TglPerawatan),
			Jam:     r.JamRawat,
			GCS:     getString(r.GCS),
			TD:      r.TD,
			HR:      getString(r.HR),
			RR:      getString(r.RR),
			Suhu:    getString(r.Suhu),
			Spo2:    r.Spo2,
			NIP:     r.NIP,
		})
	}
	return responses, nil
}

func (u *CatatanObservasiRanapUseCase) GetByNoRawat(noRawat string) ([]model.CatatanObservasiRanapResponse, error) {
	records, err := u.Repository.FindByNoRawat(noRawat)
	if err != nil {
		return nil, fmt.Errorf("catatan tidak ditemukan: %v", err)
	}

	var responses []model.CatatanObservasiRanapResponse
	for _, r := range records {
		responses = append(responses, model.CatatanObservasiRanapResponse{
			NoRawat: r.NoRawat,
			Tanggal: formatDate(r.TglPerawatan),
			Jam:     r.JamRawat,
			GCS:     getString(r.GCS),
			TD:      r.TD,
			HR:      getString(r.HR),
			RR:      getString(r.RR),
			Suhu:    getString(r.Suhu),
			Spo2:    r.Spo2,
			NIP:     r.NIP,
		})
	}
	return responses, nil
}

func (u *CatatanObservasiRanapUseCase) Update(request *model.CatatanObservasiRanapRequest) error {
	tgl, err := time.Parse("2006-01-02", request.Tanggal)
	if err != nil {
		return fmt.Errorf("format tanggal salah: %v", err)
	}

	entity := entity.CatatanObservasiRanap{
		NoRawat:      request.NoRawat,
		TglPerawatan: &tgl,
		JamRawat:     request.Jam,
		GCS:          toPtr(request.GCS),
		TD:           request.TD,
		HR:           toPtr(request.HR),
		RR:           toPtr(request.RR),
		Suhu:         toPtr(request.Suhu),
		Spo2:         request.Spo2,
		NIP:          request.NIP,
	}

	return u.Repository.Update(&entity)
}

func (u *CatatanObservasiRanapUseCase) Delete(noRawat string, tanggal string, jam string) error {
	return u.Repository.Delete(noRawat, tanggal, jam)
}

// Helper functions
func toPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func formatDate(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02")
}
