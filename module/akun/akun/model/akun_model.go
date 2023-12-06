package model

type AkunRequest struct {
	NIP      string `json:"nip" validate:"required,alphanum,max=5"`
	Email    string `json:"email" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required"`
	RoleID   int    `json:"role_id" validate:"required,numeric"`
}

type AkunResponse struct {
	NIP    string `json:"nip"`
	Email  string `json:"email"`
	RoleID int    `json:"role_id"`
}
