package model

type StokRequest struct {
	Nomor         string `json:"no_keluar" validate:"required"`
	IdPegawai     string `json:"id_pegawai" validate:"required,uuid4"`
	Tanggal       string `json:"tanggal_stok_keluar" validate:"required"`
	AsalRuangan   int    `json:"asal_ruangan" validate:"required,numeric"`
	TujuanRuangan int    `json:"tujuan_ruangan" validate:"required,numeric"`
	Keterangan    string `json:"keterangan"`
}

type StokResponse struct {
	Id            string `json:"id"`
	Nomor         string `json:"no_keluar"`
	IdPegawai     string `json:"id_pegawai"`
	Tanggal       string `json:"tanggal_stok_keluar"`
	AsalRuangan   int    `json:"asal_ruangan"`
	TujuanRuangan int    `json:"tujuan_ruangan"`
	Keterangan    string `json:"keterangan"`
}

type StokPageResponse struct {
	Page  int            `json:"page"`
	Size  int            `json:"size"`
	Total int            `json:"total"`
	Stok  []StokResponse `json:"stok_keluar_barang_medis"`
}
