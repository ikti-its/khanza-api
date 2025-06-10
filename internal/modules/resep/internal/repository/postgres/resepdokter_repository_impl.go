package postgres

import (
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/repository"
	"github.com/jmoiron/sqlx"
)

type resepDokterRepositoryImpl struct {
	DB *sqlx.DB
}

func NewResepDokterRepository(db *sqlx.DB) repository.ResepDokterRepository {
	return &resepDokterRepositoryImpl{DB: db}
}

func (r *resepDokterRepositoryImpl) Insert(p *entity.ResepDokter) error {
	query := `
		INSERT INTO resep_dokter (
			no_resep, kode_barang, jumlah, aturan_pakai
		) VALUES (
			$1, $2, $3, $4
		)
	`
	_, err := r.DB.Exec(query,
		p.NoResep, p.KodeBarang, p.Jumlah, p.AturanPakai,
	)
	return err
}

func (r *resepDokterRepositoryImpl) FindAll() ([]entity.ResepDokter, error) {
	query := `SELECT * FROM resep_dokter ORDER BY no_resep DESC`
	var list []entity.ResepDokter
	err := r.DB.Select(&list, query)
	return list, err
}

func (r *resepDokterRepositoryImpl) FindByNoResep(noResep string) ([]entity.ResepDokter, error) {
	query := `SELECT * FROM resep_dokter WHERE no_resep = $1 ORDER BY kode_barang ASC`
	var list []entity.ResepDokter
	err := r.DB.Select(&list, query, noResep)
	return list, err
}

func (r *resepDokterRepositoryImpl) Update(p *entity.ResepDokter) error {
	query := `
		UPDATE resep_dokter SET 
			jumlah = $3, aturan_pakai = $4
		WHERE no_resep = $1 AND kode_barang = $2
	`
	_, err := r.DB.Exec(query,
		p.NoResep, p.KodeBarang, p.Jumlah, p.AturanPakai,
	)
	return err
}

func (r *resepDokterRepositoryImpl) Delete(noResep, kodeBarang string) error {
	query := `DELETE FROM resep_dokter WHERE no_resep = $1 AND kode_barang = $2`
	_, err := r.DB.Exec(query, noResep, kodeBarang)
	return err
}
