package model

type ResepDokterRacikanDetail struct {
	NoResep   string  `json:"no_resep"`
	NoRacik   string  `json:"no_racik"`
	KodeBrng  string  `json:"kode_brng"`
	P1        float64 `json:"p1"`
	P2        float64 `json:"p2"`
	Kandungan string  `json:"kandungan"`
	Jml       float64 `json:"jml"`
}

type ResepDokterRacikanDetailRequest struct {
	NoResep   string  `json:"no_resep" validate:"required"`
	NoRacik   string  `json:"no_racik" validate:"required"`
	KodeBrng  string  `json:"kode_brng" validate:"required"`
	P1        float64 `json:"p1"`
	P2        float64 `json:"p2"`
	Kandungan string  `json:"kandungan"`
	Jml       float64 `json:"jml"`
}

type ResepDokterRacikanDetailResponse struct {
	NoResep   string  `json:"no_resep"`
	NoRacik   string  `json:"no_racik"`
	KodeBrng  string  `json:"kode_brng"`
	P1        float64 `json:"p1"`
	P2        float64 `json:"p2"`
	Kandungan string  `json:"kandungan"`
	Jml       float64 `json:"jml"`
}

type ResepDokterRacikanDetailPageResponse struct {
	Page                      int                                `json:"page"`
	Size                      int                                `json:"size"`
	Total                     int                                `json:"total"`
	ResepDokterRacikanDetails []ResepDokterRacikanDetailResponse `json:"resep_dokter_racikan_detail"`
}
