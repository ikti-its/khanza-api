package entity

type PasienMeninggal struct {
	NoRkmMedis    string    `json:"no_rkm_medis" db:"no_rkm_medis"`
	NmPasien      string    `json:"nm_pasien" db:"nm_pasien"`
	JK            string    `json:"jk" db:"jk"`
	TglLahir      string    `json:"tgl_lahir" db:"tgl_lahir"`
	Umur          string    `json:"umur" db:"umur"`
	GolDarah      string    `json:"gol_darah" db:"gol_darah"`
	SttsNikah     string    `json:"stts_nikah" db:"stts_nikah"`
	Agama         string    `json:"agama" db:"agama"`
	Tanggal       string    `json:"tanggal" db:"tanggal"`
	Jam           string    `json:"jam" db:"jam"`
	ICDX          string    `json:"icdx" db:"icdx"`
	ICDXAntara1   string    `json:"icdx_antara1" db:"icdx_antara1"`
	ICDXAntara2   string    `json:"icdx_antara2" db:"icdx_antara2"`
	ICDXLangsung  string    `json:"icdx_langsung" db:"icdx_langsung"`
	Keterangan    string    `json:"keterangan" db:"keterangan"`
	NamaDokter    string    `json:"nama_dokter" db:"nama_dokter"`
	KodeDokter    string    `json:"kode_dokter" db:"kode_dokter"`
}
