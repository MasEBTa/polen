package usecase

import (
	"fmt"
	"polen/model"
	"polen/model/dto"
	"polen/repository"
	"polen/utils/common"
	"polen/utils/security"
)

type UserUseCase interface {
	FindByUsername(username string) (model.UserCredential, error)
	Register(payload dto.AuthRequest) error
}

type userUseCase struct {
	repo repository.UserRepository
}

// FindByUsername implements UserUseCase.
func (u *userUseCase) FindByUsername(username string) (model.UserCredential, error) {
	return u.repo.FindByUsername(username)
}

// Register implements UserUseCase.
func (u *userUseCase) Register(payload dto.AuthRequest) error {
	hashPassword, err := security.HashPassword(payload.Password)
	if err != nil {
		return err
	}

	userCredential := model.UserCredential{
		Id:       common.GenerateID(),
		Username: payload.Username,
		Password: hashPassword,
	}

	err = u.repo.Save(userCredential)
	if err != nil {
		return fmt.Errorf("failed save user: %v", err.Error())
	}
	return nil
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}
