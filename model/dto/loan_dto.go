package dto

type LoanRequest struct {
	LoanInterestRateId string
	LoanHandlingCostId string
	LoanAmount         int
}

type LoanResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
