package model

type PemesananRequest struct {
	Tanggal     string `json:"tanggal_pesan" validate:"required"`
	Nomor       string `json:"no_pemesanan" validate:"required"`
	IdPengajuan string `json:"id_pengajuan" validate:"required,uuid4"`
	IdPegawai   string `json:"id_pegawai" validate:"required,uuid4"`
}

type PemesananResponse struct {
	Id          string `json:"id"`
	Tanggal     string `json:"tanggal_pesan"`
	Nomor       string `json:"no_pemesanan"`
	IdPengajuan string `json:"id_pengajuan"`
	IdPegawai   string `json:"id_pegawai"`
}

type PemesananPageResponse struct {
	Page      int                 `json:"page"`
	Size      int                 `json:"size"`
	Total     int                 `json:"total"`
	Pemesanan []PemesananResponse `json:"pemesanan_barang_medis"`
}
