package repository

import (
	"database/sql"
	"polen/model"
)

type TopUp interface {
	Save(payload model.TopUp) error
	FindById(id string) (model.TopUp, error)
}

type topUpRepository struct {
	db *sql.DB
}

// FindById implements TopUp.
func (t *topUpRepository) FindById(id string) (model.TopUp, error) {
	row := t.db.QueryRow(`SELECT t.id, u.user, u.username,u.email, u.password, u.role, u.virtual_account_number, u.is_active, 
	t.top_up_amount, t.countdown_time, t.status FROM top_up t JOIN user_credential u ON u.id = t.user_credential_id WHERE t.id =$1`, id)
	topup := model.TopUp{}
	err := row.Scan(
		&topup.Id,
		&topup.UserCredential.Id,
		&topup.UserCredential.Username,
		&topup.UserCredential.Password,
		&topup.UserCredential.Email,
		&topup.UserCredential.Role,
		&topup.UserCredential.VANumber,
		&topup.UserCredential.IsActive,
		&topup.TopUpAmount,
		&topup.MaturityTime,
		&topup.Status,
	)
	if err != nil {
		return model.TopUp{}, err
	}
	return topup, nil
}

// Save implements TopUp.
func (t *topUpRepository) Save(payload model.TopUp) error {
	_, err := t.db.Exec("INSERT INTO biodata_user VALUES ($1, $2, $3, $4, $5)",
		payload.Id,
		payload.UserCredential.Id,
		payload.TopUpAmount,
		payload.MaturityTime,
		payload.Status,
	)
	if err != nil {
		return err
	}
	return nil
}

func NewTopUpRepository(db *sql.DB) TopUp {
	return &topUpRepository{
		db: db,
	}
}
