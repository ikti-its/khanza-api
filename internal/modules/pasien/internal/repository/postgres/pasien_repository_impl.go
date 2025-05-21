package postgres

import (
	"github.com/ikti-its/khanza-api/internal/modules/pasien/internal/entity"
	"github.com/jmoiron/sqlx"
)

type PasienRepository struct {
	DB *sqlx.DB
}

func NewPasienRepository(db *sqlx.DB) *PasienRepository {
	return &PasienRepository{DB: db}
}

func (r *PasienRepository) Insert(pasien *entity.Pasien) error {
	query := `
		INSERT INTO pasien (
			no_rkm_medis, nm_pasien, no_ktp, jk, tmp_lahir, tgl_lahir,
			nm_ibu, alamat, gol_darah, pekerjaan, stts_nikah, agama, tgl_daftar,
			no_tlp, umur, pnd, keluarga, namakeluarga, kd_pj, no_peserta,
			kd_kel, kd_kec, kd_kab, pekerjaanpj, alamatpj, kelurahanpj, kecamatanpj,
			kabupatenpj, perusahaan_pasien, suku_bangsa, bahasa_pasien, cacat_fisik,
			email, nip, kd_prop, propinsipj
		) VALUES (
			:no_rkm_medis, :nm_pasien, :no_ktp, :jk, :tmp_lahir, :tgl_lahir,
			:nm_ibu, :alamat, :gol_darah, :pekerjaan, :stts_nikah, :agama, :tgl_daftar,
			:no_tlp, :umur, :pnd, :keluarga, :namakeluarga, :kd_pj, :no_peserta,
			:kd_kel, :kd_kec, :kd_kab, :pekerjaanpj, :alamatpj, :kelurahanpj, :kecamatanpj,
			:kabupatenpj, :perusahaan_pasien, :suku_bangsa, :bahasa_pasien, :cacat_fisik,
			:email, :nip, :kd_prop, :propinsipj
		)
	`
	_, err := r.DB.NamedExec(query, pasien)
	return err
}

func (r *PasienRepository) Find() ([]entity.Pasien, error) {
	var list []entity.Pasien
	err := r.DB.Select(&list, "SELECT * FROM pasien")
	return list, err
}

func (r *PasienRepository) FindPage(page, size int) ([]entity.Pasien, int, error) {
	offset := (page - 1) * size
	var list []entity.Pasien
	query := `SELECT * FROM pasien ORDER BY no_rkm_medis LIMIT $1 OFFSET $2`
	err := r.DB.Select(&list, query, size, offset)
	if err != nil {
		return nil, 0, err
	}

	var total int
	countQuery := `SELECT COUNT(*) FROM pasien`
	err = r.DB.Get(&total, countQuery)
	return list, total, err
}

func (r *PasienRepository) FindByNoRkmMedis(noRkmMedis string) (entity.Pasien, error) {
	var pasien entity.Pasien
	query := `SELECT * FROM pasien WHERE no_rkm_medis = $1`
	err := r.DB.Get(&pasien, query, noRkmMedis)
	return pasien, err
}

func (r *PasienRepository) Update(pasien *entity.Pasien) error {
	query := `
		UPDATE pasien SET
			nm_pasien = :nm_pasien,
			no_ktp = :no_ktp,
			jk = :jk,
			tmp_lahir = :tmp_lahir,
			tgl_lahir = :tgl_lahir,
			nm_ibu = :nm_ibu,
			alamat = :alamat,
			gol_darah = :gol_darah,
			pekerjaan = :pekerjaan,
			stts_nikah = :stts_nikah,
			agama = :agama,
			tgl_daftar = :tgl_daftar,
			no_tlp = :no_tlp,
			umur = :umur,
			pnd = :pnd,
			keluarga = :keluarga,
			namakeluarga = :namakeluarga,
			kd_pj = :kd_pj,
			no_peserta = :no_peserta,
			kd_kel = :kd_kel,
			kd_kec = :kd_kec,
			kd_kab = :kd_kab,
			pekerjaanpj = :pekerjaanpj,
			alamatpj = :alamatpj,
			kelurahanpj = :kelurahanpj,
			kecamatanpj = :kecamatanpj,
			kabupatenpj = :kabupatenpj,
			perusahaan_pasien = :perusahaan_pasien,
			suku_bangsa = :suku_bangsa,
			bahasa_pasien = :bahasa_pasien,
			cacat_fisik = :cacat_fisik,
			email = :email,
			nip = :nip,
			kd_prop = :kd_prop,
			propinsipj = :propinsipj
		WHERE no_rkm_medis = :no_rkm_medis
	`
	_, err := r.DB.NamedExec(query, pasien)
	return err
}

func (r *PasienRepository) Delete(noRkmMedis string) error {
	query := `DELETE FROM pasien WHERE no_rkm_medis = $1`
	_, err := r.DB.Exec(query, noRkmMedis)
	return err
}

func (r *PasienRepository) GetByNoKTP(noKTP string) (*entity.Pasien, error) {
	var pasien entity.Pasien
	query := `SELECT * FROM pasien WHERE no_ktp = $1`
	err := r.DB.Get(&pasien, query, noKTP)
	if err != nil {
		return nil, err
	}
	return &pasien, nil
}

func (r *PasienRepository) GetByNoPeserta(noPeserta string) (*entity.Pasien, error) {
	var pasien entity.Pasien
	query := `SELECT * FROM pasien WHERE no_peserta = $1`
	err := r.DB.Get(&pasien, query, noPeserta)
	if err != nil {
		return nil, err
	}
	return &pasien, nil
}
