package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ikti-its/khanza-api/internal/modules/pasienmeninggal/internal/entity"
)

type Repository interface {
	FindAll() ([]entity.PasienMeninggal, error)
	FindById(noRM string) (entity.PasienMeninggal, error)
	Insert(data *entity.PasienMeninggal) error
	Update(data *entity.PasienMeninggal) error
	Delete(noRM string) error
}

type RepositoryImpl struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{DB: db}
}

func (r *RepositoryImpl) FindAll() ([]entity.PasienMeninggal, error) {
	query := `SELECT * FROM sik.pasien_meninggal ORDER BY no_rkm_medis ASC`

	var records []entity.PasienMeninggal
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *RepositoryImpl) FindById(noRM string) (entity.PasienMeninggal, error) {
	query := `SELECT * FROM sik.pasien_meninggal WHERE no_rkm_medis = $1`

	var record entity.PasienMeninggal
	err := r.DB.Get(&record, query, noRM)
	return record, err
}

func (r *RepositoryImpl) Insert(data *entity.PasienMeninggal) error {
	query := `
		INSERT INTO sik.pasien_meninggal (
			no_rkm_medis, nm_pasien, jk, tgl_lahir, umur,
			gol_darah, stts_nikah, agama, tanggal, jam,
			icdx, icdx_antara1, icdx_antara2, icdx_langsung,
			keterangan, nama_dokter, kode_dokter
		) VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8, $9, $10,
			$11, $12, $13, $14,
			$15, $16, $17
		)
	`
	_, err := r.DB.Exec(query,
		data.NoRkmMedis, data.NmPasien, data.JK, data.TglLahir, data.Umur,
		data.GolDarah, data.SttsNikah, data.Agama, data.Tanggal, data.Jam,
		data.ICDX, data.ICDXAntara1, data.ICDXAntara2, data.ICDXLangsung,
		data.Keterangan, data.NamaDokter, data.KodeDokter,
	)
	return err
}

func (r *RepositoryImpl) Update(data *entity.PasienMeninggal) error {
	query := `
		UPDATE pasien_meninggal SET
			nm_pasien = $2,
			jk = $3,
			tgl_lahir = $4,
			umur = $5,
			gol_darah = $6,
			stts_nikah = $7,
			agama = $8,
			tanggal = $9,
			jam = $10,
			icdx = $11,
			icdx_antara1 = $12,
			icdx_antara2 = $13,
			icdx_langsung = $14,
			keterangan = $15,
			nama_dokter = $16,
			kode_dokter = $17
		WHERE no_rkm_medis = $1
	`
	_, err := r.DB.Exec(query,
		data.NoRkmMedis, data.NmPasien, data.JK, data.TglLahir, data.Umur,
		data.GolDarah, data.SttsNikah, data.Agama, data.Tanggal, data.Jam,
		data.ICDX, data.ICDXAntara1, data.ICDXAntara2, data.ICDXLangsung,
		data.Keterangan, data.NamaDokter, data.KodeDokter,
		
	)
	return err
}

func (r *RepositoryImpl) Delete(noRM string) error {
	query := `DELETE FROM sik.pasien_meninggal WHERE no_rkm_medis = $1`
	_, err := r.DB.Exec(query, noRM)
	return err
}
