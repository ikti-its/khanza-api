package model

type RoleRequest struct {
	ID   int    `json:"id"`
	Role string `json:"role" validate:"required,alphanum,max=20"`
}

type RoleResponse struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
}
