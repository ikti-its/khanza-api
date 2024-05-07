package postgres

import (
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/repository"
	"github.com/jmoiron/sqlx"
)

type satuanRepositoryImpl struct {
	DB *sqlx.DB
}

func NewSatuanRepository(db *sqlx.DB) repository.SatuanRepository {
	return &satuanRepositoryImpl{db}
}

func (r *satuanRepositoryImpl) Find() ([]entity.Satuan, error) {
	query := `
		SELECT id, nama
		FROM ref.satuan_barang_medis
	`

	var records []entity.Satuan
	err := r.DB.Select(&records, query)

	return records, err
}
