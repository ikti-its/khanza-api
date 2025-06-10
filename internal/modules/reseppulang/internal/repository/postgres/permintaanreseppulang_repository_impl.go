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

func (r *permintaanResepPulangRepositoryImpl) InsertMany(perms []*entity.PermintaanResepPulang) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO permintaan_resep_pulang (
			no_permintaan, tgl_permintaan, jam, no_rawat, kd_dokter,
			status, tgl_validasi, jam_validasi,
			kode_brng, jumlah, aturan_pakai
		) VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8,
			$9, $10, $11
		)
	`

	for _, p := range perms {
		_, err := tx.Exec(query,
			p.NoPermintaan, p.TglPermintaan, p.Jam, p.NoRawat, p.KdDokter,
			p.Status, p.TglValidasi, p.JamValidasi,
			p.KodeBrng, p.Jumlah, p.AturanPakai,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
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
	query := `SELECT * FROM permintaan_resep_pulang WHERE no_permintaan = $1 LIMIT 1`
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
			kd_dokter = $5, status = $6, tgl_validasi = $7, jam_validasi = $8,
			kode_brng = $9, jumlah = $10, aturan_pakai = $11
		WHERE no_permintaan = $1 AND kode_brng = $9
	`
	_, err := r.DB.Exec(query,
		p.NoPermintaan, p.TglPermintaan, p.Jam, p.NoRawat,
		p.KdDokter, p.Status, p.TglValidasi, p.JamValidasi,
		p.KodeBrng, p.Jumlah, p.AturanPakai,
	)
	return err
}

func (r *permintaanResepPulangRepositoryImpl) Delete(noPermintaan string) error {
	query := `DELETE FROM permintaan_resep_pulang WHERE no_permintaan = $1`
	_, err := r.DB.Exec(query, noPermintaan)
	return err
}

func (r *permintaanResepPulangRepositoryImpl) GetByNoPermintaan(noPermintaan string) ([]entity.PermintaanResepPulang, error) {
	var results []entity.PermintaanResepPulang
	query := `SELECT * FROM permintaan_resep_pulang WHERE no_permintaan = $1`

	err := r.DB.Select(&results, query, noPermintaan)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *permintaanResepPulangRepositoryImpl) GetByNoPermintaanWithHarga(noPermintaan string) ([]entity.ResepPulangObat, error) {
	var results []entity.ResepPulangObat

	query := `
		SELECT 
			prp.no_permintaan,
			prp.kode_brng,
			prp.jumlah,
			prp.aturan_pakai,
			db.nama_brng AS nama_obat,
			db.dasar AS harga_dasar,
			db.kelas1,
			db.kelas2,
			db.kelas3,
			db.utama,
			db.vip,
			db.vvip
		FROM permintaan_resep_pulang prp
		LEFT JOIN databarang db ON prp.kode_brng = db.kode_brng
		WHERE prp.no_permintaan = $1;
	`

	err := r.DB.Select(&results, query, noPermintaan)
	if err != nil {
		return nil, err
	}
	return results, nil
}
