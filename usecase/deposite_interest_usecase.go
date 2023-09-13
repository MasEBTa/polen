package usecase

import (
	"fmt"
	"polen/model"
	"polen/repository"
)

type DepositeInterestUseCase interface {
	CreateNew(payload model.DepositeInterest) error
}

type depositeInterestUseCase struct {
	repo repository.DepositeInterest
}

// CreateNew implements DepositeInteresetUseCase.
func (d *depositeInterestUseCase) CreateNew(payload model.DepositeInterest) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	if err := d.repo.Save(payload); err != nil {
		return fmt.Errorf("failed save Deposite Interest: %v", err.Error())
	}
	return nil
}

func NewDepositeInterestUseCase(repo repository.DepositeInterest) DepositeInterestUseCase {
	return &depositeInterestUseCase{
		repo: repo,
	}
}
