package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ikti-its/khanza-api/internal/modules/masterpasien/internal/entity"
)

type Repository interface {
	FindAll() ([]entity.MasterPasien, error)
	FindById(noRkmMedis string) (entity.MasterPasien, error)
	Insert(pasien *entity.MasterPasien) error
	Update(pasien *entity.MasterPasien) error
	Delete(noRkmMedis string) error
}

type RepositoryImpl struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{DB: db}
}

func (r *RepositoryImpl) FindAll() ([]entity.MasterPasien, error) {
	query := `SELECT * FROM pasien ORDER BY no_rkm_medis DESC`
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
		perusahaan_pasien, nip, email, cacat_fisik,
		kd_kel, kd_kec, kd_kab, kelurahanpj, kecamatanpj, kabupatenpj, kd_prop, propinsipj
	) VALUES (
		$1, $2, $3, $4, $5, $6,
		$7, $8, $9, $10, $11, $12,
		$13, $14, $15, $16, $17, $18,
		$19, $20, $21, $22, $23, $24,
		$25, $26, $27, $28,
		$29, $30, $31, $32, $33, $34, $35, $36
	)`
_, err := r.DB.Exec(query,
	data.No_rkm_medis,
	data.Nm_pasien,
	data.No_ktp,
	data.Jk,
	data.Tmp_lahir,
	data.Tgl_lahir,
	data.Nm_ibu,
	data.Alamat,
	data.Gol_darah,
	data.Pekerjaan,
	data.Stts_nikah,
	data.Agama,
	data.Tgl_daftar,
	data.No_tlp,
	data.Umur,
	data.Pnd,
	data.Keluarga,
	data.Namakeluarga,
	data.Kd_pj,
	data.No_peserta,
	data.Pekerjaanpj,
	data.Alamatpj,
	data.Suku_bangsa,
	data.Bahasa_pasien,
	data.Perusahaan_pasien,
	data.Nip,
	data.Email,
	data.Cacat_fisik,
	data.Kd_kel,
	data.Kd_kec,
	data.Kd_kab,
	data.Kelurahanpj,
	data.Kecamatanpj,
	data.Kabupatenpj,
	data.Kd_prop,
	data.Propinsipj,
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
		perusahaan_pasien = $25, nip = $26, email = $27, cacat_fisik = $28,
		kd_kel = $29, kd_kec = $30, kd_kab = $31, kelurahanpj = $32,
		kecamatanpj = $33, kabupatenpj = $34, kd_prop = $35, propinsipj = $36
	WHERE no_rkm_medis = $1
`
_, err := r.DB.Exec(query,
	data.No_rkm_medis,
	data.Nm_pasien,
	data.No_ktp,
	data.Jk,
	data.Tmp_lahir,
	data.Tgl_lahir,
	data.Nm_ibu,
	data.Alamat,
	data.Gol_darah,
	data.Pekerjaan,
	data.Stts_nikah,
	data.Agama,
	data.Tgl_daftar,
	data.No_tlp,
	data.Umur,
	data.Pnd,
	data.Keluarga,
	data.Namakeluarga,
	data.Kd_pj,
	data.No_peserta,
	data.Pekerjaanpj,
	data.Alamatpj,
	data.Suku_bangsa,
	data.Bahasa_pasien,
	data.Perusahaan_pasien,
	data.Nip,
	data.Email,
	data.Cacat_fisik,
	data.Kd_kel,
	data.Kd_kec,
	data.Kd_kab,
	data.Kelurahanpj,
	data.Kecamatanpj,
	data.Kabupatenpj,
	data.Kd_prop,
	data.Propinsipj,
)

	return err
}

func (r *RepositoryImpl) Delete(id string) error {
	query := `DELETE FROM pasien WHERE no_rkm_medis = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
