package dto

type AuthRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type AuthResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type GetAuthResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	VaNumber string `json:"va number"`
	IsActive bool   `json:"activated"`
}
