package model

type DarahRequest struct {
	IdMedis     string `json:"id_barang_medis" validate:"required,uuid4"`
	Jenis       string `json:"jenis" validate:"required,oneof='Whole Blood (WB)' 'Packed Red Cell (PRC)' 'Thrombocyte Concentrate (TC)' 'Fresh Frozen Plasma (FFP)' 'Cryoprecipitate atau AHF' 'Leucodepleted (LD)' 'Leucoreduced (LR)'"`
	Keterangan  string `json:"keterangan"`
	Kadaluwarsa string `json:"kadaluwarsa" validate:"required"`
}

type DarahResponse struct {
	Id          string `json:"id"`
	IdMedis     string `json:"id_barang_medis"`
	Jenis       string `json:"jenis"`
	Keterangan  string `json:"keterangan"`
	Kadaluwarsa string `json:"kadaluwarsa"`
}

type DarahPageResponse struct {
	Page  int             `json:"page"`
	Size  int             `json:"size"`
	Total int             `json:"total"`
	Darah []DarahResponse `json:"darah"`
}
