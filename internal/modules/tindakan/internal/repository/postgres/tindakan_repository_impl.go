package postgres

import (
	"log"

	"github.com/ikti-its/khanza-api/internal/modules/tindakan/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/tindakan/internal/repository"
	"github.com/jmoiron/sqlx"
)

type TindakanRepository interface {
	Insert(t *entity.Tindakan) error
	FindAll() ([]entity.Tindakan, error)
	FindByNomorRawat(nomorRawat string) ([]entity.Tindakan, error)
	Update(t *entity.Tindakan) error
	Delete(nomorRawat string, jamRawat string) error
	CheckDokterExists(kodeDokter string) (bool, error)
	GetAllJenisTindakan() ([]entity.JenisTindakan, error)
}

type tindakanRepositoryImpl struct {
	DB *sqlx.DB
}

func NewTindakanRepository(db *sqlx.DB) repository.TindakanRepository {
	return &tindakanRepositoryImpl{DB: db}
}

func (r *tindakanRepositoryImpl) Insert(t *entity.Tindakan) error {
	query := `
		INSERT INTO tindakan (
			nomor_rawat, nomor_rm, nama_pasien, tindakan, kode_dokter, nama_dokter,
			nip, nama_petugas, tanggal_rawat, jam_rawat, biaya
		) VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9, $10, $11
		)
	`
	_, err := r.DB.Exec(query,
		t.NomorRawat, t.NomorRM, t.NamaPasien, t.Tindakan, t.KodeDokter, t.NamaDokter,
		t.NIP, t.NamaPetugas, t.TanggalRawat, t.JamRawat, t.Biaya,
	)
	return err
}

func (r *tindakanRepositoryImpl) FindAll() ([]entity.Tindakan, error) {
	query := `SELECT * FROM tindakan ORDER BY tanggal_rawat DESC`
	var records []entity.Tindakan
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *tindakanRepositoryImpl) FindByNomorRawat(nomorRawat string) ([]entity.Tindakan, error) {
	query := `SELECT * FROM tindakan WHERE nomor_rawat = $1 ORDER BY tanggal_rawat DESC`
	var list []entity.Tindakan
	err := r.DB.Select(&list, query, nomorRawat)
	return list, err
}

func (r *tindakanRepositoryImpl) Update(t *entity.Tindakan) error {
	query := `
		UPDATE tindakan SET 
			nomor_rm = $2, nama_pasien = $3, tindakan = $4,
			kode_dokter = $5, nama_dokter = $6, nip = $7, nama_petugas = $8,
			tanggal_rawat = $9, jam_rawat = $10, biaya = $11
		WHERE nomor_rawat = $1
	`
	_, err := r.DB.Exec(query,
		t.NomorRawat, t.NomorRM, t.NamaPasien, t.Tindakan, t.KodeDokter, t.NamaDokter,
		t.NIP, t.NamaPetugas, t.TanggalRawat, t.JamRawat, t.Biaya,
	)
	return err
}

func (r *tindakanRepositoryImpl) Delete(nomorRawat, jamRawat string) error {
	query := `DELETE FROM tindakan WHERE nomor_rawat = $1 AND jam_rawat = $2`
	_, err := r.DB.Exec(query, nomorRawat, jamRawat)
	return err
}

func (r *tindakanRepositoryImpl) CheckDokterExists(kodeDokter string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM dokter WHERE kode_dokter = $1)`
	err := r.DB.QueryRow(query, kodeDokter).Scan(&exists)
	return exists, err
}

func (r *tindakanRepositoryImpl) GetAllJenisTindakan() ([]entity.JenisTindakan, error) {
	var result []entity.JenisTindakan
	query := `
	SELECT 
		kode,
		nama_tindakan,
		kode_kategori,
		material,
		bhp,
		tarif_tindakan_dokter,
		tarif_tindakan_perawat,
		kso,
		manajemen,
		total_bayar_dokter,
		total_bayar_perawat,
		(material + bhp + kso + manajemen + total_bayar_dokter + total_bayar_perawat) AS tarif,
		total_bayar_dokter_perawat,
		kode_pj,
		kode_bangsal,
		status,
		kelas
	FROM jenis_tindakan
    ORDER BY nama_tindakan ASC`

	err := r.DB.Select(&result, query)
	log.Printf("[QUERY] %s", query)
	log.Printf("[RESULT] fetched %d rows", len(result))

	if err != nil {
		log.Printf("[ERROR] Select failed: %v", err)
	}
	return result, err
}
