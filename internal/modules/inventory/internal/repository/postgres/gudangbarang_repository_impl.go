package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/inventory/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/inventory/internal/repository"
	"github.com/jmoiron/sqlx"
)

type gudangbarangRepositoryImpl struct {
	DB *sqlx.DB
}

func NewGudangBarangRepository(db *sqlx.DB) repository.GudangBarangRepository {
	return &gudangbarangRepositoryImpl{db}
}

func (r *gudangbarangRepositoryImpl) Insert(gudangbarang *entity.GudangBarang) error {
	query := "INSERT INTO gudang_barang VALUES ($1, $2, $3, $4, $5)"

	_, err := r.DB.Exec(query, gudangbarang.IdBarangMedis, gudangbarang.IdRuangan, gudangbarang.Stok, gudangbarang.NoBatch, gudangbarang.NoFaktur)

	return err
}

func (r *gudangbarangRepositoryImpl) Find() ([]entity.GudangBarang, error) {
	query := "SELECT * FROM gudang_barang"

	var records []entity.GudangBarang
	err := r.DB.Select(&records, query)

	return records, err
}

func (r *gudangbarangRepositoryImpl) FindById(id uuid.UUID) (entity.GudangBarang, error) {
	query := "SELECT * FROM gudang_barang WHERE id_barang_medis = $1"

	var record entity.GudangBarang
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *gudangbarangRepositoryImpl) Update(gudangbarang *entity.GudangBarang) error {
	query := "UPDATE gudang_barang SET id_ruangan = $2, stok = $3, no_batch = $4, no_faktur = $5 WHERE id_barang_medis = $1"

	_, err := r.DB.Exec(query, gudangbarang.IdBarangMedis, gudangbarang.IdRuangan, gudangbarang.Stok, gudangbarang.NoBatch, gudangbarang.NoFaktur)

	return err
}

func (r *gudangbarangRepositoryImpl) Delete(gudangbarang *entity.GudangBarang) error {
	query := "DELETE FROM gudang_barang WHERE id_barang_medis = $1"

	_, err := r.DB.Exec(query, gudangbarang.IdBarangMedis)

	return err
}
