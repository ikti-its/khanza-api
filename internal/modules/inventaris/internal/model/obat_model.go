package model

type ObatRequest struct {
	IdMedis     string `json:"id_barang_medis" validate:"required,uuid4"`
	Industri    int    `json:"industri_farmasi" validate:"required,numeric"`
	Kandungan   string `json:"kandungan" validate:"required"`
	SatuanBesar string `json:"satuan_besar" validate:"required,oneof=tablet kapsul ampul botol tube vial injeksi"`
	SatuanKecil string `json:"satuan_kecil" validate:"required,oneof=tablet kapsul ampul botol tube vial injeksi"`
	Isi         int    `json:"isi" validate:"required,numeric"`
	Kapasitas   int    `json:"kapasitas" validate:"required,numeric"`
	Jenis       int    `json:"jenis" validate:"required,numeric"`
	Kategori    int    `json:"kategori" validate:"required,numeric"`
	Golongan    int    `json:"golongan" validate:"required,numeric"`
	Kadaluwarsa string `json:"kadaluwarsa" validate:"required"`
}

type ObatResponse struct {
	Id          string `json:"id"`
	IdMedis     string `json:"id_barang_medis"`
	Industri    int    `json:"industri_farmasi"`
	Kandungan   string `json:"kandungan"`
	SatuanBesar string `json:"satuan_besar"`
	SatuanKecil string `json:"satuan_kecil"`
	Isi         int    `json:"isi"`
	Kapasitas   int    `json:"kapasitas"`
	Jenis       int    `json:"jenis"`
	Kategori    int    `json:"kategori"`
	Golongan    int    `json:"golongan"`
	Kadaluwarsa string `json:"kadaluwarsa"`
}

type ObatPageResponse struct {
	Page  int
	Size  int
	Total int
	Obat  []ObatResponse
}
