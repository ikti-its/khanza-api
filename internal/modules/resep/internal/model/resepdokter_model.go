package model

type ResepDokter struct {
	NoResep     string  `json:"no_resep"`
	KodeBarang  string  `json:"kode_barang"`
	Jumlah      float64 `json:"jumlah"`
	AturanPakai string  `json:"aturan_pakai"`
}

type ResepDokterRequest struct {
	NoResep     string  `json:"no_resep" validate:"required"`
	KodeBarang  string  `json:"kode_barang" validate:"required"`
	Jumlah      float64 `json:"jumlah" validate:"required"`
	AturanPakai string  `json:"aturan_pakai" validate:"required"`
}

type ResepDokterResponse struct {
	NoResep     string  `json:"no_resep"`
	KodeBarang  string  `json:"kode_barang"`
	Jumlah      float64 `json:"jumlah"`
	AturanPakai string  `json:"aturan_pakai"`
}

type ResepDokterPageResponse struct {
	Page         int                   `json:"page"`
	Size         int                   `json:"size"`
	Total        int                   `json:"total"`
	ResepDokters []ResepDokterResponse `json:"resep_dokter"`
}
