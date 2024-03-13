package model

type CutiCreateRequest struct {
	NIP            string `json:"nip" validate:"required"`
	TanggalMulai   string `json:"tanggal_mulai" validate:"required"`
	TanggalSelesai string `json:"tanggal_selesai" validate:"required"`
	Keterangan     string `json:"keterangan" validate:"required"`
}

type CutiUpdateRequest struct {
	NIP            string `json:"nip" validate:"required"`
	TanggalMulai   string `json:"tanggal_mulai" validate:"required"`
	TanggalSelesai string `json:"tanggal_selesai" validate:"required"`
	Keterangan     string `json:"keterangan" validate:"required"`
	Status         bool   `json:"status" validate:"required"`
}

type CutiResponse struct {
	ID             string `json:"id"`
	NIP            string `json:"nip"`
	TanggalMulai   string `json:"tanggal_mulai"`
	TanggalSelesai string `json:"tanggal_selesai"`
	Keterangan     string `json:"keterangan"`
	Status         bool   `json:"status"`
}
