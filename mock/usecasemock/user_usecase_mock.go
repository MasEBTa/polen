package usecasemock

import (
	"polen/model"
	"polen/model/dto"

	"github.com/stretchr/testify/mock"
)

type UserUseCaseMock struct {
	mock.Mock
}

// FindById implements usecase.UserUseCase.
func (u *UserUseCaseMock) FindById(id string) (model.UserCredential, error) {
	args := u.Called(id)
	if args.Get(1) != nil {
		return model.UserCredential{}, args.Error(1)
	}
	return args.Get(0).(model.UserCredential), nil
}

func (u *UserUseCaseMock) FindByUsername(username string) (model.UserCredential, error) {
	args := u.Called(username)
	if args.Get(1) != nil {
		return model.UserCredential{}, args.Error(1)
	}
	return args.Get(0).(model.UserCredential), nil
}
func (u *UserUseCaseMock) Register(payload dto.AuthRequest) error {
	return u.Called(payload).Error(0)
}
