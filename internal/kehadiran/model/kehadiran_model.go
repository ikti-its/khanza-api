package model

type KehadiranRequest struct {
	NIP       string `json:"nip" validate:"required"`
	Tanggal   string `json:"tanggal" validate:"required"`
	ShiftNama string `json:"shift_nama" validate:"required"`
}

type KehadiranUpdateRequest struct {
	NIP        string `json:"nip" validate:"required"`
	Tanggal    string `json:"tanggal" validate:"required"`
	ShiftNama  string `json:"shift_nama" validate:"required"`
	JamMasuk   string `json:"jam_masuk"`
	JamKeluar  string `json:"jam_keluar"`
	Keterangan string `json:"keterangan"`
}

type KehadiranShiftResponse struct {
	Nama      string `json:"nama"`
	JamMasuk  string `json:"jam_masuk"`
	JamKeluar string `json:"jam_keluar"`
}

type KehadiranResponse struct {
	ID         string                 `json:"id"`
	NIP        string                 `json:"nip"`
	Tanggal    string                 `json:"tanggal"`
	Shift      KehadiranShiftResponse `json:"shift"`
	JamMasuk   string                 `json:"jam_masuk"`
	JamKeluar  string                 `json:"jam_keluar"`
	Keterangan string                 `json:"keterangan"`
}
