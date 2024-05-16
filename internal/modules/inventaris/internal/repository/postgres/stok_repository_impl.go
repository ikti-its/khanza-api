package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/repository"
	"github.com/jmoiron/sqlx"
	"math"
	"time"
)

type stokRepositoryImpl struct {
	DB *sqlx.DB
}

func NewStokRepository(db *sqlx.DB) repository.StokRepository {
	return &stokRepositoryImpl{db}
}

func (r *stokRepositoryImpl) Insert(stok *entity.Stok) error {
	query := `
		INSERT INTO stok_keluar_barang_medis (id, no_keluar, id_pegawai, tanggal_stok_keluar, keterangan, updater)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.DB.Exec(query, stok.Id, stok.Nomor, stok.IdPegawai, stok.Tanggal, stok.Keterangan, stok.Updater)

	return err
}

func (r *stokRepositoryImpl) Find() ([]entity.Stok, error) {
	query := `
		SELECT id, no_keluar, id_pegawai, tanggal_stok_keluar, keterangan
		FROM stok_keluar_barang_medis
		WHERE deleted_at IS NULL
	`

	var records []entity.Stok
	err := r.DB.Select(&records, query)

	return records, err
}

func (r *stokRepositoryImpl) FindPage(page, size int) ([]entity.Stok, int, error) {
	query := `
		SELECT id, no_keluar, id_pegawai, tanggal_stok_keluar, keterangan
		FROM stok_keluar_barang_medis
		WHERE deleted_at IS NULL
		LIMIT $1 OFFSET $2
	`
	totalQuery := "SELECT COUNT(*) FROM stok_keluar_barang_medis WHERE deleted_at IS NULL"

	var total int64
	if err := r.DB.Get(&total, totalQuery); err != nil {
		return nil, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	var records []entity.Stok
	err := r.DB.Select(&records, query, size, offset)

	return records, totalPage, err
}

func (r *stokRepositoryImpl) FindById(id uuid.UUID) (entity.Stok, error) {
	query := `
		SELECT id, no_keluar, id_pegawai, tanggal_stok_keluar, keterangan
		FROM stok_keluar_barang_medis
		WHERE id = $1 AND deleted_at IS NULL
	`

	var record entity.Stok
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *stokRepositoryImpl) Update(stok *entity.Stok) error {
	query := `
		UPDATE stok_keluar_barang_medis
		SET no_keluar = $2, id_pegawai = $3, tanggal_stok_keluar = $4, keterangan = $5, updated_at = $6, updater = $7
		WHERE id = $1 AND deleted_at IS NULL
	`

	_, err := r.DB.Exec(query, stok.Id, stok.Nomor, stok.IdPegawai, stok.Tanggal, stok.Keterangan, time.Now(), stok.Updater)

	return err
}

func (r *stokRepositoryImpl) Delete(stok *entity.Stok) error {
	query := `
		UPDATE stok_keluar_barang_medis
		SET deleted_at = $2, updater = $3
		WHERE id = $1
	`

	_, err := r.DB.Exec(query, stok.Id, time.Now(), stok.Updater)

	return err
}
