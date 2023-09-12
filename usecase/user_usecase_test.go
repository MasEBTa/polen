package usecase

import (
	"errors"
	"polen/mock/repomock"
	"polen/model"
	"polen/model/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	urm *repomock.UserRepoMock
	uuc UserUseCase
}

func (u *UserUseCaseTestSuite) SetupTest() {
	u.urm = new(repomock.UserRepoMock)
	u.uuc = NewUserUseCase(u.urm)
}

func TestUserUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}

func (u *UserUseCaseTestSuite) TestFindByUsername_Success() {
	mockData := model.UserCredential{
		Id:       "1",
		Username: "akbar",
		Password: "123",
		Role:     "borrower",
		IsActive: false,
	}
	u.urm.On("FindByUsername", mockData.Username).Return(mockData, nil)
	uc, err := u.uuc.FindByUsername(mockData.Username)
	assert.Nil(u.T(), err)
	assert.NoError(u.T(), err)
	assert.Equal(u.T(), mockData, uc)
}
func (u *UserUseCaseTestSuite) TestFindByUsername_Fail() {
	u.urm.On("FindByUsername", "akbar").Return(model.UserCredential{}, errors.New("error"))
	uc, err := u.uuc.FindByUsername("akbar")
	assert.Error(u.T(), err)
	assert.NotNil(u.T(), err)
	assert.Equal(u.T(), model.UserCredential{}, uc)
}
func (u *UserUseCaseTestSuite) TestRegister_RoleRequired() {
	u.urm.On("Save", model.UserCredential{
		Username: "akbar",
		Password: "123",
	}).Return(errors.New("role is required"))
	err := u.uuc.Register(dto.AuthRequest{Username: "akbar", Password: "123"})
	assert.Error(u.T(), err)
	assert.NotNil(u.T(), err)
}
func (u *UserUseCaseTestSuite) TestRegister_UsernameRequired() {
	u.urm.On("Save", model.UserCredential{
		Password: "123",
	}).Return(errors.New("username required"))
	err := u.uuc.Register(dto.AuthRequest{Password: "123"})
	assert.Error(u.T(), err)
	assert.NotNil(u.T(), err)
}
func (u *UserUseCaseTestSuite) TestRegister_PasswordRequired() {
	u.urm.On("Save", model.UserCredential{
		Username: "akbar",
		Role:     "peminjam",
	}).Return(errors.New("password required"))
	err := u.uuc.Register(dto.AuthRequest{
		Username: "akbar",
		Role:     "peminjam",
	})
	assert.Error(u.T(), err)
	assert.NotNil(u.T(), err)
}
func (u *UserUseCaseTestSuite) TestRegister_ChooseRoleInvalid() {
	u.urm.On("Save", model.UserCredential{
		Username: "akbar",
		Role:     "borrower",
	}).Return(errors.New("role you has choose isnt available"))
	err := u.uuc.Register(dto.AuthRequest{
		Username: "akbar",
		Password: "123",
		Role:     "borrower",
	})
	assert.Error(u.T(), err)
	assert.NotNil(u.T(), err)
}
func (u *UserUseCaseTestSuite) TestFindById_Success() {
	mockData := model.UserCredential{
		Id:       "1",
		Username: "akbar",
		Password: "123",
		Role:     "borrower",
		IsActive: false,
	}
	u.urm.On("FindById", mockData.Id).Return(mockData, nil)
	uc, err := u.uuc.FindById(mockData.Id)
	assert.Nil(u.T(), err)
	assert.NoError(u.T(), err)
	assert.Equal(u.T(), mockData, uc)
}
func (u *UserUseCaseTestSuite) TestFindById_Fail() {
	u.urm.On("FindById", "1").Return(model.UserCredential{}, errors.New("error"))
	uc, err := u.uuc.FindById("1")
	assert.Error(u.T(), err)
	assert.NotNil(u.T(), err)
	assert.Equal(u.T(), model.UserCredential{}, uc)
}
