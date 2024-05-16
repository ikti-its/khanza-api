package model

type TransaksiRequest struct {
	IdStok  string `json:"id_stok_keluar" validate:"required,uuid4"`
	IdMedis string `json:"id_barang_medis" validate:"required,uuid4"`
	Batch   string `json:"no_batch"`
	Faktur  string `json:"no_faktur"`
	Jumlah  int    `json:"jumlah_keluar" validate:"required,numeric"`
}

type TransaksiResponse struct {
	Id      string `json:"id"`
	IdStok  string `json:"id_stok_keluar"`
	IdMedis string `json:"id_barang_medis"`
	Batch   string `json:"no_batch"`
	Faktur  string `json:"no_faktur"`
	Jumlah  int    `json:"jumlah_keluar"`
}

type TransaksiPageResponse struct {
	Page      int                 `json:"page"`
	Size      int                 `json:"size"`
	Total     int                 `json:"total"`
	Transaksi []TransaksiResponse `json:"transaksi_keluar_barang_medis"`
}
