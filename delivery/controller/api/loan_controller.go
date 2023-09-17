package api

import (
	"polen/delivery/middleware"
	"polen/model/dto"
	"polen/usecase"

	"github.com/gin-gonic/gin"
)

type LoanController struct {
	loanUC usecase.LoanUseCase
	rg     *gin.RouterGroup
}

func (l *LoanController) createHandler(c *gin.Context) {
	var payload dto.LoanRequest
	l.loanUC.Create(payload)
}

func (l *LoanController) Route() {
	l.rg.POST("/loan", middleware.AuthMiddleware(), l.createHandler)
}

func NewLoanController(loanUC usecase.LoanUseCase, rg *gin.RouterGroup) *LoanController {
	return &LoanController{
		loanUC: loanUC,
		rg:     rg,
	}
}
