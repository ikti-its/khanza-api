package model

type StokRequest struct {
	Nomor      string `json:"no_keluar" validate:"required"`
	IdPegawai  string `json:"id_pegawai" validate:"required,uuid4"`
	Tanggal    string `json:"tanggal_stok_keluar" validate:"required"`
	Keterangan string `json:"keterangan"`
}

type StokResponse struct {
	Id         string `json:"id"`
	Nomor      string `json:"no_keluar"`
	IdPegawai  string `json:"id_pegawai"`
	Tanggal    string `json:"tanggal_stok_keluar"`
	Keterangan string `json:"keterangan"`
}

type StokPageResponse struct {
	Page  int            `json:"page"`
	Size  int            `json:"size"`
	Total int            `json:"total"`
	Stok  []StokResponse `json:"stok_keluar_barang_medis"`
}
