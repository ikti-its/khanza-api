package model

type AlkesRequest struct {
	IdMedis string `json:"id_barang_medis" validate:"required,uuid4"`
	Merek   string `json:"merek" validate:"required,oneof=Omron Philips 'GE Healthcare' 'Siemens Healthineers' Medtronic 'Johnson & Johnson' 'Becton\, Dickinson and Company (BD)' Stryker 'Boston Scientific' 'Olympus Corporation' 'Roche Diagnostics'"`
}

type AlkesResponse struct {
	Id      string `json:"id"`
	IdMedis string `json:"id_barang_medis"`
	Merek   string `json:"merek"`
}

type AlkesPageResponse struct {
	Page  int
	Size  int
	Total int
	Alkes []AlkesResponse
}
