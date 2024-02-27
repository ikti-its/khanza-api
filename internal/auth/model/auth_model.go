package model

type AuthRequest struct {
	NIP      string `json:"nip" validate:"required,alphanum,max=5"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Token   string `json:"token"`
	Expired string `json:"expired"`
}
