package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/mobile/internal/repository"
	"github.com/jmoiron/sqlx"
	"math"
)

type pegawaiRepositoryImpl struct {
	DB *sqlx.DB
}

func NewPegawaiRepository(db *sqlx.DB) repository.PegawaiRepository {
	return &pegawaiRepositoryImpl{db}
}

func (r *pegawaiRepositoryImpl) Find() ([]entity.Pegawai, error) {
	query := `
		SELECT p.id AS pegawai, p.id_akun AS akun, p.nip, bp.nik, p.nama, p.jenis_kelamin, bp.tempat_lahir, bp.tanggal_lahir, bp.agama, bp.pendidikan, j.nama AS jabatan, d.nama AS departemen, s.nama AS status, p.jenis_pegawai, p.telepon, p.tanggal_masuk
		FROM pegawai p
		JOIN berkas_pegawai bp ON p.id = bp.id_pegawai
		JOIN ref.jabatan j ON p.id_jabatan = j.id
		JOIN ref.departemen d ON p.id_departemen = d.id
		JOIN ref.status_aktif_pegawai s ON p.id_status_aktif = s.id
		WHERE p.deleted_at IS NULL
	`

	var records []entity.Pegawai
	err := r.DB.Select(&records, query)

	return records, err
}

func (r *pegawaiRepositoryImpl) FindPage(page, size int) ([]entity.Pegawai, int, error) {
	query := `
		SELECT p.id AS pegawai, p.id_akun AS akun, p.nip, bp.nik, p.nama, p.jenis_kelamin, bp.tempat_lahir, bp.tanggal_lahir, bp.agama, bp.pendidikan, j.nama AS jabatan, d.nama AS departemen, s.nama AS status, p.jenis_pegawai, p.telepon, p.tanggal_masuk
		FROM pegawai p 
		JOIN berkas_pegawai bp ON p.id = bp.id_pegawai
		JOIN ref.jabatan j ON p.id_jabatan = j.id
		JOIN ref.departemen d ON p.id_departemen = d.id
		JOIN ref.status_aktif_pegawai s ON p.id_status_aktif = s.id
		WHERE p.deleted_at IS NULL
		LIMIT $1 OFFSET $2
	`
	totalQuery := "SELECT COUNT(*) FROM pegawai WHERE deleted_at IS NULL"

	var total int64
	if err := r.DB.Get(&total, totalQuery); err != nil {
		return nil, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	var records []entity.Pegawai
	err := r.DB.Select(&records, query, size, offset)

	return records, totalPage, err
}

func (r *pegawaiRepositoryImpl) FindById(id uuid.UUID) (entity.Pegawai, error) {
	query := `
		SELECT p.id AS pegawai, p.id_akun AS akun, p.nip, bp.nik, p.nama, p.jenis_kelamin, bp.tempat_lahir, bp.tanggal_lahir, bp.agama, bp.pendidikan, j.nama AS jabatan, d.nama AS departemen, s.nama AS status, p.jenis_pegawai, p.telepon, p.tanggal_masuk
		FROM pegawai p 
		JOIN berkas_pegawai bp ON p.id = bp.id_pegawai
		JOIN ref.jabatan j ON p.id_jabatan = j.id
		JOIN ref.departemen d ON p.id_departemen = d.id
		JOIN ref.status_aktif_pegawai s ON p.id_status_aktif = s.id
		WHERE p.id = $1 AND p.deleted_at IS NULL
	`

	var record entity.Pegawai
	err := r.DB.Get(&record, query, id)

	return record, err
}
