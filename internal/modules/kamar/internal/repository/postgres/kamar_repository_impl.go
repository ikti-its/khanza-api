package postgres

import (
	"github.com/ikti-its/khanza-api/internal/modules/kamar/internal/entity"
	"github.com/jmoiron/sqlx"
)

type KamarRepository interface {
	Insert(kamar *entity.Kamar) error
	Find() ([]entity.Kamar, error)
	FindAll() ([]entity.Kamar, error)
	FindByNomorBed(nomorReg string) (entity.Kamar, error)
	FindByKodeKamar(nomorReg string) (entity.Kamar, error)
	Update(kamar *entity.Kamar) error
	Delete(nomorReg string) error
	GetAvailableRooms() ([]entity.Kamar, error)
	UpdateStatusKamar(nomorBed string, status string) error
}

type kamarRepositoryImpl struct {
	DB *sqlx.DB
}

func NewKamarRepository(db *sqlx.DB) KamarRepository {
	return &kamarRepositoryImpl{DB: db}
}

func (r *kamarRepositoryImpl) Insert(kamar *entity.Kamar) error {
	query := `
		INSERT INTO kamar (
			nomor_bed, kode_kamar, nama_kamar, kelas, tarif_kamar, status_kamar
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`
	_, err := r.DB.Exec(query,
		kamar.NomorBed,    // Corresponds to nomor_bed
		kamar.KodeKamar,   // Corresponds to kode_kamar
		kamar.NamaKamar,   // Corresponds to nama_kamar
		kamar.Kelas,       // Corresponds to kelas
		kamar.TarifKamar,  // Corresponds to tarif_kamar
		kamar.StatusKamar, // Corresponds to status_kamar
	)
	return err
}

func (r *kamarRepositoryImpl) Find() ([]entity.Kamar, error) {
	query := `
		SELECT * FROM kamar ORDER BY nomor_bed DESC
	`
	var records []entity.Kamar
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *kamarRepositoryImpl) FindAll() ([]entity.Kamar, error) {
	query := `
		SELECT * FROM kamar ORDER BY nomor_bed DESC
	`
	var records []entity.Kamar
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *kamarRepositoryImpl) FindByNomorBed(nomorBed string) (entity.Kamar, error) {
	query := `SELECT * FROM kamar WHERE nomor_bed = $1`

	var record entity.Kamar
	err := r.DB.Get(&record, query, nomorBed)
	return record, err
}

func (r *kamarRepositoryImpl) FindByKodeKamar(nomorBed string) (entity.Kamar, error) {
	query := `
		SELECT * FROM kamar WHERE kode_kamar = $1
	`
	var record entity.Kamar
	err := r.DB.Get(&record, query, nomorBed)
	return record, err
}

func (r *kamarRepositoryImpl) Update(kamar *entity.Kamar) error {
	query := `
		UPDATE kamar SET 
			kode_kamar = $2, nama_kamar = $3, kelas = $4, tarif_kamar = $5, status_kamar = $6
		WHERE nomor_bed = $1
	`
	_, err := r.DB.Exec(query,
		kamar.NomorBed,    // Corresponds to nomor_bed
		kamar.KodeKamar,   // Corresponds to kode_kamar
		kamar.NamaKamar,   // Corresponds to nama_kamar
		kamar.Kelas,       // Corresponds to kelas
		kamar.TarifKamar,  // Corresponds to tarif_kamar
		kamar.StatusKamar, // Corresponds to status_kamar
	)
	return err
}

func (r *kamarRepositoryImpl) Delete(nomorReg string) error {
	query := `
		DELETE FROM kamar WHERE nomor_bed = $1
	`
	_, err := r.DB.Exec(query, nomorReg)
	return err
}

func (r *kamarRepositoryImpl) GetAvailableRooms() ([]entity.Kamar, error) {
	query := `SELECT * FROM kamar WHERE status_kamar = 'available'`
	var results []entity.Kamar
	err := r.DB.Select(&results, query)
	return results, err
}

func (r *kamarRepositoryImpl) UpdateStatusKamar(nomorBed, status string) error {
	query := `UPDATE kamar SET status_kamar = $1 WHERE nomor_bed = $2`
	_, err := r.DB.Exec(query, status, nomorBed)
	return err
}
