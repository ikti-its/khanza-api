package model

type BhpRequest struct {
	IdMedis     string `json:"id_barang_medis" validate:"required,uuid4"`
	Satuan      string `json:"satuan" validate:"required,oneof=pasang kotak paket item botol tabung ml"`
	Jumlah      int    `json:"jumlah" validate:"required,numeric"`
	Kadaluwarsa string `json:"kadaluwarsa" validate:"required"`
}

type BhpResponse struct {
	Id          string `json:"id"`
	IdMedis     string `json:"id_barang_medis"`
	Satuan      string `json:"satuan"`
	Jumlah      int    `json:"jumlah"`
	Kadaluwarsa string `json:"kadaluwarsa"`
}

type BhpPageResponse struct {
	Page  int           `json:"page"`
	Size  int           `json:"size"`
	Total int           `json:"total"`
	Bhp   []BhpResponse `json:"bahan_habis_pakai"`
}
