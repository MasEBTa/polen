package usecase

import (
	"fmt"
	"polen/model"
	"polen/repository"
	"polen/utils/common"

	"github.com/gin-gonic/gin"
)

type TopUpUseCase interface {
	CreateNew(c *gin.Context, payload model.TopUp) error
	FindById(id string) (model.TopUp, error)
	Update(c *gin.Context, payload model.TopUp) error
}

type topUpUseCase struct {
	repo   repository.TopUp
	userUC UserUseCase
}

// Update implements TopUpUseCase.
func (t *topUpUseCase) Update(c *gin.Context, payload model.TopUp) error {
	role, err := common.GetRole(c)
	if err != nil {
		return err
	}

	if role != "admin" {
		return fmt.Errorf("only admin can be update")
	}
	_, err = t.userUC.FindById(payload.UserCredential.Id)
	if err != nil {
		return err
	}

	_, err = t.FindById(payload.Id)
	if err != nil {
		return err
	}

	err = t.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to update top up: %v", err)
	}

	return nil

}

// CreateNew implements TopUpUseCase.
func (t *topUpUseCase) CreateNew(c *gin.Context, payload model.TopUp) error {
	role, err := common.GetRole(c)
	if err != nil {
		return err
	}

	if role != "pemodal" {
		return fmt.Errorf("Only pemodal users can create top-ups")
	}

	if payload.TopUpAmount <= 0 {
		return fmt.Errorf("Top Up must be greater than zero")
	}

	if payload.Id == "" {
		return fmt.Errorf("id is required")
	}

	_, err = t.userUC.FindById(payload.UserCredential.Id)
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
