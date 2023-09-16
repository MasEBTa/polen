package repository

import (
	"database/sql"
	"errors"
	"polen/mock"
	"polen/model/dto"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TopUpRepoTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSQL sqlmock.Sqlmock
	repo    TopUp
}

func (t *TopUpRepoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(t.T(), err)
	t.mockDB = db
	t.mockSQL = mock
	t.repo = NewTopUpRepository(t.mockDB)
}
func TestTopUpRepoTestSuite(t *testing.T) {
	suite.Run(t, new(TopUpRepoTestSuite))
}
func (t *TopUpRepoTestSuite) TestPagging_Success() {
	rows := sqlmock.NewRows([]string{"id", "maturity_time", "top_up_amount", "accepted_time", "accepted_status", "status_information", "transfer_confirmation_recipt", "recipt_file"})
	for _, row := range mock.MockListTopUp {
		rows.AddRow(row.Id, row.MaturityTime, row.TopUpAmount, row.AcceptedTime, row.Accepted, row.Status, row.TransferConfirmRecipe, row.File)
	}
	expectedSQL := `SELECT id, maturity_time, top_up_amount, accepted_time, accepted_status, status_information, transfer_confirmation_recipt, recipt_file FROM top_up LIMIT $2 OFFSET $1;`
	t.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(((mock.MockPageReq.Page - 1) * mock.MockPageReq.Size), mock.MockPageReq.Size).WillReturnRows(rows)

	rowCount := sqlmock.NewRows([]string{"count"})
	rowCount.AddRow(1)
	t.mockSQL.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(id) FROM top_up`)).WillReturnRows(rowCount)

	uc, p, err := t.repo.Pagging(mock.MockPageReq)
	assert.Nil(t.T(), err)
	assert.Equal(t.T(), 1, len(uc))
	assert.Equal(t.T(), 1, p.TotalRows)
}
func (t *TopUpRepoTestSuite) TestPagging_Fail() {
	// error select paging
	expectedSQL := `SELECT id, maturity_time, top_up_amount, accepted_time, accepted_status, status_information, transfer_confirmation_recipt, recipt_file FROM top_up LIMIT $2 OFFSET $1;`
	t.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WillReturnError(errors.New("failed"))
	uc, p, err := t.repo.Pagging(dto.PageRequest{})
	assert.Error(t.T(), err)
	assert.Nil(t.T(), uc)
	assert.Equal(t.T(), 0, p.TotalRows)
	// error select count
	rows := sqlmock.NewRows([]string{"id", "maturity_time", "top_up_amount", "accepted_time", "accepted_status", "status_information", "transfer_confirmation_recipt", "recipt_file"})
	for _, row := range mock.MockListTopUp {
		rows.AddRow(row.Id, row.MaturityTime, row.TopUpAmount, row.AcceptedTime, row.Accepted, row.Status, row.TransferConfirmRecipe, row.File)
	}
	expectedSQL = `SELECT id, maturity_time, top_up_amount, accepted_time, accepted_status, status_information, transfer_confirmation_recipt, recipt_file FROM top_up LIMIT $2 OFFSET $1;`
	t.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(((mock.MockPageReq.Page - 1) * mock.MockPageReq.Size), mock.MockPageReq.Size).WillReturnRows(rows)
	t.mockSQL.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(id) FROM top_up`)).WillReturnError(errors.New("failed"))
	uc, p, err = t.repo.Pagging(mock.MockPageReq)
	assert.Error(t.T(), err)
	assert.Nil(t.T(), uc)
	assert.Equal(t.T(), 0, p.TotalRows)
}
