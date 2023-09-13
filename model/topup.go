package model

type TopUp struct {
	Id             string
	UserCredential UserCredential
	TopUpAmount    int
	MaturityTime   string
	Status         string
}
