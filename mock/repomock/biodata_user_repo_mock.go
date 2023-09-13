package repomock

import (
	"polen/model"

	"github.com/stretchr/testify/mock"
)

type BiodataUserRepoMock struct {
	mock.Mock
}

// FindById implements BiodataUser.
func (b *BiodataUserRepoMock) FindById(id string) (model.BiodataUser, error) {
	args := b.Called(id)
	if args.Get(1) != nil {
		return model.BiodataUser{}, args.Error(1)
	}
	return args.Get(0).(model.BiodataUser), nil
}

// DeleteById implements BiodataUser.
func (b *BiodataUserRepoMock) DeleteById(id string) error {
	return b.Called(id).Error(0)
}

// FindAll implements BiodataUser.
func (b *BiodataUserRepoMock) FindAll() ([]model.BiodataUser, error) {
	args := b.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.BiodataUser), nil
}

// FindByNIK implements BiodataUser.
func (b *BiodataUserRepoMock) FindByNIK(nik string) (model.BiodataUser, error) {
	args := b.Called(nik)
	if args.Get(1) != nil {
		return model.BiodataUser{}, args.Error(1)
	}
	return args.Get(0).(model.BiodataUser), nil
}

// Save implements BiodataUser.
func (b *BiodataUserRepoMock) Save(payload model.BiodataUser) error {
	return b.Called(payload).Error(0)
}

// Update implements BiodataUser.
func (b *BiodataUserRepoMock) Update(payload model.BiodataUser) error {
	return b.Called(payload).Error(0)
}
