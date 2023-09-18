package repository

import (
	"database/sql"
	"errors"
	"polen/mock"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DepositeRepoTestSuite struct {
	suite.Suite
	mockSQL sqlmock.Sqlmock
	mockDB  *sql.DB
	user    BiodataUser
	repo    DepositeRepository
}

func (d *DepositeRepoTestSuite) SetupTest() {
	db, sql, err := sqlmock.New()
	assert.NoError(d.T(), err)
	d.mockDB = db
	d.mockSQL = sql
	d.repo = NewDepositeRepository(d.mockDB, d.user)
}

func TestDepositeRepoTestSuite(t *testing.T) {
	suite.Run(t, new(DepositeRepoTestSuite))
}
func (d *DepositeRepoTestSuite) TestPagging_Success() {
	rows := sqlmock.NewRows([]string{"id", "deposit_amount", "interest_rate", "tax_rate", "duration", "created_date", "maturity_date", "status", "gross_profit", "tax", "net_profit", "total_return"})
	for _, deposite := range mock.MockDepositeDto {
		rows.AddRow(deposite.Id, deposite.DepositeAmount, deposite.InterestRate, deposite.TaxRate, deposite.DurationMounth, deposite.CreateDate, deposite.MaturityDate, deposite.Status, deposite.GrossProfit, deposite.Tax, deposite.NetProfit, deposite.TotalReturn)
	}
	expectedSQL := `SELECT id, deposit_amount, interest_rate, tax_rate, duration, created_date, maturity_date, status, gross_profit, tax, net_profit, total_return FROM deposit LIMIT $2 OFFSET $1;`
	d.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(((mock.MockPageReq.Page - 1) * mock.MockPageReq.Size), mock.MockPageReq.Size).WillReturnRows(rows)

	rowCount := sqlmock.NewRows([]string{"count"})
	rowCount.AddRow(1)
	d.mockSQL.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(id) FROM deposit_interest`)).WillReturnRows(rowCount)

	uc, p, err := d.repo.Pagging(mock.MockPageReq)
	assert.Nil(d.T(), err)
	assert.Equal(d.T(), 1, len(uc))
	assert.Equal(d.T(), 1, p.TotalRows)
}
func (d *DepositeRepoTestSuite) TestPagging_Failed() {
	// error select paging
	expectedSQL := `SELECT id, deposit_amount, interest_rate, tax_rate, duration, created_date, maturity_date, status, gross_profit, tax, net_profit, total_return FROM deposit LIMIT $2 OFFSET $1;`
	d.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(((mock.MockPageReq.Page - 1) * mock.MockPageReq.Size), mock.MockPageReq.Size).WillReturnError(errors.New("error"))
	uc, p, err := d.repo.Pagging(mock.MockPageReq)
	assert.Error(d.T(), err)
	assert.Nil(d.T(), uc)
	assert.Equal(d.T(), 0, p.TotalRows)
	// error select count
	rows := sqlmock.NewRows([]string{"id", "deposit_amount", "interest_rate", "tax_rate", "duration", "created_date", "maturity_date", "status", "gross_profit", "tax", "net_profit", "total_return"})
	for _, deposite := range mock.MockDepositeDto {
		rows.AddRow(deposite.Id, deposite.DepositeAmount, deposite.InterestRate, deposite.TaxRate, deposite.DurationMounth, deposite.CreateDate, deposite.MaturityDate, deposite.Status, deposite.GrossProfit, deposite.Tax, deposite.NetProfit, deposite.TotalReturn)
	}
	expectedSQL = `SELECT id, deposit_amount, interest_rate, tax_rate, duration, created_date, maturity_date, status, gross_profit, tax, net_profit, total_return FROM deposit LIMIT $2 OFFSET $1;`
	d.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(((mock.MockPageReq.Page - 1) * mock.MockPageReq.Size), mock.MockPageReq.Size).WillReturnRows(rows)
	rowCount := sqlmock.NewRows([]string{"count"})
	rowCount.AddRow(1)
	d.mockSQL.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(id) FROM deposit_interest`)).WillReturnError(errors.New("error"))
	uc, p, err = d.repo.Pagging(mock.MockPageReq)
	assert.Error(d.T(), err)
	assert.Nil(d.T(), uc)
	assert.Equal(d.T(), 0, p.TotalRows)
}
