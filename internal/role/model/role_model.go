package model

type RoleRequest struct {
	Nama string `json:"nama" validate:"required,alphanum,max=20"`
}

type RoleResponse struct {
	Nama string `json:"nama"`
}
