package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/repository"
	"github.com/jmoiron/sqlx"
	"math"
	"time"
)

type bhpRepositoryImpl struct {
	DB *sqlx.DB
}

func NewBhpRepository(db *sqlx.DB) repository.BhpRepository {
	return &bhpRepositoryImpl{db}
}

func (r *bhpRepositoryImpl) Insert(bhp *entity.Bhp) error {
	query := `
		INSERT INTO bahan_habis_pakai (id, id_barang_medis, satuan, jumlah, kadaluwarsa, updater)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.DB.Exec(query, bhp.Id, bhp.IdMedis, bhp.Satuan, bhp.Jumlah, bhp.Kadaluwarsa, bhp.Updater)

	return err
}

func (r *bhpRepositoryImpl) Find() ([]entity.Bhp, error) {
	query := `
		SELECT id, id_barang_medis, satuan, jumlah, kadaluwarsa
		FROM bahan_habis_pakai
		WHERE deleted_at IS NULL
	`

	var records []entity.Bhp
	err := r.DB.Select(&records, query)

	return records, err
}

func (r *bhpRepositoryImpl) FindPage(page, size int) ([]entity.Bhp, int, error) {
	query := `
		SELECT id, id_barang_medis, satuan, jumlah, kadaluwarsa
		FROM bahan_habis_pakai
		WHERE deleted_at IS NULL
		LIMIT $1 OFFSET $2
	`
	totalQuery := "SELECT COUNT(*) FROM bahan_habis_pakai WHERE deleted_at IS NULL"

	var total int64
	if err := r.DB.Get(&total, totalQuery); err != nil {
		return nil, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	var records []entity.Bhp
	err := r.DB.Select(&records, query, size, offset)

	return records, totalPage, err
}

func (r *bhpRepositoryImpl) FindById(id uuid.UUID) (entity.Bhp, error) {
	query := `
		SELECT id, id_barang_medis, satuan, jumlah, kadaluwarsa
		FROM bahan_habis_pakai
		WHERE id = $1 AND deleted_at IS NULL
	`

	var record entity.Bhp
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *bhpRepositoryImpl) Update(bhp *entity.Bhp) error {
	query := `
		UPDATE bahan_habis_pakai
		SET id_barang_medis = $2, satuan = $3, jumlah = $4, kadaluwarsa = $5, updated_at = $6, updater = $7
		WHERE id = $1 AND deleted_at IS NULL
	`

	_, err := r.DB.Exec(query, bhp.Id, bhp.IdMedis, bhp.Satuan, bhp.Jumlah, bhp.Kadaluwarsa, time.Now(), bhp.Updater)

	return err
}

func (r *bhpRepositoryImpl) Delete(bhp *entity.Bhp) error {
	query := `
		UPDATE bahan_habis_pakai
		SET deleted_at = $2, updater = $3
		WHERE id = $1
	`

	_, err := r.DB.Exec(query, bhp.Id, time.Now(), bhp.Updater)

	return err
}
