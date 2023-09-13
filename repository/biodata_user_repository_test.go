package repository

import (
	"database/sql"
	"errors"
	"polen/model"
	"regexp"
	"testing"
	"time"

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
func (b *BiodataUserRepoTestSuite) TestFindById_Success() {
	mockData := model.BiodataUser{
		Id: "1",
		UserCredential: model.UserCredential{
			Id:       "1",
			Username: "akbar",
			Password: "123",
			Role:     "peminjam",
			IsActive: false,
		},
		NamaLengkap:  "akbar ismail",
		Nik:          "32010",
		NomorTelepon: "081123785743",
		Pekerjaan:    "mahasiswa",
		TempatLahir:  "jakarta",
		TanggalLahir: time.Date(2000, time.December, 1, 0, 0, 0, 0, time.UTC),
		KodePos:      "1610",
	}
	rows := sqlmock.NewRows([]string{"id", "user_credential_id", "user_credential_username", "user_credential_password", "user_credential_role", "user_credential_is_active", "nama_lengkap", "nik", "nomor_telepon", "pekerjaan", "tempat_lahir", "tanggal_lahir", "kode_pos"})
	rows.AddRow(mockData.Id, mockData.UserCredential.Id, mockData.UserCredential.Username, mockData.UserCredential.Password, mockData.UserCredential.Role, mockData.UserCredential.IsActive, mockData.NamaLengkap, mockData.Nik, mockData.NomorTelepon, mockData.Pekerjaan, mockData.TempatLahir, mockData.TanggalLahir, mockData.KodePos)
	expectedSQL := `SELECT b.id, u.user, u.username, u.password, u.role, u.is_active, 
	b.nama_lengkap, b.nik, b.nomor_telepon, b.pekerjaan, b.tempat_lahir, b.tanggal_lahir, b.kode_pos 
	FROM biodata_user b JOIN user_credential u ON u.id = b.user_credential_id WHERE b.id =$1`
	b.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(mockData.Id).WillReturnRows(rows)
	bu, err := b.repo.FindById(mockData.Id)
	assert.Nil(b.T(), err)
	assert.NoError(b.T(), err)
	assert.Equal(b.T(), mockData, bu)
}
func (b *BiodataUserRepoTestSuite) TestFindById_Fail() {
	expectedSQL := `SELECT b.id, u.user, u.username, u.password, u.role, u.is_active, 
	b.nama_lengkap, b.nik, b.nomor_telepon, b.pekerjaan, b.tempat_lahir, b.tanggal_lahir, b.kode_pos 
	FROM biodata_user b JOIN user_credential u ON u.id = b.user_credential_id WHERE b.id =$1`
	b.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs("123").WillReturnError(errors.New("error"))
	bu, err := b.repo.FindById("123")
	assert.Error(b.T(), err)
	assert.NotNil(b.T(), err)
	assert.Equal(b.T(), model.BiodataUser{}, bu)
}
func (b *BiodataUserRepoTestSuite) TestDeleteById_Success() {
	expectedSQL := `DELETE FROM biodata_user`
	b.mockSQL.ExpectExec(expectedSQL).WithArgs("123").WillReturnResult(sqlmock.NewResult(1, 1))
	err := b.repo.DeleteById("123")
	assert.Nil(b.T(), err)
	assert.NoError(b.T(), err)
}
func (b *BiodataUserRepoTestSuite) TestDeleteById_Fail() {
	expectedSQL := `DELETE FROM biodata_user`
	b.mockSQL.ExpectExec(expectedSQL).WithArgs("123").WillReturnError(errors.New("error"))
	err := b.repo.DeleteById("123")
	assert.Error(b.T(), err)
	assert.NotNil(b.T(), err)
}
func (b *BiodataUserRepoTestSuite) TestFindAll_Success() {
	mockData := []model.BiodataUser{
		{
			Id: "1",
			UserCredential: model.UserCredential{
				Id:       "1",
				Username: "akbar",
				Password: "123",
				Role:     "peminjam",
				IsActive: false,
			},
			NamaLengkap:  "akbar ismail",
			Nik:          "32010",
			NomorTelepon: "081123785743",
			Pekerjaan:    "mahasiswa",
			TempatLahir:  "jakarta",
			TanggalLahir: time.Date(2000, time.December, 1, 0, 0, 0, 0, time.UTC),
			KodePos:      "1610",
		},
	}
	rows := sqlmock.NewRows([]string{"id", "user_credential_id", "user_credential_username", "user_credential_password", "user_credential_role", "user_credential_is_active", "nama_lengkap", "nik", "nomor_telepon", "pekerjaan", "tempat_lahir", "tanggal_lahir", "kode_pos"})
	for _, data := range mockData {
		rows.AddRow(data.Id, data.UserCredential.Id, data.UserCredential.Username, data.UserCredential.Password, data.UserCredential.Role, data.UserCredential.IsActive, data.NamaLengkap, data.Nik, data.NomorTelepon, data.Pekerjaan, data.TempatLahir, data.TanggalLahir, data.KodePos)
	}
	expectedSQL := `SELECT b.id, u.user, u.username, u.password, u.role, u.is_active, 
	b.nama_lengkap, b.nik, b.nomor_telepon, b.pekerjaan, b.tempat_lahir, b.tanggal_lahir, b.kode_pos 
	FROM biodata_user b JOIN user_credential u ON u.id = b.user_credential_id`
	b.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WillReturnRows(rows)
	_, err := b.repo.FindAll()
	assert.Nil(b.T(), err)
	assert.NoError(b.T(), err)
}
func (b *BiodataUserRepoTestSuite) TestFindAll_Fail() {
	expectedSQL := `SELECT b.id, u.user, u.username, u.password, u.role, u.is_active, 
	b.nama_lengkap, b.nik, b.nomor_telepon, b.pekerjaan, b.tempat_lahir, b.tanggal_lahir, b.kode_pos 
	FROM biodata_user b JOIN user_credential u ON u.id = b.user_credential_id`
	b.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WillReturnError(errors.New("error"))
	bu, err := b.repo.FindAll()
	assert.Error(b.T(), err)
	assert.Nil(b.T(), bu)
}
func (b *BiodataUserRepoTestSuite) TestFindByNIK_Success() {
	mockData := model.BiodataUser{
		Id: "1",
		UserCredential: model.UserCredential{
			Id:       "1",
			Username: "akbar",
			Password: "123",
			Role:     "peminjam",
			IsActive: false,
		},
		NamaLengkap:  "akbar ismail",
		Nik:          "32010",
		NomorTelepon: "081123785743",
		Pekerjaan:    "mahasiswa",
		TempatLahir:  "jakarta",
		TanggalLahir: time.Date(2000, time.December, 1, 0, 0, 0, 0, time.UTC),
		KodePos:      "1610",
	}
	rows := sqlmock.NewRows([]string{"id", "user_credential_id", "user_credential_username", "user_credential_password", "user_credential_role", "user_credential_is_active", "nama_lengkap", "nik", "nomor_telepon", "pekerjaan", "tempat_lahir", "tanggal_lahir", "kode_pos"})
	rows.AddRow(mockData.Id, mockData.UserCredential.Id, mockData.UserCredential.Username, mockData.UserCredential.Password, mockData.UserCredential.Role, mockData.UserCredential.IsActive, mockData.NamaLengkap, mockData.Nik, mockData.NomorTelepon, mockData.Pekerjaan, mockData.TempatLahir, mockData.TanggalLahir, mockData.KodePos)
	expectedSQL := `SELECT b.id, u.user, u.username, u.password, u.role, u.is_active, 
	b.nama_lengkap, b.nik, b.nomor_telepon, b.pekerjaan, b.tempat_lahir, b.tanggal_lahir, b.kode_pos 
	FROM biodata_user b JOIN user_credential u ON u.id = b.user_credential_id WHERE b.nik =$1`
	b.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs(mockData.Nik).WillReturnRows(rows)
	bu, err := b.repo.FindByNIK(mockData.Nik)
	assert.Nil(b.T(), err)
	assert.NoError(b.T(), err)
	assert.Equal(b.T(), mockData, bu)
}
func (b *BiodataUserRepoTestSuite) TestFindByNIK_Fail() {
	expectedSQL := `SELECT b.id, u.user, u.username, u.password, u.role, u.is_active, 
	b.nama_lengkap, b.nik, b.nomor_telepon, b.pekerjaan, b.tempat_lahir, b.tanggal_lahir, b.kode_pos 
	FROM biodata_user b JOIN user_credential u ON u.id = b.user_credential_id WHERE b.nik =$1`
	b.mockSQL.ExpectQuery(regexp.QuoteMeta(expectedSQL)).WithArgs("32010").WillReturnError(errors.New("error"))
	bu, err := b.repo.FindByNIK("32010")
	assert.Error(b.T(), err)
	assert.NotNil(b.T(), err)
	assert.Equal(b.T(), model.BiodataUser{}, bu)
}
func (b *BiodataUserRepoTestSuite) TestSave_Success() {
	mockData := model.BiodataUser{
		Id: "1",
		UserCredential: model.UserCredential{
			Id:       "1",
			Username: "akbar",
			Password: "123",
			Role:     "peminjam",
			IsActive: false,
		},
		NamaLengkap:  "akbar ismail",
		Nik:          "32010",
		NomorTelepon: "081123785743",
		Pekerjaan:    "mahasiswa",
		TempatLahir:  "jakarta",
		TanggalLahir: time.Date(2000, time.December, 1, 0, 0, 0, 0, time.UTC),
		KodePos:      "1610",
	}
	expectedSQL := `INSERT INTO biodata_user`
	b.mockSQL.ExpectExec(expectedSQL).WithArgs(mockData.Id, mockData.UserCredential.Id, mockData.NamaLengkap, mockData.Nik, mockData.NomorTelepon, mockData.Pekerjaan, mockData.TempatLahir, mockData.TanggalLahir, mockData.KodePos).WillReturnResult(sqlmock.NewResult(1, 1))
	err := b.repo.Save(mockData)
	assert.Nil(b.T(), err)
	assert.NoError(b.T(), err)
}
func (b *BiodataUserRepoTestSuite) TestSave_Fail() {
	mockData := model.BiodataUser{
		Id: "1",
		UserCredential: model.UserCredential{
			Id:       "1",
			Username: "akbar",
			Password: "123",
			Role:     "peminjam",
			IsActive: false,
		},
		NamaLengkap:  "akbar ismail",
		Nik:          "32010",
		NomorTelepon: "081123785743",
		Pekerjaan:    "mahasiswa",
		TempatLahir:  "jakarta",
		TanggalLahir: time.Date(2000, time.December, 1, 0, 0, 0, 0, time.UTC),
		KodePos:      "1610",
	}
	expectedSQL := `INSERT INTO biodata_user`
	b.mockSQL.ExpectExec(expectedSQL).WithArgs(mockData.Id, mockData.UserCredential.Id, mockData.NamaLengkap, mockData.Nik, mockData.NomorTelepon, mockData.Pekerjaan, mockData.TempatLahir, mockData.TanggalLahir, mockData.KodePos).WillReturnError(errors.New("error"))
	err := b.repo.Save(mockData)
	assert.Error(b.T(), err)
	assert.NotNil(b.T(), err)
}
func (b *BiodataUserRepoTestSuite) TestUpdate_Success() {
	mockData := model.BiodataUser{
		Id: "1",
		UserCredential: model.UserCredential{
			Id:       "1",
			Username: "akbar",
			Password: "123",
			Role:     "peminjam",
			IsActive: false,
		},
		NamaLengkap:  "akbar ismail",
		Nik:          "32010",
		NomorTelepon: "081123785743",
		Pekerjaan:    "mahasiswa",
		TempatLahir:  "jakarta",
		TanggalLahir: time.Date(2000, time.December, 1, 0, 0, 0, 0, time.UTC),
		KodePos:      "1610",
	}
	expectedSQL := `UPDATE biodata_user SET user_credential_id = $2, nama_lengkap = $3, nik = $4, nomor_telepon = $5 
	, pekerjaan = $6, tempat_lahir = $7, tanggal_lahir = $8, kode_pos = $9 WHERE id = $1`
	b.mockSQL.ExpectExec(regexp.QuoteMeta(expectedSQL)).WithArgs(mockData.Id, mockData.UserCredential.Id, mockData.NamaLengkap, mockData.Nik, mockData.NomorTelepon, mockData.Pekerjaan, mockData.TempatLahir, mockData.TanggalLahir, mockData.KodePos).WillReturnResult(sqlmock.NewResult(1, 1))
	err := b.repo.Update(mockData)
	assert.Nil(b.T(), err)
	assert.NoError(b.T(), err)
}
func (b *BiodataUserRepoTestSuite) TestUpdate_Fail() {
	mockData := model.BiodataUser{
		Id: "1",
		UserCredential: model.UserCredential{
			Id:       "1",
			Username: "akbar",
			Password: "123",
			Role:     "peminjam",
			IsActive: false,
		},
		NamaLengkap:  "akbar ismail",
		Nik:          "32010",
		NomorTelepon: "081123785743",
		Pekerjaan:    "mahasiswa",
		TempatLahir:  "jakarta",
		TanggalLahir: time.Date(2000, time.December, 1, 0, 0, 0, 0, time.UTC),
		KodePos:      "1610",
	}
	expectedSQL := `UPDATE biodata_user SET user_credential_id = $2, nama_lengkap = $3, nik = $4, nomor_telepon = $5 
	, pekerjaan = $6, tempat_lahir = $7, tanggal_lahir = $8, kode_pos = $9 WHERE id = $1`
	b.mockSQL.ExpectExec(regexp.QuoteMeta(expectedSQL)).WithArgs(mockData.Id, mockData.UserCredential.Id, mockData.NamaLengkap, mockData.Nik, mockData.NomorTelepon, mockData.Pekerjaan, mockData.TempatLahir, mockData.TanggalLahir, mockData.KodePos).WillReturnError(errors.New("error"))
	err := b.repo.Update(mockData)
	assert.Error(b.T(), err)
	assert.NotNil(b.T(), err)
}
