package model

type UserCredential struct {
	Id       string
	Username string
	Password string
	Role     string
	IsActive bool
}
