package model

import "time"

type Loan struct {
	Id                     string
	UserCredentialId       string
	LoanDateCreate         time.Time
	LoanAmount             int
	LoanInterestRate       float64
	LoanInterestNominal    int
	AppHandlingCostNominal int
	AppHandlingCostUnit    string
	TotalAmountOfDepth     int
	Status                 bool
}

type InstallenmentLoan struct {
	Id                     string
	UserCredId             string
	LoanId                 string
	IsPayed                bool
	PaymentInstallment     int
	PaymentDeadLine        time.Time
	AppHandlingCostNominal int
	AppHandlingCostUnit    string
	TotalAmountOfDepth     int
	LoanDateCreate         time.Time
	LatePaymentFees        int
	PaymentDate            time.Time
	status                 string
}
