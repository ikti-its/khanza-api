package model

type StokRequest struct {
	Nomor      string `json:"no_keluar" validate:"required"`
	IdMedis    string `json:"id_barang_medis" validate:"required,uuid4"`
	IdPegawai  string `json:"id_pegawai" validate:"required,uuid4"`
	Tanggal    string `json:"tanggal_stok_keluar" validate:"required"`
	Jumlah     int    `json:"jumlah_keluar" validate:"required,numeric"`
	Keterangan string `json:"keterangan"`
}

type StokResponse struct {
	Id         string `json:"id"`
	Nomor      string `json:"no_keluar"`
	IdMedis    string `json:"id_barang_medis"`
	IdPegawai  string `json:"id_pegawai"`
	Tanggal    string `json:"tanggal_stok_keluar"`
	Jumlah     int    `json:"jumlah_keluar"`
	Keterangan string `json:"keterangan"`
}

type StokPageResponse struct {
	Page  int            `json:"page"`
	Size  int            `json:"size"`
	Total int            `json:"total"`
	Stok  []StokResponse `json:"stok_keluar_barang_medis"`
}
