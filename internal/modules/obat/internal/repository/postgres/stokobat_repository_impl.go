package postgres

import (
	"log"

	"github.com/ikti-its/khanza-api/internal/modules/obat/internal/entity"
	gudangEntity "github.com/ikti-its/khanza-api/internal/modules/obat/internal/entity"
	"github.com/jmoiron/sqlx"
)

type gudangBarangRepository struct {
	db *sqlx.DB
}

type GudangBarangRepository interface {
	Insert(barang *entity.GudangBarang) error
	FindAll() ([]entity.GudangBarang, error)
	FindByID(id string) (*entity.GudangBarang, error)
	Update(barang *entity.GudangBarang) error
	Delete(id string) error
	FindByIDBarangMedis(idBarangMedis string) ([]entity.GudangBarang, error)
}

func NewGudangBarangRepository(db *sqlx.DB) GudangBarangRepository {
	return &gudangBarangRepository{db: db}
}

func (r *gudangBarangRepository) Insert(barang *entity.GudangBarang) error {
	query := `
		INSERT INTO sik.gudang_barang (id, id_barang_medis, id_ruangan, stok, no_batch, no_faktur)
		VALUES (:id, :id_barang_medis, :id_ruangan, :stok, :no_batch, :no_faktur)
	`
	_, err := r.db.NamedExec(query, barang)
	return err
}

func (r *gudangBarangRepository) FindAll() ([]entity.GudangBarang, error) {
	var result []entity.GudangBarang
	query := `
		SELECT id, id_barang_medis, id_ruangan, stok, no_batch, no_faktur
		FROM sik.gudang_barang
		ORDER BY id ASC
	`
	err := r.db.Select(&result, query)
	return result, err
}

func (r *gudangBarangRepository) FindByID(id string) (*entity.GudangBarang, error) {
	var result entity.GudangBarang
	query := `
		SELECT id, id_barang_medis, id_ruangan, stok, no_batch, no_faktur
		FROM sik.gudang_barang
		WHERE id_barang_medis = $1
	`
	err := r.db.Get(&result, query, id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *gudangBarangRepository) Update(barang *entity.GudangBarang) error {
	query := `
		UPDATE sik.gudang_barang
		SET id_barang_medis = :id_barang_medis,
			id_ruangan = :id_ruangan,
			stok = :stok,
			no_batch = :no_batch,
			no_faktur = :no_faktur
		WHERE id = :id
	`
	_, err := r.db.NamedExec(query, barang)
	return err
}

func (r *gudangBarangRepository) Delete(id string) error {
	query := `DELETE FROM sik.gudang_barang WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *gudangBarangRepository) FindByIDBarangMedis(idBarangMedis string) ([]gudangEntity.GudangBarang, error) {
	var result []gudangEntity.GudangBarang
	log.Printf("→ TYPE CHECK: IDBarangMedis is %T\n", result[0].IDBarangMedis)
	log.Printf(">> Using GudangBarang from: %T\n", result)
	query := `
		SELECT id, id_barang_medis, id_ruangan, stok, no_batch, no_faktur
		FROM sik.gudang_barang
		WHERE id_barang_medis = $1
	`

	err := r.db.Select(&result, query, idBarangMedis)
	return result, err
}
