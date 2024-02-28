package model

type DepartemenRequest struct {
	Nama string `json:"nama" validate:"required,alpha,max=25"`
}

type DepartemenResponse struct {
	Nama string `json:"nama"`
}
