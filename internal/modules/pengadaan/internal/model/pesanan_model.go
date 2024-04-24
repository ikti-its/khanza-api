package model

type PesananRequest struct {
	IdPengajuan string  `json:"id_pengajuan" validate:"required,uuid4"`
	IdMedis     string  `json:"id_barang_medis" validate:"required,uuid4"`
	Harga       float64 `json:"harga_satuan" validate:"required"`
	Pesanan     int     `json:"jumlah_pesanan" validate:"required,numeric"`
	Diterima    int     `json:"jumlah_diterima" validate:"numeric"`
	Kadaluwarsa string  `json:"kadaluwarsa"`
	Batch       string  `json:"no_batch"`
}

type PesananResponse struct {
	Id          string  `json:"id"`
	IdPengajuan string  `json:"id_pengajuan"`
	IdMedis     string  `json:"id_barang_medis"`
	Harga       float64 `json:"harga_satuan"`
	Pesanan     int     `json:"jumlah_pesanan"`
	Diterima    int     `json:"jumlah_diterima"`
	Kadaluwarsa string  `json:"kadaluwarsa"`
	Batch       string  `json:"no_batch"`
}

type PesananPageResponse struct {
	Page    int               `json:"page"`
	Size    int               `json:"size"`
	Total   int               `json:"total"`
	Pesanan []PesananResponse `json:"pesanan_barang_medis"`
}
