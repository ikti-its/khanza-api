package model

type MedisRequest struct {
	Nama   string  `json:"nama" validate:"required"`
	Jenis  string  `json:"jenis" validate:"required,oneof=Obat 'Alat Kesehatan' 'Bahan Habis Pakai' Darah"`
	Satuan int     `json:"satuan" validate:"required,numeric"`
	Harga  float64 `json:"harga" validate:"numeric"`
	Stok   int     `json:"stok" validate:"numeric"`
}

type MedisResponse struct {
	Id     string  `json:"id"`
	Nama   string  `json:"nama"`
	Jenis  string  `json:"jenis"`
	Satuan int     `json:"satuan"`
	Harga  float64 `json:"harga"`
	Stok   int     `json:"stok"`
}

type MedisPageResponse struct {
	Page  int             `json:"page"`
	Size  int             `json:"size"`
	Total int             `json:"total"`
	Medis []MedisResponse `json:"barang_medis"`
}
