package repository

import (
	"database/sql"
	"polen/model/dto"
)

type LoanRepository interface {
	create(loanReq dto.LoanRequest) error
}

type loanRepository struct {
	db *sql.DB
}

// create implements LoanRepository.
func (l *loanRepository) create(loan dto.LoanRequest) error {
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
