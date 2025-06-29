package entity



type MasterPasien struct {
	No_rkm_medis             string    `json:"no_rkm_medis" db:"no_rkm_medis"`
	Nm_pasien       string    `json:"nm_pasien" db:"nm_pasien"`
	No_ktp            string    `json:"no_ktp" db:"no_ktp"`
	Jk     string    `json:"jk" db:"jk"`
	Tmp_lahir      string    `json:"tmp_lahir" db:"tmp_lahir"`
	Tgl_lahir     string `json:"tgl_lahir" db:"tgl_lahir"`
	Nm_ibu          string    `json:"nm_ibu" db:"nm_ibu"`
	Alamat           string    `json:"alamat" db:"alamat"`
	Gol_darah    string    `json:"gol_darah" db:"gol_darah"`
	Pekerjaan        string    `json:"pekerjaan" db:"pekerjaan"`
	Stts_nikah string    `json:"stts_nikah" db:"stts_nikah"`
	Agama            string    `json:"agama" db:"agama"`
	Tgl_daftar    string `json:"tgl_daftar" db:"tgl_daftar"`
	No_tlp        string    `json:"no_tlp" db:"no_tlp"`
	Umur             string    `json:"umur" db:"umur"`
	Pnd       string    `json:"pnd" db:"pnd"`
	Keluarga  string    `json:"keluarga" db:"keluarga"`
	Namakeluarga          string    `json:"namakeluarga" db:"namakeluarga"`
	Kd_pj        string    `json:"kd_pj" db:"kd_pj"`
	No_peserta        string    `json:"no_peserta" db:"no_peserta"`
	Kd_kel           string    `json:"kd_kel" db:"kd_kel"`
	Kd_kec           string    `json:"kd_kec" db:"kd_kec"`
	Kd_kab           string    `json:"kd_kab" db:"kd_kab"`
	Pekerjaanpj      string    `json:"pekerjaanpj" db:"pekerjaanpj"`
	Alamatpj         string    `json:"alamatpj" db:"alamatpj"`
	Kelurahanpj      string    `json:"kelurahanpj" db:"kelurahanpj"`
	Kecamatanpj      string    `json:"kecamatanpj" db:"kecamatanpj"`
	Kabupatenpj      string    `json:"kabupatenpj" db:"kabupatenpj"`
	Suku_bangsa             string    `json:"suku_bangsa" db:"suku_bangsa"`
	Bahasa_pasien           string    `json:"bahasa_pasien" db:"bahasa_pasien"`
	Perusahaan_pasien         string    `json:"perusahaan_pasien" db:"perusahaan_pasien"`
	Nip                     string    `json:"nip" db:"nip"`
	Email                   string    `json:"email" db:"email"`
	Cacat_fisik             string    `json:"cacat_fisik" db:"cacat_fisik"`
	Kd_prop                 string    `json:"kd_prop" db:"kd_prop"`
	Propinsipj              string    `json:"propinsipj" db:"propinsipj"`
}
