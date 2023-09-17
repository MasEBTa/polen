package dto

type LoanRequest struct {
	UserCredentialId   string
	LoanInterestRateId string
	LoanHandlingCostId string
	LoanAmount         int
}

type LoanResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
