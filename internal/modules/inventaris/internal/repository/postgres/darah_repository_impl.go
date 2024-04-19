package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/repository"
	"github.com/jmoiron/sqlx"
	"math"
	"time"
)

type darahRepositoryImpl struct {
	DB *sqlx.DB
}

func NewDarahRepository(db *sqlx.DB) repository.DarahRepository {
	return &darahRepositoryImpl{db}
}

func (r *darahRepositoryImpl) Insert(darah *entity.Darah) error {
	query := `
		INSERT INTO darah (id, id_barang_medis, keterangan, kadaluwarsa, updater)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.DB.Exec(query, darah.Id, darah.IdMedis, darah.Keterangan, darah.Kadaluwarsa, darah.Updater)

	return err
}

func (r *darahRepositoryImpl) Find() ([]entity.Darah, error) {
	query := `
		SELECT id, id_barang_medis, keterangan, kadaluwarsa
		FROM darah
		WHERE deleted_at IS NULL
	`

	var records []entity.Darah
	err := r.DB.Select(&records, query)

	return records, err
}

func (r *darahRepositoryImpl) FindPage(page, size int) ([]entity.Darah, int, error) {
	query := `
		SELECT id, id_barang_medis, keterangan, kadaluwarsa
		FROM darah
		WHERE deleted_at IS NULL
		LIMIT $1 OFFSET $2
	`
	totalQuery := "SELECT COUNT(*) FROM darah WHERE deleted_at IS NULL"

	var total int64
	if err := r.DB.Get(&total, totalQuery); err != nil {
		return nil, 0, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(size)))
	offset := (page - 1) * size

	var records []entity.Darah
	err := r.DB.Select(&records, query, size, offset)

	return records, totalPage, err
}

func (r *darahRepositoryImpl) FindById(id uuid.UUID) (entity.Darah, error) {
	query := `
		SELECT id, id_barang_medis, keterangan, kadaluwarsa
		FROM darah
		WHERE id = $1 AND deleted_at IS NULL
	`

	var record entity.Darah
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *darahRepositoryImpl) FindByIdMedis(id uuid.UUID) (entity.Darah, error) {
	query := `
		SELECT id, id_barang_medis, keterangan, kadaluwarsa
		FROM darah
		WHERE id_barang_medis = $1 AND deleted_at IS NULL
	`

	var record entity.Darah
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *darahRepositoryImpl) Update(darah *entity.Darah) error {
	query := `
		UPDATE darah
		SET id_barang_medis = $2, keterangan = $3, kadaluwarsa = $4, updated_at = $5, updater = $6
		WHERE id = $1 AND deleted_at IS NULL
	`

	_, err := r.DB.Exec(query, darah.Id, darah.IdMedis, darah.Keterangan, darah.Kadaluwarsa, time.Now(), darah.Updater)

	return err
}

func (r *darahRepositoryImpl) Delete(darah *entity.Darah) error {
	query := `
		UPDATE darah
		SET deleted_at = $2, updater = $3
		WHERE id = $1
	`

	_, err := r.DB.Exec(query, darah.Id, time.Now(), darah.Updater)

	return err
}
