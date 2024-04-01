package model

type AlamatRequest struct {
	IdAkun    string  `json:"id_akun" validate:"required,uuid4"`
	Alamat    string  `json:"alamat" validate:"required"`
	AlamatLat float64 `json:"alamat_lat" validate:"required,latitude"`
	AlamatLon float64 `json:"alamat_lon" validate:"required,longitude"`
	Kota      string  `json:"kota" validate:"required"`
	KodePos   string  `json:"kode_pos" validate:"required,max=5"`
}

type AlamatResponse struct {
	IdAkun    string  `json:"id_akun"`
	Alamat    string  `json:"alamat"`
	AlamatLat float64 `json:"alamat_lat"`
	AlamatLon float64 `json:"alamat_lon"`
	Kota      string  `json:"kota"`
	KodePos   string  `json:"kode_pos"`
}
