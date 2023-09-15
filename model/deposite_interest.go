package model

import "time"

type DepositeInterest struct {
	Id             string
	CreateDate     time.Time
	InterestRate   float64
	TaxRate        float64
	DurationMounth int
}

// profit
// (setoran pokok*interest rate*durationDay)/365

// tax
