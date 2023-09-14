package repository

import (
	"database/sql"
	"polen/model"
)

type DepositeInterest interface {
	Save(payload model.DepositeInterest) error
	Update(payload model.DepositeInterest) error
	DeleteById(id string) error
	FindById(id string) (model.DepositeInterest, error)
}

type depositeInterest struct {
	db *sql.DB
}

// DeleteById implements DepositeInterest.
func (d *depositeInterest) DeleteById(id string) error {
	_, err := d.db.Exec("DELETE FROM deposite_interest WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// FindById implements DepositeInterest.
func (d *depositeInterest) FindById(id string) (model.DepositeInterest, error) {
	row := d.db.QueryRow(`SELECT id, created_date, interest_date, duration WHERE  =$1`, id)
	deposite := model.DepositeInterest{}
	err := row.Scan(
		&deposite.Id,
		&deposite.CreateDate,
		&deposite.InterestDate,
		&deposite.Duration,
	)
	if err != nil {
		return model.DepositeInterest{}, err
	}
	return deposite, nil
}

// Update implements DepositeInterest.
func (d *depositeInterest) Update(payload model.DepositeInterest) error {
	_, err := d.db.Exec(`UPDATE deposite_interest SET created_date = $2, interest_date = $3, duration= $4  WHERE id = $1`,
		payload.Id,
		payload.CreateDate,
		payload.InterestDate,
		payload.Duration,
	)
	if err != nil {
		return err
	}
	return nil
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
