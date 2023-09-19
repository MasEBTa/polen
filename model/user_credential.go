package model

type UserCredential struct {
<<<<<<< HEAD
	Id       string
	Username string
	Email    string
	Password string
	VANumber string
	Role     string
	IsActive bool
=======
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
>>>>>>> dev/akbar
}
