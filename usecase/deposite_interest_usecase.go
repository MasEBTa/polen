package usecase

import (
	"fmt"
	"polen/model"
	"polen/repository"
)

type DepositeInterestUseCase interface {
	CreateNew(payload model.DepositeInterest) error
	FindById(id string) (model.DepositeInterest, error)
	Update(payload model.DepositeInterest) error
	DeleteById(id string) error
}

type depositeInterestUseCase struct {
	repo repository.DepositeInterest
}

// DeleteById implements DepositeInterestUseCase.
func (d *depositeInterestUseCase) DeleteById(id string) error {
	deposite, err := d.repo.FindById(id)
	if err != nil {
		return err
	}

	err = d.repo.DeleteById(deposite.Id)
	if err != nil {
		return fmt.Errorf("failed to delete deposite: %v", err)
	}

	return nil
}

// FindById implements DepositeInterestUseCase.
func (d *depositeInterestUseCase) FindById(id string) (model.DepositeInterest, error) {
	deposite, err := d.repo.FindById(id)
	if err != nil {
		return model.DepositeInterest{}, err
	}
	return deposite, nil
}

// Update implements DepositeInterestUseCase.
func (d *depositeInterestUseCase) Update(payload model.DepositeInterest) error {
	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	_, err := d.FindById(payload.Id)
	if err != nil {
		return err
	}

	err = d.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update deposite: %v", err)
	}

	return nil
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
