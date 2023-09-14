package model

import "time"

type Loan struct {
	Id                     string
	UserCredentialId       string
	LoanDateCreate         time.Time
	LoanAmount             int
	LoanInterestRate       float64
	AppHandlingCostNominal int
	AppHandlingCostUnit    string
	TotalAmountOfDepth     int
	Status                 bool
}

type InstallenmentLoan struct {
	Id                 string
	UserCredId         string
	LoanId             string
	IsPayed            bool
	PaymentInstallment int
	PaymentDeadLine    time.Time
	LatePaymentFees    int
	LatePaymentDate    time.Time
}
