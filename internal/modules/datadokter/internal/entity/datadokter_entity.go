package entity

type Dokter struct {
	Kode_dokter    string `json:"kode_dokter" db:"kode_dokter"`
	Nama_dokter    string `json:"nama_dokter" db:"nama_dokter"`
	Jenis_kelamin  string `json:"jenis_kelamin" db:"jenis_kelamin"`
	Alamat_tinggal string `json:"alamat_tinggal" db:"alamat_tinggal"`
	No_telp        string `json:"no_telp" db:"no_telp"`
	Spesialis      string `json:"spesialis" db:"spesialis"`
	Izin_praktik   string `json:"izin_praktik" db:"izin_praktik"`
}
