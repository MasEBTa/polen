package usecase

import (
	"polen/model/dto"
	"polen/repository"
)

type LoanUseCase interface {
	Create(payload dto.LoanRequest) (int, dto.LoanResponse, error)
}

type loanUseCase struct {
	repo   repository.LoanRepository
	loanIr LoanInterestUseCase
	loanHc AppHandlingCostUsecase
}

// Create implements LoanUseCase.
func (loan *loanUseCase) Create(payload dto.LoanRequest) (int, dto.LoanResponse, error) {
	// if payload.LoanInterestRateId == "" {
	// 	return 400, dto.LoanResponse{}, fmt.Errorf("loan interest rate is required")
	// }
	// if payload.LoanHandlingCostId == "" {
	// 	return 400, dto.LoanResponse{}, fmt.Errorf("loan handling cost is required")
	// }
	// if payload.LoanAmount <= 0 {
	// 	return 400, dto.LoanResponse{}, fmt.Errorf("loan amount cost is must greather than zero")
	// }
	// loanIr, err := loan.loanIr.FindById(payload.LoanInterestRateId)
	// if err != nil {
	// 	return 404, dto.LoanResponse{}, fmt.Errorf("loan duration you choose arent available,")
	// }
	// loanHc, err := loan.loanHc.FindById(payload.LoanHandlingCostId)
	// if err != nil {
	// 	return 404, dto.LoanResponse{}, fmt.Errorf("handling cost you choose arent available,")
	// }
	// // building payload
	// var loanpayload model.Loan
	// loanpayload.Id = common.GenerateID()
	// loanpayload.UserCredentialId = payload.UserCredentialId
	// loanpayload.LoanDateCreate = time.Now()
	// loanpayload.LoanAmount = payload.LoanAmount
	// loanpayload.LoanInterestRate = loanIr.LoanInterestRate
	// loanpayload.AppHandlingCostUnit = loanHc.Unit
	// // app handling cost
	// if loanpayload.AppHandlingCostUnit == "rupiah" {
	// 	loanpayload.AppHandlingCostNominal = loanHc.Nominal
	// } else if loanpayload.AppHandlingCostUnit == "percent" {
	// 	loanpayload.AppHandlingCostNominal = payload.LoanAmount * loanHc.Nominal
	// }
	// // loan rate
	// loanpayload.LoanInterestNominal = payload.LoanAmount + int(float64(payload.LoanAmount)*loanIr.LoanInterestRate)
	// loanpayload.TotalAmountOfDepth = loanpayload.LoanAmount + loanpayload.LoanInterestNominal + loanpayload.AppHandlingCostNominal
	// loanpayload.Status = true

	// // var instalenmentpayload []model.InstallenmentLoan
	// type paymentInstallment struct {
	// 	paymentDeadLine time.Time
	// 	total           int
	// }
	// var pis []paymentInstallment

	// for i := 1; i <= loanIr.DurationMonths; i++ {
	// 	pi := paymentInstallment{
	// 		paymentDeadLine: time.Now().AddDate(0, i, 0),
	// 		total:           loanpayload.TotalAmountOfDepth / loanIr.DurationMonths,
	// 	}
	// 	pis = append(pis, pi)
	// }

	// for _, v := range pis {
	// 	// build payload installentment
	// 	var instalenment model.InstallenmentLoan
	// 	instalenment.Id = common.GenerateID()
	// 	instalenment.UserCredId = payload.UserCredentialId
	// 	instalenment.LoanId = loanpayload.Id
	// 	instalenment.IsPayed = false
	// 	instalenment.PaymentInstallment = v.total
	// 	instalenment.PaymentDeadLine = v.paymentDeadLine
	// }

	return 200, dto.LoanResponse{}, nil
}

func NewLoanUseCase(repo repository.LoanRepository, loanir LoanInterestUseCase, loanhc AppHandlingCostUsecase) LoanUseCase {
	return &loanUseCase{
		repo:   repo,
		loanIr: loanir,
		loanHc: loanhc,
	}
}
