package repository

import "github.com/ikti-its/khanza-api/internal/modules/rekammedis/internal/entity"

type DiagnosaPasienRepository interface {
	Insert(diagnosa *entity.DiagnosaPasien) error
	FindAll() ([]entity.DiagnosaPasien, error)
	FindByNoRawat(noRawat string) ([]entity.DiagnosaPasien, error)
	FindByKodePenyakit(kode string) ([]entity.DiagnosaPasien, error)
	FindByNoRawatAndStatus(noRawat string, status string) ([]entity.DiagnosaPasien, error)
	Update(diagnosa *entity.DiagnosaPasien) error
	Delete(noRawat string, kdPenyakit string) error
}
