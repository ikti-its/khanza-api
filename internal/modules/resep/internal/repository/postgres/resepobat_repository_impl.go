package postgres

import (
	"context"

	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/entity"
	"github.com/ikti-its/khanza-api/internal/modules/resep/internal/repository"
	"github.com/jmoiron/sqlx"
)

type resepObatRepositoryImpl struct {
	DB *sqlx.DB
}

func NewResepObatRepository(db *sqlx.DB) repository.ResepObatRepository {
	return &resepObatRepositoryImpl{DB: db}
}

func (r *resepObatRepositoryImpl) Insert(p *entity.ResepObat) error {
	query := `
		INSERT INTO resep_obat (
			no_resep, tgl_perawatan, jam, no_rawat, kd_dokter,
			tgl_peresepan, jam_peresepan, status, tgl_penyerahan, jam_penyerahan, validasi
		) VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8, $9, $10, $11
		)
	`
	_, err := r.DB.Exec(query,
		p.NoResep, p.TglPerawatan, p.Jam, p.NoRawat, p.KdDokter,
		p.TglPeresepan, p.JamPeresepan, p.Status, p.TglPenyerahan, p.JamPenyerahan, p.Validasi,
	)
	return err
}

func (r *resepObatRepositoryImpl) FindAll() ([]entity.ResepObat, error) {
	query := `SELECT * FROM resep_obat ORDER BY no_resep DESC `
	var list []entity.ResepObat
	err := r.DB.Select(&list, query)
	return list, err
}

func (r *resepObatRepositoryImpl) FindByNoResep(noResep string) (*entity.ResepObat, error) {
	query := `SELECT * FROM resep_obat WHERE no_resep = $1`
	var resep entity.ResepObat
	err := r.DB.Get(&resep, query, noResep)
	if err != nil {
		return nil, err
	}
	return &resep, nil
}

func (r *resepObatRepositoryImpl) Update(p *entity.ResepObat) error {
	query := `
		UPDATE resep_obat SET 
			tgl_perawatan = $2,
			jam = $3,
			no_rawat = $4,
			kd_dokter = $5,
			tgl_peresepan = $6,
			jam_peresepan = $7,
			status = $8,
			tgl_penyerahan = $9,
			jam_penyerahan = $10,
			validasi = $11
		WHERE no_resep = $1
	`
	_, err := r.DB.Exec(query,
		p.NoResep, p.TglPerawatan, p.Jam, p.NoRawat, p.KdDokter,
		p.TglPeresepan, p.JamPeresepan, p.Status, p.TglPenyerahan, p.JamPenyerahan,
	)
	return err
}

func (r *resepObatRepositoryImpl) Delete(noResep string) error {
	query := `DELETE FROM resep_obat WHERE no_resep = $1`
	_, err := r.DB.Exec(query, noResep)
	return err
}

func (r *resepObatRepositoryImpl) GetByNomorRawat(nomorRawat string) ([]entity.ResepObat, error) {
	var resepObats []entity.ResepObat
	query := `SELECT * FROM sik.resep_obat WHERE no_rawat = $1`
	err := r.DB.Select(&resepObats, query, nomorRawat)
	return resepObats, err
}

func (r *resepObatRepositoryImpl) UpdateValidasi(ctx context.Context, noResep string, validasi bool) error {
	query := `UPDATE resep_obat SET validasi = $1 WHERE no_resep = $2`
	_, err := r.DB.ExecContext(ctx, query, validasi, noResep)
	return err
}
