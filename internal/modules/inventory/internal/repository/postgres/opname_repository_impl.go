package postgres

import (
	"github.com/google/uuid"
	"github.com/ikti-its/khanza-api/internal/modules/inventory/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/inventory/internal/repository"
	"github.com/jmoiron/sqlx"
)

type opnameRepositoryImpl struct {
	DB *sqlx.DB
}

func NewOpnameRepository(db *sqlx.DB) repository.OpnameRepository {
	return &opnameRepositoryImpl{db}
}

func (r *opnameRepositoryImpl) Insert(opname *entity.Opname) error {
	query := "INSERT INTO opname VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	_, err := r.DB.Exec(query, opname.IdBarangMedis, opname.IdRuangan, opname.HBeli, opname.Tanggal, opname.Real, opname.Stok, opname.Keterangan, opname.NoBatch, opname.NoFaktur)

	return err
}

func (r *opnameRepositoryImpl) Find() ([]entity.Opname, error) {
	query := "SELECT * FROM opname"

	var records []entity.Opname
	err := r.DB.Select(&records, query)

	return records, err
}

func (r *opnameRepositoryImpl) FindById(id uuid.UUID) (entity.Opname, error) {
	query := "SELECT * FROM opname WHERE id_barang_medis = $1"

	var record entity.Opname
	err := r.DB.Get(&record, query, id)

	return record, err
}

func (r *opnameRepositoryImpl) Update(opname *entity.Opname) error {
	query := "UPDATE opname SET id_ruangan = $2, h_beli = $3, tanggal = $4, real = $5, stok = $6, keterangan = $7, no_batch = $8, no_faktur = $9 WHERE id_barang_medis = $1"

	_, err := r.DB.Exec(query, opname.IdBarangMedis, opname.IdRuangan, opname.HBeli, opname.Tanggal, opname.Real, opname.Stok, opname.Keterangan, opname.NoBatch, opname.NoFaktur)

	return err
}

func (r *opnameRepositoryImpl) Delete(opname *entity.Opname) error {
	query := "DELETE FROM opname WHERE id_barang_medis = $1"

	_, err := r.DB.Exec(query, opname.IdBarangMedis)

	return err
}
