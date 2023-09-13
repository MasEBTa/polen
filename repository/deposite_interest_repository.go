package repository

import (
	"database/sql"
	"polen/model"
)

type DepositeInterest interface {
	Save(payload model.DepositeInterest) error
}

type depositeInterest struct {
	db *sql.DB
}

// Save implements DepositeIntereset.
func (d *depositeInterest) Save(payload model.DepositeInterest) error {
	_, err := d.db.Exec("INSERT INTO deposite_interest VALUES ($1, $2, $3, $4)", payload.Id, payload.CreateDate, payload.InterestDate, payload.Duration)
	if err != nil {
		return err
	}
	return nil
}

func NewDepositeInterestRepository(db *sql.DB) DepositeInterest {
	return &depositeInterest{
		db: db,
	}
}
