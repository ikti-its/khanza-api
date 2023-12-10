package model

type AkunRequest struct {
	NIP      string `json:"nip" validate:"required,alphanum,max=5"`
	Email    string `json:"email" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	RoleName string `json:"role_name" validate:"required,alphanum,max=20"`
}

type AkunUpdateRequest struct {
	Email    string `json:"email" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

type AkunResponse struct {
	NIP      string `json:"nip"`
	Email    string `json:"email"`
	RoleName string `json:"role_name"`
}
