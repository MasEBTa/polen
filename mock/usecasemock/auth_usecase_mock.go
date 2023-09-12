package usecasemock

import (
	"polen/model/dto"

	"github.com/stretchr/testify/mock"
)

type AuthUseCaseMock struct {
	mock.Mock
}

func (a *AuthUseCaseMock) Login(payload dto.AuthRequest) (dto.AuthResponse, error) {
	args := a.Called(payload)
	if args.Get(1) != nil {
		return dto.AuthResponse{}, args.Error(1)
	}
	return args.Get(0).(dto.AuthResponse), nil
}
