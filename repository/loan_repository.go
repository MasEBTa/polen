package repository

import (
	"database/sql"
	"polen/model"
)

type LoanRepository interface {
	Create(loanReq model.Loan) error
}

type loanRepository struct {
	db *sql.DB
}

// create implements LoanRepository.
func (l *loanRepository) Create(loan model.Loan) error {
	_, err := l.db.Begin()
	if err != nil {
		return err
	}
	return nil
}

func NewLoanRepository(db *sql.DB) LoanRepository {
	return &loanRepository{
		db: db,
	}
}
