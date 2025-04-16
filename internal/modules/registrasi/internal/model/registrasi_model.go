package model

type Registrasi struct {
	NomorReg         string  `json:"nomor_reg" validate:"required"`
	NomorRawat       string  `json:"nomor_rawat"`
	Tanggal          string  `json:"tanggal"`
	Jam              string  `json:"jam"`
	KodeDokter       string  `json:"kode_dokter"`
	NamaDokter       string  `json:"nama_dokter"`
	NomorRM          string  `json:"nomor_rm"`
	Nama             string  `json:"nama_pasien"`
	JenisKelamin     string  `json:"jenis_kelamin" validate:"oneof=L P"` // L (Laki-laki), P (Perempuan)
	Umur             string  `json:"umur"`
	Poliklinik       string  `json:"poliklinik"`
	JenisBayar       string  `json:"jenis_bayar"`
	PenanggungJawab  string  `json:"penanggung_jawab"`
	Alamat           string  `json:"alamat_pj"`
	HubunganPJ       string  `json:"hubungan_pj"`
	BiayaRegistrasi  float64 `json:"biaya_registrasi"`
	StatusRegistrasi string  `json:"status_registrasi" validate:"oneof=baru lama"`
	NoTelepon        string  `json:"no_telepon" validate:"e164"` // Validates phone number format
	StatusRawat      string  `json:"status_rawat"`
	StatusPoli       string  `json:"status_poli"`
	StatusBayar      string  `json:"status_bayar"`
	StatusKamar      string  `json:"status_kamar"`
	Kelas            string  `json:"kelas"`
}

type RegistrasiRequest struct {
	NomorReg         string  `json:"nomor_reg" validate:"required"`
	NomorRawat       string  `json:"nomor_rawat"`
	Tanggal          string  `json:"tanggal"`
	Jam              string  `json:"jam"`
	KodeDokter       string  `json:"kode_dokter"`
	NamaDokter       string  `json:"nama_dokter"`
	NomorRM          string  `json:"nomor_rm"`
	Nama             string  `json:"nama_pasien"`
	JenisKelamin     string  `json:"jenis_kelamin" validate:"oneof=L P"` // L (Laki-laki), P (Perempuan)
	Umur             string  `json:"umur"`
	Poliklinik       string  `json:"poliklinik"`
	JenisBayar       string  `json:"jenis_bayar"`
	PenanggungJawab  string  `json:"penanggung_jawab"`
	Alamat           string  `json:"alamat_pj"`
	HubunganPJ       string  `json:"hubungan_pj"`
	BiayaRegistrasi  float64 `json:"biaya_registrasi"`
	StatusRegistrasi string  `json:"status_registrasi" validate:"oneof=baru lama"`
	NoTelepon        string  `json:"no_telepon" validate:"e164"` // Validates phone number format
	StatusRawat      string  `json:"status_rawat"`
	StatusPoli       string  `json:"status_poli"`
	StatusBayar      string  `json:"status_bayar"`
	StatusKamar      string  `json:"status_kamar"`
	Kelas            string  `json:"kelas"`
}

type RegistrasiResponse struct {
	NomorReg         string  `json:"nomor_reg"`
	NomorRawat       string  `json:"nomor_rawat"`
	Tanggal          string  `json:"tanggal"`
	Jam              string  `json:"jam"`
	KodeDokter       string  `json:"kode_dokter"`
	NamaDokter       string  `json:"nama_dokter"`
	NomorRM          string  `json:"nomor_rm"`
	Nama             string  `json:"nama_pasien"`
	JenisKelamin     string  `json:"jenis_kelamin"`
	Umur             string  `json:"umur"`
	Poliklinik       string  `json:"poliklinik"`
	JenisBayar       string  `json:"jenis_bayar"`
	PenanggungJawab  string  `json:"penanggung_jawab,omitempty"`
	Alamat           string  `json:"alamat_pj,omitempty"`
	HubunganPJ       string  `json:"hubungan_pj,omitempty"`
	BiayaRegistrasi  float64 `json:"biaya_registrasi"`
	StatusRegistrasi string  `json:"status_registrasi"`
	NoTelepon        string  `json:"no_telepon"`
	StatusRawat      string  `json:"status_rawat"`
	StatusPoli       string  `json:"status_poli"`
	StatusBayar      string  `json:"status_bayar"`
	StatusKamar      string  `json:"status_kamar"`
	Kelas            string  `json:"kelas"`
}

type PendingRoomRequest struct {
	NomorReg   string `db:"nomor_reg" json:"nomor_reg"`
	NamaPasien string `db:"nama_pasien" json:"nama_pasien"`
	Kelas      string `db:"kelas" json:"kelas"`
}

type RegistrasiPageResponse struct {
	Page       int                  `json:"page"`
	Size       int                  `json:"size"`
	Total      int                  `json:"total"`
	Registrasi []RegistrasiResponse `json:"registrasi"`
}
