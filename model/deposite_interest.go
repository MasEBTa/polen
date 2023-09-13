package model

import "time"

type DepositeInterest struct {
	Id           string
	CreateDate   time.Time
	InterestDate time.Time
	Duration     time.Time
}
