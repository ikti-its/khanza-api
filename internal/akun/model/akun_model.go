package model

type AkunRequest struct {
	NIP      string `json:"nip" validate:"required,alphanum,max=5"`
	Email    string `json:"email" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	RoleNama string `json:"role_nama" validate:"required,alphanum,max=20"`
}

type AkunUpdateRequest struct {
	Email    string `json:"email" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

type AkunResponse struct {
	NIP      string `json:"nip"`
	Email    string `json:"email"`
	RoleNama string `json:"role_nama"`
}

type AkunPageResponse struct {
	Akun  []AkunResponse `json:"akun"`
	Page  int            `json:"page"`
	Size  int            `json:"size"`
	Total int            `json:"total"`
}
