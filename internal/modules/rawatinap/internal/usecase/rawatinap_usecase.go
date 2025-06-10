package usecase

import (
	"database/sql"
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
		AlamatPasien:    sql.NullString{String: request.AlamatPasien, Valid: request.AlamatPasien != ""},
		PenanggungJawab: sql.NullString{String: request.PenanggungJawab, Valid: request.PenanggungJawab != ""},
		HubunganPJ:      sql.NullString{String: request.HubunganPJ, Valid: request.HubunganPJ != ""},
		JenisBayar:      sql.NullString{String: request.JenisBayar, Valid: request.JenisBayar != ""},
		Kamar:           sql.NullString{String: request.Kamar, Valid: request.Kamar != ""},
		TarifKamar:      sql.NullFloat64{Float64: request.TarifKamar, Valid: true}, // adjust condition if needed
		DiagnosaAwal:    sql.NullString{String: request.DiagnosaAwal, Valid: request.DiagnosaAwal != ""},
		DiagnosaAkhir:   sql.NullString{String: request.DiagnosaAkhir, Valid: request.DiagnosaAkhir != ""},
		TanggalMasuk:    tglMasuk,
		JamMasuk:        sql.NullTime{Time: jamMasuk, Valid: true},
		TanggalKeluar:   sql.NullTime{Time: tanggalKeluar, Valid: true}, // adjust if nullable
		JamKeluar:       sql.NullTime{Time: jamKeluar, Valid: true},     // adjust if nullable
		TotalBiaya:      sql.NullFloat64{Float64: request.TotalBiaya, Valid: true},
		StatusPulang:    sql.NullString{String: request.StatusPulang, Valid: request.StatusPulang != ""},
		LamaRanap:       sql.NullFloat64{Float64: request.LamaRanap, Valid: true},
		DokterPJ:        sql.NullString{String: request.DokterPJ, Valid: request.DokterPJ != ""},
		StatusBayar:     sql.NullString{String: request.StatusBayar, Valid: request.StatusBayar != ""},
	}

	err = u.Repository.Insert(&rawat)
	if err != nil {
		return model.RawatInapResponse{}, fmt.Errorf("failed to create rawat inap: %v", err)
	}

	return model.FromEntity(rawat), nil
}

func (u *RawatInapUseCase) GetAll() ([]model.RawatInapResponse, error) {
	fmt.Println("📥 Fetching all rawat inap records...")

	rawats, err := u.Repository.FindAll()
	if err != nil {
		fmt.Println("❌ DB error in FindAll():", err)
		return nil, err
	}

	fmt.Println("✅ Retrieved records:", len(rawats))

	// Optional: log one sample
	if len(rawats) > 0 {
		fmt.Printf("🔍 First record: %+v\n", rawats[0])
	}

	// Transform to response model
	// If you're using a `FromEntity()` mapper, watch for panics there too
	var responses []model.RawatInapResponse
	for _, r := range rawats {
		responses = append(responses, model.FromEntity(r)) // <- this might panic if any value is nil
	}

	return responses, nil
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
		parsedJamMasuk, err := time.Parse("15:04:05", request.JamMasuk)
		if err == nil {
			rawat.JamMasuk = sql.NullTime{Time: parsedJamMasuk, Valid: true}
		} else {
			rawat.JamMasuk = sql.NullTime{Valid: false}
		}
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
	rawat.AlamatPasien = sql.NullString{String: request.AlamatPasien, Valid: request.AlamatPasien != ""}
	rawat.PenanggungJawab = sql.NullString{String: request.PenanggungJawab, Valid: request.PenanggungJawab != ""}
	rawat.HubunganPJ = sql.NullString{String: request.HubunganPJ, Valid: request.HubunganPJ != ""}
	rawat.JenisBayar = sql.NullString{String: request.JenisBayar, Valid: request.JenisBayar != ""}
	rawat.Kamar = sql.NullString{String: request.Kamar, Valid: request.Kamar != ""}
	rawat.TarifKamar = sql.NullFloat64{Float64: request.TarifKamar, Valid: true}
	rawat.DiagnosaAwal = sql.NullString{String: request.DiagnosaAwal, Valid: request.DiagnosaAwal != ""}
	rawat.DiagnosaAkhir = sql.NullString{String: request.DiagnosaAkhir, Valid: request.DiagnosaAkhir != ""}
	rawat.TanggalKeluar = sql.NullTime{Time: tanggalKeluar, Valid: true}
	rawat.JamKeluar = sql.NullTime{Time: jamKeluar, Valid: true}
	rawat.TotalBiaya = sql.NullFloat64{Float64: request.TotalBiaya, Valid: true}
	rawat.StatusPulang = sql.NullString{String: request.StatusPulang, Valid: request.StatusPulang != ""}
	rawat.LamaRanap = sql.NullFloat64{Float64: request.LamaRanap, Valid: true}
	rawat.DokterPJ = sql.NullString{String: request.DokterPJ, Valid: request.DokterPJ != ""}
	rawat.StatusBayar = sql.NullString{String: request.StatusBayar, Valid: request.StatusBayar != ""}

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
