package model

type RoleRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required,alphanum,max=20"`
}

type RoleResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
