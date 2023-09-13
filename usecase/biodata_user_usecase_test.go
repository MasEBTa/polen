package usecase

import (
	"errors"
	"polen/mock/repomock"
	"polen/mock/usecasemock"
	"polen/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BiodataUserUseCaseTestSuite struct {
	suite.Suite
	biurm *repomock.BiodataUserRepoMock
	uucm  *usecasemock.UserUseCaseMock
	buuc  BiodataUserUseCase
}

func (b *BiodataUserUseCaseTestSuite) SetupTest() {
	b.biurm = new(repomock.BiodataUserRepoMock)
	b.uucm = new(usecasemock.UserUseCaseMock)
	b.buuc = NewBiodataUserUseCase(b.biurm, b.uucm)
}
func TestBiodataUserUseCaseSuite(t *testing.T) {
	suite.Run(t, new(BiodataUserUseCaseTestSuite))
}

var mockData = model.BiodataUser{
	Id: "1",
	UserCredential: model.UserCredential{
		Id:       "1",
		Username: "akbar",
		Email:    "akbar@gmail.com",
		Password: "123",
		VANumber: "fdsfdhfdhfb",
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
var mockUserCred = model.UserCredential{
	Id:       "1",
	Username: "akbar",
	Email:    "akbar@gmail.com",
	Password: "123",
	VANumber: "fdsfdhfdhfb",
	Role:     "peminjam",
	IsActive: false,
}

func (b *BiodataUserUseCaseTestSuite) TestFindById_Success() {
	b.biurm.On("FindById", mockData.Id).Return(mockData, nil)
	bu, err := b.buuc.FindById(mockData.Id)
	assert.Nil(b.T(), err)
	assert.Equal(b.T(), mockData, bu)
}
func (b *BiodataUserUseCaseTestSuite) TestFindById_Fail() {
	b.biurm.On("FindById", "fhsdbfhd").Return(model.BiodataUser{}, errors.New("error"))
	bu, err := b.buuc.FindById("fhsdbfhd")
	assert.Error(b.T(), err)
	assert.NotNil(b.T(), err)
	assert.Equal(b.T(), model.BiodataUser{}, bu)
}
func (b *BiodataUserUseCaseTestSuite) TestFindAll_Success() {
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
	b.biurm.On("FindAll").Return(mockData, nil)
	bu, err := b.buuc.FindAll()
	assert.Nil(b.T(), err)
	assert.Equal(b.T(), mockData, bu)
}
func (b *BiodataUserUseCaseTestSuite) TestFindAll_Fail() {
	b.biurm.On("FindAll").Return(nil, errors.New("error"))
	_, err := b.buuc.FindAll()
	assert.Error(b.T(), err)
}
func (b *BiodataUserUseCaseTestSuite) TestFindByNIK_Success() {
	b.biurm.On("FindByNIK", mockData.Nik).Return(mockData, nil)
	bu, err := b.buuc.FindByNIK(mockData.Nik)
	assert.Nil(b.T(), err)
	assert.Equal(b.T(), mockData, bu)
}
func (b *BiodataUserUseCaseTestSuite) TestFindByNIK_Fail() {
	b.biurm.On("FindByNIK", "32010").Return(model.UserCredential{}, errors.New("error"))
	_, err := b.buuc.FindByNIK("32010")
	assert.Error(b.T(), err)
}
func (b *BiodataUserUseCaseTestSuite) TestDelete_Success() {
	b.biurm.On("FindById", mockData.Id).Return(mockData, nil)
	b.biurm.On("DeleteById", mockData.Id).Return(nil)
	err := b.buuc.Delete(mockData.Id)
	assert.Nil(b.T(), err)
}
func (b *BiodataUserUseCaseTestSuite) TestDelete_Fail() {
	b.biurm.On("FindById", mockData.Id).Return(mockData, nil)
	b.biurm.On("DeleteById", mockData.Id).Return(errors.New("failed to delete biodata"))
	err := b.buuc.Delete(mockData.Id)
	assert.Error(b.T(), err)
}
func (b *BiodataUserUseCaseTestSuite) TestDelete_BiodataIdInvalid() {
	b.biurm.On("FindById", "1").Return(model.BiodataUser{}, errors.New("error"))
	b.biurm.On("DeleteById", "1").Return(nil)
	err := b.buuc.Delete("1")
	assert.Error(b.T(), err)
}
func (b *BiodataUserUseCaseTestSuite) TestCreateNew_Success() {
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, nil)
	b.biurm.On("Save", mockData).Return(nil)
	err := b.buuc.CreateNew(mockData)
	assert.Nil(b.T(), err)
}
func (b *BiodataUserUseCaseTestSuite) TestCreateNew_Fail() {
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, nil)
	b.biurm.On("Save", mockData).Return(errors.New("failed to save new Biodata"))
	err := b.buuc.CreateNew(mockData)
	assert.Error(b.T(), err)
}
func (b *BiodataUserUseCaseTestSuite) TestCreateNew_UserIdInvalid() {
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, errors.New("error"))
	b.biurm.On("Save", mockData).Return(nil)
	err := b.buuc.CreateNew(mockData)
	assert.Error(b.T(), err)
}
func (b *BiodataUserUseCaseTestSuite) TestCreateNew_EmptyInvalid() {
	// nik required
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, nil)
	b.biurm.On("Save", model.BiodataUser{Id: "1", NamaLengkap: "akbar"}).Return(errors.New("nik is required"))
	err := b.buuc.CreateNew(model.BiodataUser{Id: "1", NamaLengkap: "akbar"})
	assert.Error(b.T(), err)
	// id user required
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, nil)
	b.biurm.On("Save", model.BiodataUser{Id: "1", NamaLengkap: "akbar", Nik: "32010"}).Return(errors.New("id user is required"))
	err = b.buuc.CreateNew(model.BiodataUser{Id: "1", NamaLengkap: "akbar", Nik: "32010"})
	assert.Error(b.T(), err)
	// name required
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, nil)
	b.biurm.On("Save", model.BiodataUser{Id: "1", Nik: "32010", UserCredential: model.UserCredential{Id: "1"}}).Return(errors.New("name is required"))
	err = b.buuc.CreateNew(model.BiodataUser{Id: "1", Nik: "32010", UserCredential: model.UserCredential{Id: "1"}})
	assert.Error(b.T(), err)
	// phone required
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, nil)
	b.biurm.On("Save", model.BiodataUser{Id: "1", NamaLengkap: "akbar", Nik: "32010", UserCredential: model.UserCredential{Id: "1"}}).Return(errors.New("phone is required"))
	err = b.buuc.CreateNew(model.BiodataUser{Id: "1", NamaLengkap: "akbar", Nik: "32010", UserCredential: model.UserCredential{Id: "1"}})
	assert.Error(b.T(), err)
	// job required
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, nil)
	b.biurm.On("Save", model.BiodataUser{Id: "1", NamaLengkap: "akbar", Nik: "32010", UserCredential: model.UserCredential{Id: "1"}, NomorTelepon: "081243345656"}).Return(errors.New("job is required"))
	err = b.buuc.CreateNew(model.BiodataUser{Id: "1", NamaLengkap: "akbar", Nik: "32010", UserCredential: model.UserCredential{Id: "1"}, NomorTelepon: "081243345656"})
	assert.Error(b.T(), err)
	// birth place required
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, nil)
	b.biurm.On("Save", model.BiodataUser{Id: "1", NamaLengkap: "akbar", Nik: "32010", UserCredential: model.UserCredential{Id: "1"}, NomorTelepon: "081243345656", Pekerjaan: "mahasiswa"}).Return(errors.New("birth place is required"))
	err = b.buuc.CreateNew(model.BiodataUser{Id: "1", NamaLengkap: "akbar", Nik: "32010", UserCredential: model.UserCredential{Id: "1"}, NomorTelepon: "081243345656", Pekerjaan: "mahasiswa"})
	assert.Error(b.T(), err)
}
func (b *BiodataUserUseCaseTestSuite) TestUpdate_Success() {
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, nil)
	b.biurm.On("Update", mockData).Return(nil)
	err := b.buuc.Update(mockData)
	assert.Nil(b.T(), err)
}
func (b *BiodataUserUseCaseTestSuite) TestUpdate_Fail() {
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, nil)
	b.biurm.On("Update", mockData).Return(errors.New("failed to update biodata"))
	err := b.buuc.Update(mockData)
	assert.Error(b.T(), err)
}
func (b *BiodataUserUseCaseTestSuite) TestUpdate_UserIdInvalid() {
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, errors.New("error"))
	b.biurm.On("Update", mockData).Return(nil)
	err := b.buuc.Update(mockData)
	assert.Error(b.T(), err)
}
func (b *BiodataUserUseCaseTestSuite) TestUpdate_EmptyInvalid() {
	// id biodata required
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, nil)
	b.biurm.On("Update", model.BiodataUser{}).Return(errors.New("id is required"))
	err := b.buuc.Update(model.BiodataUser{})
	assert.Error(b.T(), err)
	// nik required
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, nil)
	b.biurm.On("Update", model.BiodataUser{Id: "1", NamaLengkap: "akbar"}).Return(errors.New("nik is required"))
	err = b.buuc.Update(model.BiodataUser{Id: "1", NamaLengkap: "akbar"})
	assert.Error(b.T(), err)
	// id user required
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, nil)
	b.biurm.On("Update", model.BiodataUser{Id: "1", NamaLengkap: "akbar", Nik: "32010"}).Return(errors.New("id user is required"))
	err = b.buuc.Update(model.BiodataUser{Id: "1", NamaLengkap: "akbar", Nik: "32010"})
	assert.Error(b.T(), err)
	// name required
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, nil)
	b.biurm.On("Update", model.BiodataUser{Id: "1", Nik: "32010", UserCredential: model.UserCredential{Id: "1"}}).Return(errors.New("name is required"))
	err = b.buuc.Update(model.BiodataUser{Id: "1", Nik: "32010", UserCredential: model.UserCredential{Id: "1"}})
	assert.Error(b.T(), err)
	// phone required
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, nil)
	b.biurm.On("Update", model.BiodataUser{Id: "1", NamaLengkap: "akbar", Nik: "32010", UserCredential: model.UserCredential{Id: "1"}}).Return(errors.New("phone is required"))
	err = b.buuc.Update(model.BiodataUser{Id: "1", NamaLengkap: "akbar", Nik: "32010", UserCredential: model.UserCredential{Id: "1"}})
	assert.Error(b.T(), err)
	// job required
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, nil)
	b.biurm.On("Update", model.BiodataUser{Id: "1", NamaLengkap: "akbar", Nik: "32010", UserCredential: model.UserCredential{Id: "1"}, NomorTelepon: "081243345656"}).Return(errors.New("job is required"))
	err = b.buuc.Update(model.BiodataUser{Id: "1", NamaLengkap: "akbar", Nik: "32010", UserCredential: model.UserCredential{Id: "1"}, NomorTelepon: "081243345656"})
	assert.Error(b.T(), err)
	// birth place required
	b.uucm.On("FindById", mockUserCred.Id).Return(mockUserCred, nil)
	b.biurm.On("Update", model.BiodataUser{Id: "1", NamaLengkap: "akbar", Nik: "32010", UserCredential: model.UserCredential{Id: "1"}, NomorTelepon: "081243345656", Pekerjaan: "mahasiswa"}).Return(errors.New("birth place is required"))
	err = b.buuc.Update(model.BiodataUser{Id: "1", NamaLengkap: "akbar", Nik: "32010", UserCredential: model.UserCredential{Id: "1"}, NomorTelepon: "081243345656", Pekerjaan: "mahasiswa"})
	assert.Error(b.T(), err)
}
