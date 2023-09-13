package usecase

import (
	"fmt"
	"polen/model"
	"polen/repository"
)

type TopUpUseCase interface {
	CreateNew(payload model.TopUp) error
	FindById(id string) (model.TopUp, error)
}

type topUpUseCase struct {
	repo   repository.TopUp
	userUC UserUseCase
}

// CreateNew implements TopUpUseCase.
func (t *topUpUseCase) CreateNew(payload model.TopUp) error {
	if payload.TopUpAmount <= 0 {
		return fmt.Errorf("Top Up must be greater than zero")
	}

	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	_, err := t.userUC.FindById(payload.UserCredential.Id)
	if err != nil {
		return err
	}

	err = t.repo.Save(payload)
	if err != nil {
		return fmt.Errorf("failed to save new topup: %v", err)
	}

	return nil
}

// FindById implements TopUpUseCase.
func (t *topUpUseCase) FindById(id string) (model.TopUp, error) {
	topup, err := t.repo.FindById(id)
	if err != nil {
		return model.TopUp{}, err
	}
	return topup, nil
}

func NewTopUpUseCase(repo repository.TopUp, userUC UserUseCase) TopUpUseCase {
	return &topUpUseCase{
		repo:   repo,
		userUC: userUC,
	}
}
