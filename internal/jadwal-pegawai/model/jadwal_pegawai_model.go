package model

type JadwalPegawaiRequest struct {
	NIP       string `json:"nip" validate:"required,alphanum,max=5"`
	Tahun     int16  `json:"tahun" validate:"required,numeric"`
	Bulan     int16  `json:"bulan" validate:"required,numeric"`
	Hari      int16  `json:"hari" validate:"required,numeric"`
	ShiftNama string `json:"shift_nama" validate:"required,alpha,max=10"`
}

type JadwalPegawaiResponse struct {
	NIP       string `json:"nip"`
	Tahun     int16  `json:"tahun"`
	Bulan     int16  `json:"bulan"`
	Hari      int16  `json:"hari"`
	ShiftNama string `json:"shift_nama"`
}
