package usecase

import (
	"fmt"
	"polen/model/dto"
	"polen/repository"
	"polen/utils/common"

	"github.com/gin-gonic/gin"
)

type LoanUseCase interface {
	Create(payload dto.LoanRequest) (dto.LoanResponse, error)
}

type loanUseCase struct {
	repo repository.LoanRepository
	ctx  *gin.Context
}

// Create implements LoanUseCase.
func (loan *loanUseCase) Create(payload dto.LoanRequest) (dto.LoanResponse, error) {
	role, err := common.GetRole(loan.ctx)
	if err != nil {
		return dto.LoanResponse{}, err
	}
	if role != "peminjam" {
		return dto.LoanResponse{}, fmt.Errorf("you arent allow to do this transaction")
	}
	if payload.LoanInterestRateId == "" {
		return dto.LoanResponse{}, fmt.Errorf("loan interest rate is required")
	}
	if payload.LoanHandlingCostId == "" {
		return dto.LoanResponse{}, fmt.Errorf("loan handling cost is required")
	}
	if payload.LoanAmount <= 0 {
		return dto.LoanResponse{}, fmt.Errorf("loan amount cost is must greather than zero")
	}
	return dto.LoanResponse{}, err
}

func NewLoanUseCase(repo repository.LoanRepository, ctx *gin.Context) LoanUseCase {
	return &loanUseCase{
		repo: repo,
		ctx:  ctx,
	}
}
