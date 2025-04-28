package postgres

import (
	"database/sql"

	"github.com/ikti-its/khanza-api/internal/modules/reseppulang/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/reseppulang/internal/repository"
	"github.com/jmoiron/sqlx"
)

type resepPulangRepositoryImpl struct {
	DB *sqlx.DB
}

func NewResepPulangRepository(db *sqlx.DB) repository.ResepPulangRepository {
	return &resepPulangRepositoryImpl{DB: db}
}

func (r *resepPulangRepositoryImpl) Insert(p *entity.ResepPulang) error {
	query := `
		INSERT INTO resep_pulang (
			no_rawat, kode_brng, jml_barang, harga, total,
			dosis, tanggal, jam, kd_bangsal, no_batch, no_faktur
		) VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8, $9, $10, $11
		)
	`
	_, err := r.DB.Exec(query,
		p.NoRawat, p.KodeBrng, p.JmlBarang, p.Harga, p.Total,
		p.Dosis, p.Tanggal, p.Jam, p.KdBangsal, p.NoBatch, p.NoFaktur,
	)
	return err
}

func (r *resepPulangRepositoryImpl) FindAll() ([]entity.ResepPulang, error) {
	query := `SELECT * FROM resep_pulang ORDER BY tanggal DESC, jam DESC`
	var list []entity.ResepPulang
	err := r.DB.Select(&list, query)
	return list, err
}

func (r *resepPulangRepositoryImpl) FindByNoRawat(noRawat string) ([]entity.ResepPulang, error) {
	query := `SELECT * FROM resep_pulang WHERE no_rawat = $1 ORDER BY tanggal DESC, jam DESC`
	var list []entity.ResepPulang
	err := r.DB.Select(&list, query, noRawat)
	return list, err
}

func (r *resepPulangRepositoryImpl) FindByCompositeKey(noRawat, kodeBrng, tanggal, jam string) (*entity.ResepPulang, error) {
	var data entity.ResepPulang
	query := `
		SELECT * FROM resep_pulang 
		WHERE no_rawat = $1 AND kode_brng = $2 AND tanggal = $3 AND jam = $4
	`
	err := r.DB.Get(&data, query, noRawat, kodeBrng, tanggal, jam)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // ❗ Data not found, not a fatal error
		}
		return nil, err // ❗ Fatal database error
	}
	return &data, nil
}

func (r *resepPulangRepositoryImpl) Update(p *entity.ResepPulang) error {
	query := `
		UPDATE resep_pulang SET 
			jml_barang = $5, harga = $6, total = $7,
			dosis = $8, kd_bangsal = $9, no_batch = $10, no_faktur = $11
		WHERE no_rawat = $1 AND kode_brng = $2 AND tanggal = $3 AND jam = $4
	`
	_, err := r.DB.Exec(query,
		p.NoRawat, p.KodeBrng, p.Tanggal, p.Jam,
		p.JmlBarang, p.Harga, p.Total,
		p.Dosis, p.KdBangsal, p.NoBatch, p.NoFaktur,
	)
	return err
}

func (r *resepPulangRepositoryImpl) Delete(noRawat, kodeBrng, tanggal, jam string) error {
	query := `
    DELETE FROM resep_pulang
    WHERE no_rawat = $1
    AND kode_brng = $2
    AND tanggal = $3::DATE
    AND jam = $4::TIME
`

	_, err := r.DB.Exec(query, noRawat, kodeBrng, tanggal, jam)
	return err
}
