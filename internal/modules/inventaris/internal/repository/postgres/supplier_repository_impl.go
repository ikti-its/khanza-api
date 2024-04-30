package postgres

import (
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/inventaris/internal/repository"
	"github.com/jmoiron/sqlx"
)

type supplierRepositoryImpl struct {
	DB *sqlx.DB
}

func NewSupplierRepository(db *sqlx.DB) repository.SupplierRepository {
	return &supplierRepositoryImpl{db}
}

func (r *supplierRepositoryImpl) Find() ([]entity.Supplier, error) {
	query := `
		SELECT id, nama, alamat, no_telp AS telepon, kota, nama_bank AS bank, no_rekening AS rekening
		FROM ref.supplier_barang_medis
	`

	var records []entity.Supplier
	err := r.DB.Select(&records, query)

	return records, err
}
