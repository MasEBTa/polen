package repository

import (
	"database/sql"
	"polen/model"
)

type UserRepository interface {
	Save(payload model.UserCredential) error
	Saldo(payload model.UserCredential, idsaldo string) error
	FindByUsername(username string) (model.UserCredential, error)

	FindById(id string) (model.UserCredential, error)
}

type userRepository struct {
	db *sql.DB
}

// saldo implements UserRepository.
func (u *userRepository) Saldo(payload model.UserCredential, idsaldo string) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO user_credential VALUES ($1, $2, $3, $4, $5, $6)", payload.Id, payload.Username, payload.Email, payload.Password, payload.Role, true)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec("INSERT INTO saldo VALUES ($1, $2, $3)", idsaldo, payload.Id, 0)
	if err != nil {
		tx.Rollback()
		return err
	}
	return err
}

// FindById implements UserRepository.
func (u *userRepository) FindById(id string) (model.UserCredential, error) {
	row := u.db.QueryRow("SELECT id, username, role, password FROM user_credential WHERE id =$1", id)
	var userCredential model.UserCredential
	err := row.Scan(&userCredential.Id, &userCredential.Username, &userCredential.Role, &userCredential.Password)
	if err != nil {
		return model.UserCredential{}, err
	}
	return userCredential, nil
}

// FindByUsername implements UserRepository.
func (u *userRepository) FindByUsername(username string) (model.UserCredential, error) {
	row := u.db.QueryRow("SELECT id, username, role, password FROM user_credential WHERE username = $1", username)
	var userCredential model.UserCredential
	err := row.Scan(&userCredential.Id, &userCredential.Username, &userCredential.Role, &userCredential.Password)
	if err != nil {
		return model.UserCredential{}, err
	}
	return userCredential, nil
}

// Save implements UserRepository.
func (u *userRepository) Save(payload model.UserCredential) error {
	_, err := u.db.Exec("INSERT INTO user_credential VALUES ($1, $2, $3, $4, $5, $6)", payload.Id, payload.Username, payload.Email, payload.Password, payload.Role, true)
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
