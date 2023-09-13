package usecasemock

import (
	"polen/model"

	"github.com/stretchr/testify/mock"
)

type BiodataUserUseCaseMock struct {
	mock.Mock
}

// FindById implements BiodataUserUseCase.
func (b *BiodataUserUseCaseMock) FindById(id string) (model.BiodataUser, error) {
	args := b.Called(id)
	if args.Get(1) != nil {
		return model.BiodataUser{}, args.Error(1)
	}
	return args.Get(0).(model.BiodataUser), nil
}

// CreateNew implements BiodataUserUseCaseMock.
func (b *BiodataUserUseCaseMock) CreateNew(payload model.BiodataUser) error {
	return b.Called(payload).Error(0)
}

// Delete implements BiodataUserUseCaseMock.
func (b *BiodataUserUseCaseMock) Delete(id string) error {
	return b.Called(id).Error(0)
}

// FindAll implements BiodataUserUseCaseMock.
func (b *BiodataUserUseCaseMock) FindAll() ([]model.BiodataUser, error) {
	args := b.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.BiodataUser), nil
}

// FindByNIK implements BiodataUserUseCaseMock.
func (b *BiodataUserUseCaseMock) FindByNIK(nik string) (model.BiodataUser, error) {
	args := b.Called(nik)
	if args.Get(1) != nil {
		return model.BiodataUser{}, args.Error(1)
	}
	return args.Get(0).(model.BiodataUser), nil
}

// Update implements BiodataUserUseCaseMock.
func (b *BiodataUserUseCaseMock) Update(payload model.BiodataUser) error {
	return b.Called(payload).Error(0)
}
