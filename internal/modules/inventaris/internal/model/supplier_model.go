package model

type SupplierResponse struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Alamat   string `json:"alamat"`
	Telepon  string `json:"telepon"`
	Kota     string `json:"kota"`
	Bank     string `json:"bank"`
	Rekening string `json:"rekening"`
}
