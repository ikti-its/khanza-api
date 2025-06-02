package postgres

import (
	"log"

	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/model"
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/repository"
	"github.com/jmoiron/sqlx"
)

type resepDokterRacikanDetailRepositoryImpl struct {
	DB *sqlx.DB
}

func NewResepDokterRacikanDetailRepository(db *sqlx.DB) repository.ResepDokterRacikanDetailRepository {
	return &resepDokterRacikanDetailRepositoryImpl{DB: db}
}

func (r *resepDokterRacikanDetailRepositoryImpl) Insert(d *entity.ResepDokterRacikanDetail) error {
	query := `
		INSERT INTO resep_dokter_racikan_detail (
			no_resep, no_racik, kode_brng, p1, p2, kandungan, jml
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		)
	`
	_, err := r.DB.Exec(query,
		d.NoResep, d.NoRacik, d.KodeBrng, d.P1, d.P2, d.Kandungan, d.Jml,
	)
	return err
}

func (r *resepDokterRacikanDetailRepositoryImpl) FindAll() ([]entity.ResepDokterRacikanDetail, error) {
	query := `SELECT * FROM resep_dokter_racikan_detail ORDER BY no_resep DESC, no_racik ASC`
	var list []entity.ResepDokterRacikanDetail
	err := r.DB.Select(&list, query)
	return list, err
}

func (r *resepDokterRacikanDetailRepositoryImpl) FindByNoResepAndNoRacik(noResep, noRacik string) ([]entity.ResepDokterRacikanDetail, error) {
	query := `
		SELECT * FROM resep_dokter_racikan_detail
		WHERE no_resep = $1 AND no_racik = $2
		ORDER BY kode_brng ASC
	`
	var list []entity.ResepDokterRacikanDetail
	err := r.DB.Select(&list, query, noResep, noRacik)
	return list, err
}

func (r *resepDokterRacikanDetailRepositoryImpl) Update(d *entity.ResepDokterRacikanDetail) error {
	query := `
		UPDATE resep_dokter_racikan_detail SET 
			p1 = $4,
			p2 = $5,
			kandungan = $6,
			jml = $7
		WHERE no_resep = $1 AND no_racik = $2 AND kode_brng = $3
	`
	_, err := r.DB.Exec(query,
		d.NoResep, d.NoRacik, d.KodeBrng, d.P1, d.P2, d.Kandungan, d.Jml,
	)
	return err
}

func (r *resepDokterRacikanDetailRepositoryImpl) Delete(noResep, noRacik, kodeBrng string) error {
	query := `
		DELETE FROM resep_dokter_racikan_detail
		WHERE no_resep = $1 AND no_racik = $2 AND kode_brng = $3
	`
	_, err := r.DB.Exec(query, noResep, noRacik, kodeBrng)
	return err
}

func (r *resepDokterRacikanDetailRepositoryImpl) FindByNoResep(noResep string) ([]model.ResepDokterRacikanDetail, error) {
	var results []model.ResepDokterRacikanDetail

	query := `
		SELECT 
			no_resep, 
			no_racik, 
			kode_brng, 
			p1, 
			p2, 
			kandungan, 
			jml 
		FROM resep_dokter_racikan_detail
		WHERE no_resep = $1
	`

	err := r.DB.Select(&results, query, noResep)
	if err != nil {
		// Hanya log dan return kosong jika tidak ditemukan
		log.Printf("❌ Query failed: %v", err)
		return nil, err
	}

	// Log jika hasilnya kosong tapi tidak error
	log.Printf("✅ Query success. Found %d rows", len(results))
	return results, nil
}
