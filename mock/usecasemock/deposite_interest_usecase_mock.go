package usecasemock

import (
	"polen/model"

	"github.com/stretchr/testify/mock"
)

type DepositeInterestUseCaseMock struct {
	mock.Mock
}

// DeleteById implements DepositeInterestUseCase.
func (d *DepositeInterestUseCaseMock) DeleteById(id string) error {
	return d.Called(id).Error(0)
}

// FindById implements DepositeInterestUseCase.
func (d *DepositeInterestUseCaseMock) FindById(id string) (model.DepositeInterest, error) {
	args := d.Called(id)
	if args.Get(1) != nil {
		return model.DepositeInterest{}, args.Error(1)
	}
	return args.Get(0).(model.DepositeInterest), nil
}

// Update implements DepositeInterestUseCase.
func (d *DepositeInterestUseCaseMock) Update(payload model.DepositeInterest) error {
	return d.Called(payload).Error(0)
}

// CreateNew implements DepositeInteresetUseCase.
func (d *DepositeInterestUseCaseMock) CreateNew(payload model.DepositeInterest) error {
	return d.Called(payload).Error(0)
}
