package repository

import (
	"database/sql"
	"errors"
	"polen/mock"
	"polen/model"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DepositeInterestRepoTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSQL sqlmock.Sqlmock
	repo    DepositeInterest
}

func (d *DepositeInterestRepoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(d.T(), err)
	d.mockDB = db
	d.mockSQL = mock
	d.repo = NewDepositeInterestRepository(d.mockDB)
}
func TestDepositeInterestRepoTestSuite(t *testing.T) {
	suite.Run(t, new(DepositeInterestRepoTestSuite))
}
func (d *DepositeInterestRepoTestSuite) TestDeleteById_Success() {
	d.mockSQL.ExpectExec(regexp.QuoteMeta(`DELETE FROM deposite_interest WHERE id = $1`)).WithArgs(mock.MockDepositeInterest.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	err := d.repo.DeleteById(mock.MockDepositeInterest.Id)
	assert.Nil(d.T(), err)
}
func (d *DepositeInterestRepoTestSuite) TestDeleteById_Failed() {
	d.mockSQL.ExpectExec(regexp.QuoteMeta(`DELETE FROM deposite_interest WHERE id = $1`)).WithArgs(mock.MockDepositeInterest.Id).WillReturnError(errors.New("failed delete"))
	err := d.repo.DeleteById(mock.MockDepositeInterest.Id)
	assert.Error(d.T(), err)
}
func (d *DepositeInterestRepoTestSuite) TestFindById_Success() {
	rows := sqlmock.NewRows([]string{"id", "created_date", "interest_date", "duration"})
	rows.AddRow(mock.MockDepositeInterest.Id, mock.MockDepositeInterest.CreateDate, mock.MockDepositeInterest.InterestDate, mock.MockDepositeInterest.Duration)
	d.mockSQL.ExpectQuery(regexp.QuoteMeta(`SELECT id, created_date, interest_date, duration WHERE  =$1`)).WithArgs(mock.MockDepositeInterest.Id).WillReturnRows(rows)
	di, err := d.repo.FindById(mock.MockDepositeInterest.Id)
	assert.Nil(d.T(), err)
	assert.Equal(d.T(), mock.MockDepositeInterest.Id, di.Id)
}
func (d *DepositeInterestRepoTestSuite) TestFindById_Failed() {
	d.mockSQL.ExpectQuery(regexp.QuoteMeta(`SELECT id, created_date, interest_date, duration WHERE  =$1`)).WithArgs(mock.MockDepositeInterest.Id).WillReturnError(errors.New("failed find "))
	di, err := d.repo.FindById(mock.MockDepositeInterest.Id)
	assert.Error(d.T(), err)
	assert.Equal(d.T(), model.DepositeInterest{}, di)
}
func (d *DepositeInterestRepoTestSuite) TestUpdate_Success() {
	d.mockSQL.ExpectExec(regexp.QuoteMeta(`UPDATE deposite_interest SET created_date = $2, interest_date = $3, duration= $4  WHERE id = $1`)).WithArgs(mock.MockDepositeInterest.Id, mock.MockDepositeInterest.CreateDate, mock.MockDepositeInterest.InterestDate, mock.MockDepositeInterest.Duration).WillReturnResult(sqlmock.NewResult(1, 1))
	err := d.repo.Update(mock.MockDepositeInterest)
	assert.Nil(d.T(), err)
}
func (d *DepositeInterestRepoTestSuite) TestUpdate_Failed() {
	d.mockSQL.ExpectExec(regexp.QuoteMeta(`UPDATE deposite_interest SET created_date = $2, interest_date = $3, duration= $4  WHERE id = $1`)).WithArgs(mock.MockDepositeInterest.Id, mock.MockDepositeInterest.CreateDate, mock.MockDepositeInterest.InterestDate, mock.MockDepositeInterest.Duration).WillReturnError(errors.New("failed update"))
	err := d.repo.Update(mock.MockDepositeInterest)
	assert.Error(d.T(), err)
}
func (d *DepositeInterestRepoTestSuite) TestSave_Success() {
	d.mockSQL.ExpectExec(`INSERT INTO deposite_interest`).WithArgs(mock.MockDepositeInterest.Id, mock.MockDepositeInterest.CreateDate, mock.MockDepositeInterest.InterestDate, mock.MockDepositeInterest.Duration).WillReturnResult(sqlmock.NewResult(1, 1))
	err := d.repo.Save(mock.MockDepositeInterest)
	assert.Nil(d.T(), err)
	assert.NoError(d.T(), err)
}
func (d *DepositeInterestRepoTestSuite) TestSave_Failed() {
	d.mockSQL.ExpectExec(`INSERT INTO deposite_interest`).WithArgs(mock.MockDepositeInterest.Id, mock.MockDepositeInterest.CreateDate, mock.MockDepositeInterest.InterestDate, mock.MockDepositeInterest.Duration).WillReturnError(errors.New("failed save"))
	err := d.repo.Save(mock.MockDepositeInterest)
	assert.Error(d.T(), err)
	assert.NotNil(d.T(), err)
}
