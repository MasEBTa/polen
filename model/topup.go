package model

import "time"

type TopUp struct {
	Id             string
	UserCredential UserCredential
	TopUpAmount    int
	MaturityTime   time.Time
	Status         bool
}
