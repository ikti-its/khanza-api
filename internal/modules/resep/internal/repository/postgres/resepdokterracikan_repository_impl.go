package postgres

import (
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/repository"
	"github.com/jmoiron/sqlx"
)

type resepDokterRacikanRepositoryImpl struct {
	DB *sqlx.DB
}

func NewResepDokterRacikanRepository(db *sqlx.DB) repository.ResepDokterRacikanRepository {
	return &resepDokterRacikanRepositoryImpl{DB: db}
}

func (r *resepDokterRacikanRepositoryImpl) Insert(p *entity.ResepDokterRacikan) error {
	query := `
		INSERT INTO resep_dokter_racikan (
			no_resep, no_racik, nama_racik, kd_racik, jml_dr, aturan_pakai, keterangan
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		)
	`
	_, err := r.DB.Exec(query,
		p.NoResep, p.NoRacik, p.NamaRacik, p.KdRacik, p.JmlDr, p.AturanPakai, p.Keterangan,
	)
	return err
}

func (r *resepDokterRacikanRepositoryImpl) FindAll() ([]entity.ResepDokterRacikan, error) {
	query := `SELECT * FROM resep_dokter_racikan ORDER BY no_resep DESC, no_racik ASC`
	var list []entity.ResepDokterRacikan
	err := r.DB.Select(&list, query)
	return list, err
}

func (r *resepDokterRacikanRepositoryImpl) FindByNoResep(noResep string) ([]entity.ResepDokterRacikan, error) {
	query := `SELECT * FROM resep_dokter_racikan WHERE no_resep = $1 ORDER BY no_racik ASC`
	var list []entity.ResepDokterRacikan
	err := r.DB.Select(&list, query, noResep)
	return list, err
}

func (r *resepDokterRacikanRepositoryImpl) Update(p *entity.ResepDokterRacikan) error {
	query := `
		UPDATE resep_dokter_racikan SET 
			nama_racik = $3,
			kd_racik = $4,
			jml_dr = $5,
			aturan_pakai = $6,
			keterangan = $7
		WHERE no_resep = $1 AND no_racik = $2
	`
	_, err := r.DB.Exec(query,
		p.NoResep, p.NoRacik, p.NamaRacik, p.KdRacik, p.JmlDr, p.AturanPakai, p.Keterangan,
	)
	return err
}

func (r *resepDokterRacikanRepositoryImpl) Delete(noResep, noRacik string) error {
	query := `DELETE FROM resep_dokter_racikan WHERE no_resep = $1 AND no_racik = $2`
	_, err := r.DB.Exec(query, noResep, noRacik)
	return err
}
