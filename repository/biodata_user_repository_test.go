package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BiodataUserRepoTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSQL sqlmock.Sqlmock
	repo    BiodataUser
}

func (b *BiodataUserRepoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(b.T(), err)
	b.mockDB = db
	b.mockSQL = mock
	b.repo = NewBiodataUserRepository(b.mockDB)
}
func TestBiodataUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(BiodataUserRepoTestSuite))
}
