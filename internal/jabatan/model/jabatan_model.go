package model

type JabatanRequest struct {
	Nama      string  `json:"nama" validate:"required;alpha;max=25"`
	Jenjang   string  `json:"jenjang" validate:"required;alpha;max=25"`
	GajiPokok float64 `json:"gaji_pokok" validate:"positive"`
	Tunjangan float64 `json:"tunjangan" validate:"positive"`
}

type JabatanResponse struct {
	Nama      string  `json:"nama"`
	Jenjang   string  `json:"jenjang"`
	GajiPokok float64 `json:"gaji_pokok"`
	Tunjangan float64 `json:"tunjangan"`
}
