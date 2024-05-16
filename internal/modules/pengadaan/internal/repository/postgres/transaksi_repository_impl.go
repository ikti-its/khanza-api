package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/pengadaan/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/pengadaan/internal/repository"
	"github.com/jmoiron/sqlx"
	"math"
	"time"
)

type transaksiRepositoryImpl struct {
	DB *sqlx.DB
}

func NewTransaksiRepository(db *sqlx.DB) repository.TransaksiRepository {
	return &transaksiRepositoryImpl{db}
}

func (r *transaksiRepositoryImpl) Insert(transaksi *entity.Transaksi) error {
	query := `
		INSERT INTO transaksi_keluar_barang_medis (id, id_stok_keluar, id_barang_medis, no_batch, no_faktur, jumlah_keluar, updater)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.DB.Exec(query, transaksi.Id, transaksi.IdStok, transaksi.IdMedis, transaksi.Batch, transaksi.Faktur, transaksi.Jumlah, transaksi.Updater)

	return err
}

func (r *transaksiRepositoryImpl) Find() ([]entity.Transaksi, error) {
	query := `
		SELECT id, id_stok_keluar, id_barang_medis, no_batch, no_faktur, jumlah_keluar
		FROM transaksi_keluar_barang_medis
		WHERE deleted_at IS NULL
	`

	var records []entity.Transaksi
	err := r.DB.Select(&records, query)

	return records, err
}

func (r *transaksiRepositoryImpl) FindPage(page, size int) ([]entity.Transaksi, int, error) {
	query := `
		SELECT id, id_stok_keluar, id_barang_medis, no_batch, no_faktur, jumlah_keluar
		FROM transaksi_keluar_barang_medis
		WHERE deleted_at IS NULL
		LIMIT $1 OFFSET $2
	`
	totalQuery := "SELECT COUNT(*) FROM transaksi_keluar_barang_medis WHERE deleted_at IS NULL"

	var total int64
	if err := r.DB.Get(&total, totalQuery); err != nil {
		return nil, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	var records []entity.Transaksi
	err := r.DB.Select(&records, query, size, offset)

	return records, totalPage, err
}

func (r *transaksiRepositoryImpl) FindById(id uuid.UUID) (entity.Transaksi, error) {
	query := `
		SELECT id, id_stok_keluar, id_barang_medis, no_batch, no_faktur, jumlah_keluar
		FROM transaksi_keluar_barang_medis
		WHERE id = $1 AND deleted_at IS NULL
	`

	var record entity.Transaksi
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *transaksiRepositoryImpl) Update(transaksi *entity.Transaksi) error {
	query := `
		UPDATE transaksi_keluar_barang_medis
		SET id_stok_keluar = $2, id_barang_medis = $3, no_batch = $4, no_faktur = $5, jumlah_keluar = $6, updated_at = $7, updater = $8
		WHERE id = $1 AND deleted_at IS NULL
	`

	_, err := r.DB.Exec(query, transaksi.Id, transaksi.IdStok, transaksi.IdMedis, transaksi.Batch, transaksi.Faktur, transaksi.Jumlah, time.Now(), transaksi.Updater)

	return err
}

func (r *transaksiRepositoryImpl) Delete(transaksi *entity.Transaksi) error {
	query := `
		UPDATE transaksi_keluar_barang_medis
		SET deleted_at = $2, updater = $3
		WHERE id = $1
	`

	_, err := r.DB.Exec(query, transaksi.Id, time.Now(), transaksi.Updater)

	return err
}
