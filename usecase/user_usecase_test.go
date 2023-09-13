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

func (u *UserUseCaseTestSuite) TestRegister_EmptyInvalid() {
	// username required
	u.urm.On("Save", model.UserCredential{Id: "1"}).Return(errors.New("username required"))
	err := u.uuc.Register(dto.AuthRequest{})
	assert.Error(u.T(), err)
	// password required
	u.urm.On("Save", model.UserCredential{Id: "1", Username: "akbar", Email: "akbar@gmail.com"}).Return(errors.New("password required"))
	err = u.uuc.Register(dto.AuthRequest{Username: "akbar", Email: "akbar@gmail.com"})
	assert.Error(u.T(), err)
	// email required
	u.urm.On("Save", model.UserCredential{Id: "1", Username: "akbar", Password: "123"}).Return(errors.New("email required"))
	err = u.uuc.Register(dto.AuthRequest{Username: "akbar", Password: "123"})
	assert.Error(u.T(), err)
	// role required
	u.urm.On("Save", model.UserCredential{Id: "1", Username: "akbar", Email: "akbar@gmail.com", Password: "123"}).Return(errors.New("role is required"))
	err = u.uuc.Register(dto.AuthRequest{Username: "akbar", Email: "akbar@gmail.com", Password: "123"})
	assert.Error(u.T(), err)
}
func (u *UserUseCaseTestSuite) TestRegisterCheckRole_Fail() {
	mockUserCred := model.UserCredential{
		Id:       "1",
		Username: "akbar",
		Email:    "akbar@gmail.com",
		Password: "123",
		VANumber: "dhfbdsfds123",
		Role:     "borrower",
		IsActive: false,
	}
	mockAuthReq := dto.AuthRequest{
		Username: "akbar",
		Email:    "akbar@gmail.com",
		Password: "123",
		Role:     "borrower",
	}
	u.urm.On("Save", mockUserCred).Return(errors.New("role you has choose isnt available"))
	err := u.uuc.Register(mockAuthReq)
	assert.Error(u.T(), err)
}
func (u *UserUseCaseTestSuite) TestRegister_InvalidEmail() {
	mockUserCred := model.UserCredential{
		Id:       "1",
		Username: "akbar",
		Email:    "akbar@gmail",
		Password: "123",
		VANumber: "dhfbdsfds123",
		Role:     "peminjam",
		IsActive: false,
	}
	mockAuthReq := dto.AuthRequest{
		Username: "akbar",
		Email:    "akbar@gmail",
		Password: "123",
		Role:     "peminjam",
	}
	u.urm.On("Save", mockUserCred).Return(errors.New("is not valid email"))
	err := u.uuc.Register(mockAuthReq)
	assert.Error(u.T(), err)
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
