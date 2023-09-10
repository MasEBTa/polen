package repository

import (
	"database/sql"
	"polen/model"
)

type UserRepository interface {
	Save(payload model.UserCredential) error
	FindByUsername(username string) (model.UserCredential, error)
}

type userRepository struct {
	db *sql.DB
}

// FindByUsername implements UserRepository.
func (u *userRepository) FindByUsername(username string) (model.UserCredential, error) {
	row := u.db.QueryRow("SELECT id, username, password FROM user_credential WHERE username = $1", username)
	var userCredential model.UserCredential
	err := row.Scan(&userCredential.Id, &userCredential.Username, &userCredential.Password)
	if err != nil {
		return model.UserCredential{}, err
	}
	return userCredential, nil
}

// Save implements UserRepository.
func (u *userRepository) Save(payload model.UserCredential) error {
	_, err := u.db.Exec("INSERT INTO user_credential VALUES ($1, $2, $3, $4)", payload.Id, payload.Username, payload.Password, true)
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
