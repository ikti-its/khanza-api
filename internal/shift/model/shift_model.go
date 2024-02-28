package model

type ShiftRequest struct {
	Nama      string `json:"nama" validate:"required,alpha,max=10"`
	JamMasuk  string `json:"jam_masuk" validate:"required"`
	JamKeluar string `json:"jam_keluar" validate:"required"`
}

type ShiftResponse struct {
	Nama      string `json:"nama"`
	JamMasuk  string `json:"jam_masuk"`
	JamKeluar string `json:"jam_keluar"`
}
