package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/kehadiran/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/kehadiran/internal/repository"
	"github.com/jmoiron/sqlx"
	"math"
	"time"
)

type kehadiranRepositoryImpl struct {
	DB *sqlx.DB
}

func NewKehadiranRepository(db *sqlx.DB) repository.KehadiranRepository {
	return &kehadiranRepositoryImpl{DB: db}
}

func (r *kehadiranRepositoryImpl) Insert(kehadiran *entity.Kehadiran) error {
	query := `
		INSERT INTO presensi (id, id_pegawai, id_jadwal_pegawai, tanggal, jam_masuk, keterangan)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.DB.Exec(query, kehadiran.Id, kehadiran.IdPegawai, kehadiran.IdJadwalPegawai, kehadiran.Tanggal, kehadiran.JamMasuk, kehadiran.Keterangan)

	return err
}

func (r *kehadiranRepositoryImpl) Find() ([]entity.Kehadiran, error) {
	query := `
		SELECT id, id_pegawai, id_jadwal_pegawai, tanggal, jam_masuk, jam_pulang, keterangan
		FROM presensi
		WHERE deleted_at IS NULL
		ORDER BY tanggal
	`

	var records []entity.Kehadiran
	err := r.DB.Select(&records, query)

	return records, err
}

func (r *kehadiranRepositoryImpl) FindPage(page, size int) ([]entity.Kehadiran, int, error) {
	query := `
		SELECT id, id_pegawai, id_jadwal_pegawai, tanggal, jam_masuk, jam_pulang, keterangan
		FROM presensi
		WHERE deleted_at IS NULL
		ORDER BY tanggal
		LIMIT $1 OFFSET $2
	`
	totalQuery := "SELECT COUNT(*) FROM presensi WHERE deleted_at IS NULL"

	var total int64
	if err := r.DB.Get(&total, totalQuery); err != nil {
		return nil, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	var records []entity.Kehadiran
	err := r.DB.Select(&records, query, size, offset)

	return records, totalPage, err
}

func (r *kehadiranRepositoryImpl) FindByPegawaiId(id uuid.UUID) ([]entity.Kehadiran, error) {
	query := `
		SELECT id, id_pegawai, id_jadwal_pegawai, tanggal, jam_masuk, jam_pulang, keterangan
		FROM presensi
		WHERE id_pegawai = $1 AND deleted_at IS NULL
		ORDER BY tanggal
	`

	var records []entity.Kehadiran
	err := r.DB.Select(&records, query, id)

	return records, err
}

func (r *kehadiranRepositoryImpl) FindByTanggal(tanggal string) ([]entity.Kehadiran, error) {
	query := `
		SELECT id, id_pegawai, id_jadwal_pegawai, tanggal, jam_masuk, jam_pulang, keterangan
		FROM presensi
		WHERE tanggal = $1 AND deleted_at IS NULL
		ORDER BY tanggal
	`

	var records []entity.Kehadiran
	err := r.DB.Select(&records, query, tanggal)

	return records, err
}

func (r *kehadiranRepositoryImpl) FindById(id uuid.UUID) (entity.Kehadiran, error) {
	query := `
		SELECT id, id_pegawai, id_jadwal_pegawai, tanggal, jam_masuk, jam_pulang, keterangan
		FROM presensi
		WHERE id = $1 AND deleted_at IS NULL
	`

	var record entity.Kehadiran
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *kehadiranRepositoryImpl) Update(kehadiran *entity.Kehadiran) error {
	query := `
		UPDATE presensi
		SET id_pegawai = $1, id_jadwal_pegawai = $2, tanggal = $3, jam_masuk = $4, jam_pulang = $5, 
		    keterangan = (
		        CASE WHEN $4 > (
		            SELECT s.jam_masuk
		            FROM ref.shift s
		            JOIN jadwal_pegawai jp ON s.id = jp.id_shift
		            WHERE jp.id = $2
		        ) THEN 'Terlambat' ELSE 'Hadir' END
		    ), 
		    updated_at = $6, updater = $7
		WHERE id = $8 AND deleted_at IS NULL
	`

	_, err := r.DB.Exec(query, kehadiran.IdPegawai, kehadiran.IdJadwalPegawai, kehadiran.Tanggal, kehadiran.JamMasuk, kehadiran.JamPulang, time.Now(), kehadiran.Updater, kehadiran.Id)

	return err
}
