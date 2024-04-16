package model

type DarahRequest struct {
	IdMedis     string `json:"id_barang_medis" validate:"required,uuid4"`
	Keterangan  string `json:"keterangan"`
	Kadaluwarsa string `json:"kadaluwarsa" validate:"required"`
}

type DarahResponse struct {
	Id          string `json:"id"`
	IdMedis     string `json:"id_barang_medis"`
	Keterangan  string `json:"keterangan"`
	Kadaluwarsa string `json:"kadaluwarsa"`
}

type DarahPageResponse struct {
	Page  int             `json:"page"`
	Size  int             `json:"size"`
	Total int             `json:"total"`
	Darah []DarahResponse `json:"darah"`
}
