package usecase

import (
	"errors"
	"polen/mock/repomock"
	"polen/model"
	"polen/model/dto"
	"polen/utils/security"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthUseCaseTestSuite struct {
	suite.Suite
	urm *repomock.UserRepoMock
	auc AuthUseCase
}

func (a *AuthUseCaseTestSuite) SetupTest() {
	a.urm = new(repomock.UserRepoMock)
	a.auc = NewAuthUseCase(a.urm)
}

func TestAuthUseCaseSuite(t *testing.T) {
	suite.Run(t, new(AuthUseCaseTestSuite))
}

func (a *AuthUseCaseTestSuite) TestLogin_UsernameInvalid() {
	a.urm.On("FindByUsername", "akbr").Return(model.UserCredential{}, errors.New("unauthorized: invalid credential"))
	ar, err := a.auc.Login(dto.AuthLoginRequest{Username: "akbr"})
	assert.Error(a.T(), err)
	assert.NotNil(a.T(), err)
	assert.Equal(a.T(), dto.AuthResponse{}, ar)
}
func (a *AuthUseCaseTestSuite) TestLogin_PasswordInvalid() {
	mockUserCredential := model.UserCredential{
		Username: "akbar",
		Password: "123",
		Role:     "peminjam",
	}
	mockAuthRequest := dto.AuthLoginRequest{
		Username: "akbar",
		Password: "123",
	}
	a.urm.On("FindByUsername", mockUserCredential.Username).Return(mockUserCredential, nil)
	err := security.VerifyPassword(mockUserCredential.Password, "123")
	assert.Error(a.T(), err)
	ar, err2 := a.auc.Login(mockAuthRequest)
	assert.Error(a.T(), err2)
	assert.Equal(a.T(), dto.AuthResponse{}, ar)
}
