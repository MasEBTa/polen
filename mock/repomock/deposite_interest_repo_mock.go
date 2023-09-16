package repomock

import (
	"polen/model"

	"github.com/stretchr/testify/mock"
)

type DepositeInterestRepoMock struct {
	mock.Mock
}

// DeleteById implements DepositeInterest.
func (d *DepositeInterestRepoMock) DeleteById(id string) error {
	return d.Called(id).Error(0)
}

// FindById implements DepositeInterest.
func (d *DepositeInterestRepoMock) FindById(id string) (model.DepositeInterest, error) {
	args := d.Called(id)
	if args.Get(1) != nil {
		return model.DepositeInterest{}, args.Error(1)
	}
	return args.Get(0).(model.DepositeInterest), nil
}

// Update implements DepositeInterest.
func (d *DepositeInterestRepoMock) Update(payload model.DepositeInterest) error {
	return d.Called(payload).Error(0)
}

// Save implements DepositeIntereset.
func (d *DepositeInterestRepoMock) Save(payload model.DepositeInterest) error {
	return d.Called(payload).Error(0)
}
