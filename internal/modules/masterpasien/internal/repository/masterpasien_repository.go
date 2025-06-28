package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ikti-its/khanza-api/internal/modules/masterpasien/internal/entity"
)

type Repository interface {
	FindAll() ([]entity.MasterPasien, error)
	FindById(id string) (entity.MasterPasien, error)
	Insert(data *entity.MasterPasien) error
	Update(data *entity.MasterPasien) error
	Delete(id string) error
}

type RepositoryImpl struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{DB: db}
}

func (r *RepositoryImpl) FindAll() ([]entity.MasterPasien, error) {
	query := `
SELECT 
  no_rkm_medis,
  nm_pasien,
  no_ktp,
  jk,
  tmp_lahir,
  tgl_lahir,
  nm_ibu,
  alamat,
  gol_darah,
  pekerjaan,
  stts_nikah,
  agama,
  tgl_daftar,
  no_tlp,
  umur,
  pnd,
  namakeluarga,
  kd_kel,
  kd_pj,
  no_peserta,
  pekerjaanpj,
  alamatpj,
  suku_bangsa,
  bahasa_pasien,
  perusahaan_pasien,
  nip,
  email,
  cacat_fisik
FROM pasien
ORDER BY no_rkm_medis DESC
`

	var result []entity.MasterPasien
	err := r.DB.Select(&result, query)
	return result, err
}

func (r *RepositoryImpl) FindById(id string) (entity.MasterPasien, error) {
	query := `SELECT * FROM pasien WHERE no_rkm_medis = $1`
	var pasien entity.MasterPasien
	err := r.DB.Get(&pasien, query, id)
	return pasien, err
}

func (r *RepositoryImpl) Insert(data *entity.MasterPasien) error {
	query := `
		INSERT INTO pasien (
			no_rkm_medis, nm_pasien, no_ktp, jk, tmp_lahir, tgl_lahir,
			nm_ibu, alamat, gol_darah, pekerjaan, stts_nikah, agama,
			tgl_daftar, no_tlp, umur, pnd, keluarga, namakeluarga,
			kd_pj, no_peserta, pekerjaanpj, alamatpj, suku_bangsa, bahasa_pasien,
			perusahaan_pasien, nip, email, cacat_fisik
		) VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9, $10, $11, $12,
			$13, $14, $15, $16, $17, $18,
			$19, $20, $21, $22, $23, $24,
			$25, $26, $27, $28
		)
	`
	_, err := r.DB.Exec(query,
		data.NoRM, data.NamaPasien, data.NoKTP, data.JenisKelamin, data.TempatLahir, data.TanggalLahir,
		data.NamaIbu, data.Alamat, data.GolonganDarah, data.Pekerjaan, data.StatusPernikahan, data.Agama,
		data.TanggalDaftar, data.NoTelepon, data.Umur, data.Pendidikan, data.PenanggungJawab, data.NamaPJ,
		data.Asuransi, data.NoPeserta, data.PekerjaanPJ, data.AlamatPJ, data.Suku, data.Bahasa,
		data.Instansi, data.NIPNRP, data.Email, data.CacatFisik,
	)
	return err
}

func (r *RepositoryImpl) Update(data *entity.MasterPasien) error {
	query := `
		UPDATE pasien SET 
			nm_pasien = $2, no_ktp = $3, jk = $4, tmp_lahir = $5, tgl_lahir = $6,
			nm_ibu = $7, alamat = $8, gol_darah = $9, pekerjaan = $10, stts_nikah = $11,
			agama = $12, tgl_daftar = $13, no_tlp = $14, umur = $15, pnd = $16,
			keluarga = $17, namakeluarga = $18, kd_pj = $19, no_peserta = $20,
			pekerjaanpj = $21, alamatpj = $22, suku_bangsa = $23, bahasa_pasien = $24,
			perusahaan_pasien = $25, nip = $26, email = $27, cacat_fisik = $28
		WHERE no_rkm_medis = $1
	`
	_, err := r.DB.Exec(query,
		data.NoRM, data.NamaPasien, data.NoKTP, data.JenisKelamin, data.TempatLahir, data.TanggalLahir,
		data.NamaIbu, data.Alamat, data.GolonganDarah, data.Pekerjaan, data.StatusPernikahan, data.Agama,
		data.TanggalDaftar, data.NoTelepon, data.Umur, data.Pendidikan, data.PenanggungJawab, data.NamaPJ,
		data.Asuransi, data.NoPeserta, data.PekerjaanPJ, data.AlamatPJ, data.Suku, data.Bahasa,
		data.Instansi, data.NIPNRP, data.Email, data.CacatFisik,
	)
	return err
}

func (r *RepositoryImpl) Delete(id string) error {
	query := `DELETE FROM pasien WHERE no_rkm_medis = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
