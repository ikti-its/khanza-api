package model

type PengajuanRequest struct {
	Tanggal      string  `json:"tanggal_pengajuan" validate:"required"`
	Nomor        string  `json:"nomor_pengajuan" validate:"required"`
	Supplier     int     `json:"id_supplier" validate:"required,numeric"`
	Pegawai      string  `json:"id_pegawai" validate:"required,uuid4"`
	DiskonPersen float64 `json:"diskon_persen" validate:"required"`
	DiskonJumlah float64 `json:"diskon_jumlah" validate:"required"`
	PajakPersen  float64 `json:"pajak_persen" validate:"required"`
	PajakJumlah  float64 `json:"pajak_jumlah" validate:"required"`
	Materai      float64 `json:"materai" validate:"required"`
	Catatan      string  `json:"catatan"`
	Status       string  `json:"status_pesanan" validate:"required,oneof=0 1 2 3 4 5"`
}

type PengajuanResponse struct {
	Id           string  `json:"id"`
	Tanggal      string  `json:"tanggal_pengajuan"`
	Nomor        string  `json:"nomor_pengajuan"`
	Supplier     int     `json:"id_supplier"`
	Pegawai      string  `json:"id_pegawai"`
	DiskonPersen float64 `json:"diskon_persen"`
	DiskonJumlah float64 `json:"diskon_jumlah"`
	PajakPersen  float64 `json:"pajak_persen"`
	PajakJumlah  float64 `json:"pajak_jumlah"`
	Materai      float64 `json:"materai"`
	Catatan      string  `json:"catatan"`
	Status       string  `json:"status_pesanan"`
}

type PengajuanPageResponse struct {
	Page      int                 `json:"page"`
	Size      int                 `json:"size"`
	Total     int                 `json:"total"`
	Pengajuan []PengajuanResponse `json:"pengajuan_barang_medis"`
}
