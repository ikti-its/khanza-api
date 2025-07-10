package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ikti-its/khanza-api/internal/modules/kelahiranbayi/internal/entity"
)

type Repository interface {
	FindAll() ([]entity.KelahiranBayi, error)
	FindById(noRM string) (entity.KelahiranBayi, error)
	Insert(data *entity.KelahiranBayi) error
	Update(data *entity.KelahiranBayi) error
	Delete(noRM string) error
}

type RepositoryImpl struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{DB: db}
}

func (r *RepositoryImpl) FindAll() ([]entity.KelahiranBayi, error) {
	query := `SELECT * FROM sik.kelahiran_bayi ORDER BY tgl_lahir DESC`

	var records []entity.KelahiranBayi
	err := r.DB.Select(&records, query)
	return records, err
}

func (r *RepositoryImpl) FindById(noRM string) (entity.KelahiranBayi, error) {
	query := `SELECT * FROM sik.kelahiran_bayi WHERE no_rkm_medis = $1`

	var record entity.KelahiranBayi
	err := r.DB.Get(&record, query, noRM)
	return record, err
}

func (r *RepositoryImpl) Insert(data *entity.KelahiranBayi) error {
	query := `
		INSERT INTO sik.kelahiran_bayi (
			no_rkm_medis, nm_pasien, jk, tmp_lahir, tgl_lahir, jam, umur, tgl_daftar,
			nm_ibu, umur_ibu, nm_ayah, umur_ayah, alamat,
			bb, pb, proses_lahir, kelahiran_ke, keterangan, diagnosa, penyulit_kehamilan, ketuban,
			lk_perut, lk_kepala, lk_dada, penolong, no_skl, gravida, para, abortus,
			f1, u1, t1, r1, w1, n1,
			f5, u5, t5, r5, w5, n5,
			f10, u10, t10, r10, w10, n10,
			resusitas, obat, mikasi, mikonium
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7,
			$8, $9, $10, $11, $12,
			$13, $14, $15, $16, $17, $18, $19, $20,
			$21, $22, $23, $24, $25, $26, $27,
			$28, $29, $30, $31, $32, $33,
			$34, $35, $36, $37, $38, $39,
			$40, $41, $42, $43, $44, $45,
			$46, $47, $48, $49, $50, $51
		)
	`

	_, err := r.DB.Exec(query,
		data.NoRkmMedis, data.NmPasien, data.JK, data.TmpLahir, data.TglLahir, data.Jam, data.Umur, data.TglDaftar,
		data.NmIbu, data.UmurIbu, data.NmAyah, data.UmurAyah, data.Alamat,
		data.BB, data.PB, data.ProsesLahir, data.KelahiranKe, data.Keterangan, data.Diagnosa, data.PenyulitKehamilan, data.Ketuban,
		data.LKPerut, data.LKKepala, data.LKDada, data.Penolong, data.NoSKL, data.Gravida, data.Para, data.Abortus,
		data.F1, data.U1, data.T1, data.R1, data.W1, data.N1,
		data.F5, data.U5, data.T5, data.R5, data.W5, data.N5,
		data.F10, data.U10, data.T10, data.R10, data.W10, data.N10,
		data.Resusitas, data.Obat, data.Mikasi, data.Mikonium,
	)

	return err
}

func (r *RepositoryImpl) Update(data *entity.KelahiranBayi) error {
	query := `
		UPDATE sik.kelahiran_bayi SET
			nm_pasien = $2, jk = $3, tgl_lahir = $4, jam = $5, umur = $6, tmp_lahir = $7, tgl_daftar = $8,
			nm_ibu = $9, umur_ibu = $10, nm_ayah = $11, umur_ayah = $12, alamat = $13,
			bb = $14, pb = $15, proses_lahir = $16, kelahiran_ke = $17, keterangan = $18, diagnosa = $19,
			penyulit_kehamilan = $20, ketuban = $21, lk_perut = $22, lk_kepala = $23, lk_dada = $24, penolong = $25, no_skl = $26,
			gravida = $27, para = $28, abortus = $29,
			f1 = $30, u1 = $31, t1 = $32, r1 = $33, w1 = $34, n1 = $35,
			f5 = $36, u5 = $37, t5 = $38, r5 = $39, w5 = $40, n5 = $41,
			f10 = $42, u10 = $43, t10 = $44, r10 = $45, w10 = $46, n10 = $47,
			resusitas = $48, obat = $49, mikasi = $50, mikonium = $51
		WHERE no_rkm_medis = $1
	`

	_, err := r.DB.Exec(query,
		data.NoRkmMedis, data.NmPasien, data.JK, data.TmpLahir, data.TglLahir, data.Jam, data.Umur, data.TglDaftar,
		data.NmIbu, data.UmurIbu, data.NmAyah, data.UmurAyah, data.Alamat,
		data.BB, data.PB, data.ProsesLahir, data.KelahiranKe, data.Keterangan, data.Diagnosa, data.PenyulitKehamilan, data.Ketuban,
		data.LKPerut, data.LKKepala, data.LKDada, data.Penolong, data.NoSKL, data.Gravida, data.Para, data.Abortus,
		data.F1, data.U1, data.T1, data.R1, data.W1, data.N1,
		data.F5, data.U5, data.T5, data.R5, data.W5, data.N5,
		data.F10, data.U10, data.T10, data.R10, data.W10, data.N10,
		data.Resusitas, data.Obat, data.Mikasi, data.Mikonium,
	)

	return err
}

func (r *RepositoryImpl) Delete(noRM string) error {
	query := `DELETE FROM sik.kelahiran_bayi WHERE no_rkm_medis = $1`
	_, err := r.DB.Exec(query, noRM)
	return err
}
