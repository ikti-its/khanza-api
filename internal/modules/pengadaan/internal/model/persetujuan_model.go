package model

type PersetujuanRequest struct {
	IdPengajuan    string `json:"id_pengajuan" validate:"required,uuid4"`
	Status         string `json:"status" validate:"required,oneof=Disetujui Ditolak"`
	StatusApoteker string `json:"status_apoteker" validate:"required,oneof=Disetujui Ditolak"`
	StatusKeuangan string `json:"status_keuangan" validate:"required,oneof=Disetujui Ditolak"`
	Apoteker       string `json:"id_apoteker" validate:"uuid4"`
	Keuangan       string `json:"id_keuangan" validate:"uuid4"`
}

type PersetujuanResponse struct {
	IdPengajuan    string `json:"id_pengajuan"`
	Status         string `json:"status"`
	StatusApoteker string `json:"status_apoteker"`
	StatusKeuangan string `json:"status_keuangan"`
	Apoteker       string `json:"id_apoteker"`
	Keuangan       string `json:"id_keuangan"`
}

type PersetujuanPageResponse struct {
	Page        int                   `json:"page"`
	Size        int                   `json:"size"`
	Total       int                   `json:"total"`
	Persetujuan []PersetujuanResponse `json:"persetujuan_pengajuan"`
}
