package dto

type AuthRequest struct {
	Username string
	Email    string
	Password string
	Role     string
}

type AuthResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
