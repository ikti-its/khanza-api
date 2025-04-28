package postgres

import (
	"database/sql"

	"github.com/ikti-its/khanza-api/internal/modules/reseppulang/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/reseppulang/internal/repository"
	"github.com/jmoiron/sqlx"
)

type permintaanResepPulangRepositoryImpl struct {
	DB *sqlx.DB
}

func NewPermintaanResepPulangRepository(db *sqlx.DB) repository.PermintaanResepPulangRepository {
	return &permintaanResepPulangRepositoryImpl{DB: db}
}

func (r *permintaanResepPulangRepositoryImpl) Insert(p *entity.PermintaanResepPulang) error {
	query := `
		INSERT INTO permintaan_resep_pulang (
			no_permintaan, tgl_permintaan, jam, no_rawat, kd_dokter,
			status, tgl_validasi, jam_validasi
		) VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8
		)
	`
	_, err := r.DB.Exec(query,
		p.NoPermintaan, p.TglPermintaan, p.Jam,
		p.NoRawat, p.KdDokter, p.Status,
		p.TglValidasi, p.JamValidasi,
	)
	return err
}

func (r *permintaanResepPulangRepositoryImpl) FindAll() ([]entity.PermintaanResepPulang, error) {
	query := `SELECT * FROM permintaan_resep_pulang ORDER BY tgl_permintaan DESC, jam DESC`
	var list []entity.PermintaanResepPulang
	err := r.DB.Select(&list, query)
	return list, err
}

func (r *permintaanResepPulangRepositoryImpl) FindByNoRawat(noRawat string) ([]entity.PermintaanResepPulang, error) {
	query := `SELECT * FROM permintaan_resep_pulang WHERE no_rawat = $1 ORDER BY tgl_permintaan DESC, jam DESC`
	var list []entity.PermintaanResepPulang
	err := r.DB.Select(&list, query, noRawat)
	return list, err
}

func (r *permintaanResepPulangRepositoryImpl) FindByNoPermintaan(noPermintaan string) (*entity.PermintaanResepPulang, error) {
	var data entity.PermintaanResepPulang
	query := `SELECT * FROM permintaan_resep_pulang WHERE no_permintaan = $1`
	err := r.DB.Get(&data, query, noPermintaan)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // ❗ Data tidak ketemu, tapi bukan error fatal
		}
		return nil, err // ❗ Error fatal lain, kayak query syntax error
	}
	return &data, nil
}

func (r *permintaanResepPulangRepositoryImpl) Update(p *entity.PermintaanResepPulang) error {
	query := `
		UPDATE permintaan_resep_pulang SET 
			tgl_permintaan = $2, jam = $3, no_rawat = $4,
			kd_dokter = $5, status = $6, tgl_validasi = $7, jam_validasi = $8
		WHERE no_permintaan = $1
	`
	_, err := r.DB.Exec(query,
		p.NoPermintaan, p.TglPermintaan, p.Jam, p.NoRawat,
		p.KdDokter, p.Status, p.TglValidasi, p.JamValidasi,
	)
	return err
}

func (r *permintaanResepPulangRepositoryImpl) Delete(noPermintaan string) error {
	query := `DELETE FROM permintaan_resep_pulang WHERE no_permintaan = $1`
	_, err := r.DB.Exec(query, noPermintaan)
	return err
}
