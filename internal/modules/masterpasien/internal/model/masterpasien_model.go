package model

import "time"

type Model struct {
	NoRM             string    `json:"no_rkm_medis" db:"no_rkm_medis"`
	NamaPasien       string    `json:"nama_pasien" db:"nm_pasien"`
	NoKTP            string    `json:"no_ktp" db:"no_ktp"`
	JenisKelamin     string    `json:"jk" db:"jk"`
	TempatLahir      string    `json:"tmp_lahir" db:"tmp_lahir"`
	TanggalLahir     time.Time `json:"tgl_lahir" db:"tgl_lahir"`
	NamaIbu          string    `json:"nm_ibu" db:"nm_ibu"`
	Alamat           string    `json:"alamat" db:"alamat"`
	GolonganDarah    string    `json:"gol_darah" db:"gol_darah"`
	Pekerjaan        string    `json:"pekerjaan" db:"pekerjaan"`
	StatusPernikahan string    `json:"stts_nikah" db:"stts_nikah"`
	Agama            string    `json:"agama" db:"agama"`
	TanggalDaftar    time.Time `json:"tgl_daftar" db:"tgl_daftar"`
	NoTelepon        string    `json:"no_tlp" db:"no_tlp"`
	Umur             string    `json:"umur" db:"umur"`
	Pendidikan       string    `json:"pnd" db:"pnd"`
	PenanggungJawab  string    `json:"keluarga" db:"keluarga"`
	NamaPJ           string    `json:"namakeluarga" db:"namakeluarga"`
	HubunganPJ       string    `json:"kd_kel" db:"kd_kel"`
	Asuransi         string    `json:"kd_pj" db:"kd_pj"`
	NoPeserta        string    `json:"no_peserta" db:"no_peserta"`
	PekerjaanPJ      string    `json:"pekerjaanpj" db:"pekerjaanpj"`
	AlamatPJ         string    `json:"alamatpj" db:"alamatpj"`
	Suku             string    `json:"suku_bangsa" db:"suku_bangsa"`
	Bahasa           string    `json:"bahasa_pasien" db:"bahasa_pasien"`
	Instansi         string    `json:"perusahaan_pasien" db:"perusahaan_pasien"`
	NIPNRP           string    `json:"nip" db:"nip"`
	Email            string    `json:"email" db:"email"`
	CacatFisik       string    `json:"cacat_fisik" db:"cacat_fisik"`
}

type Request = Model
type Response = Model
