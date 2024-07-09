package model

type RoleResponse struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
}

type JabatanResponse struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
}

type DepartemenResponse struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
}

type StatusAktifResponse struct {
	Id   string `json:"id"`
	Nama string `json:"nama"`
}

type ShiftResponse struct {
	Id   string `json:"id"`
	Nama string `json:"nama"`
}

type AlasanCutiResponse struct {
	Id   string `json:"id"`
	Nama string `json:"nama"`
}
