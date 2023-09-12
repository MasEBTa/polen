package repository

import (
	"database/sql"
	"errors"
	"polen/model"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRepoTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSQL sqlmock.Sqlmock
	repo    UserRepository
}

func (u *UserRepoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(u.T(), err)
	u.mockDB = db
	u.mockSQL = mock
	u.repo = NewUserRepository(u.mockDB)
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

func (u *UserRepoTestSuite) TestFindByUsername_Success() {
	mockData := model.UserCredential{
		Id:       "1",
		Username: "akbar",
		Password: "123",
		Role:     "borrower",
	}
	rows := sqlmock.NewRows([]string{"id", "username", "role", "password"})
	rows.AddRow(mockData.Id, mockData.Username, mockData.Role, mockData.Password)
	expectedSQL := `SELECT id, username, role, password FROM user_credential WHERE username = $1`
	u.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(mockData.Username).WillReturnRows(rows)
	uc, err := u.repo.FindByUsername(mockData.Username)
	assert.Nil(u.T(), err)
	assert.NoError(u.T(), err)
	assert.Equal(u.T(), mockData, uc)
}
func (u *UserRepoTestSuite) TestFindByUsername_Fail() {
	expectedSQL := `SELECT id, username, role, password FROM user_credential WHERE username = $1`
	u.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs("ismail").WillReturnError(errors.New("error"))
	uc, err := u.repo.FindByUsername("ismail")
	assert.Error(u.T(), err)
	assert.NotNil(u.T(), err)
	assert.Equal(u.T(), model.UserCredential{}, uc)
}

func (u *UserRepoTestSuite) TestSave_Success() {
	mockData := model.UserCredential{
		Id:       "1",
		Username: "akbar",
		Password: "123",
		Role:     "borrower",
		IsActive: true,
	}
	expectedSQL := `INSERT INTO user_credential`
	u.mockSQL.ExpectExec(expectedSQL).WithArgs(mockData.Id, mockData.Username, mockData.Password, mockData.Role, mockData.IsActive).WillReturnResult(sqlmock.NewResult(1, 1))
	err := u.repo.Save(mockData)
	assert.Nil(u.T(), err)
	assert.NoError(u.T(), err)
}
func (u *UserRepoTestSuite) TestSave_Failed() {
	mockData := model.UserCredential{
		Id:       "1",
		Username: "akbar",
		Password: "123",
		Role:     "borrower",
		IsActive: true,
	}
	expectedSQL := `INSERT INTO user_credential`
	u.mockSQL.ExpectExec(expectedSQL).WithArgs(mockData.Id, mockData.Username, mockData.Password, mockData.Role, mockData.IsActive).WillReturnError(errors.New("error"))
	err := u.repo.Save(mockData)
	assert.Error(u.T(), err)
	assert.NotNil(u.T(), err)
}
func (u *UserRepoTestSuite) TestFindById_Success() {
	mockData := model.UserCredential{
		Id:       "1",
		Username: "akbar",
		Password: "123",
		Role:     "borrower",
	}
	rows := sqlmock.NewRows([]string{"id", "username", "role", "password"})
	rows.AddRow(mockData.Id, mockData.Username, mockData.Role, mockData.Password)
	expectedSQL := `SELECT id, username, role, password FROM user_credential WHERE id =$1`
	u.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(mockData.Id).WillReturnRows(rows)
	uc, err := u.repo.FindById(mockData.Id)
	assert.Nil(u.T(), err)
	assert.NoError(u.T(), err)
	assert.Equal(u.T(), mockData, uc)
}
func (u *UserRepoTestSuite) TestFindById_Fail() {
	expectedSQL := `SELECT id, username, role, password FROM user_credential WHERE id =$1`
	u.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs("ismail").WillReturnError(errors.New("error"))
	uc, err := u.repo.FindById("ismail")
	assert.Error(u.T(), err)
	assert.NotNil(u.T(), err)
	assert.Equal(u.T(), model.UserCredential{}, uc)
}
