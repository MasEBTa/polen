package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"polen/mock/usecasemock"
	"polen/model"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BiodataUserControllerTestSuite struct {
	suite.Suite
	buucm  *usecasemock.BiodataUserUseCaseMock
	router *gin.Engine
}

func (b *BiodataUserControllerTestSuite) SetupTest() {
	b.buucm = new(usecasemock.BiodataUserUseCaseMock)
	b.router = gin.Default()
}
func TestBiodataUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BiodataUserControllerTestSuite))
}

var mockBiodataUser = model.BiodataUser{
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

func (b *BiodataUserControllerTestSuite) TestDeleteHandler_Success() {
	b.buucm.On("Delete", "1").Return(nil)
	rg := b.router.Group("/api/v1")
	NewBiodataController(b.buucm, rg).Route()
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "/api/v1/biodata/1", nil)
	b.router.ServeHTTP(recorder, request)
	response := recorder.Body.Bytes()
	var deleteResponse struct {
		Message string
	}
	json.Unmarshal(response, &deleteResponse)
	assert.Equal(b.T(), http.StatusOK, recorder.Code)
	assert.Equal(b.T(), "successfully delete biodata with id 1", deleteResponse.Message)
}
func (b *BiodataUserControllerTestSuite) TestDeleteHandler_ServerError() {
	b.buucm.On("Delete", "1").Return(errors.New("failed"))
	rg := b.router.Group("/api/v1")
	NewBiodataController(b.buucm, rg).Route()
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "/api/v1/biodata/1", nil)
	b.router.ServeHTTP(recorder, request)
	assert.Equal(b.T(), http.StatusInternalServerError, recorder.Code)
}
func (b *BiodataUserControllerTestSuite) TestGetByIdHandler_Success() {
	b.buucm.On("FindById", mockBiodataUser.Id).Return(mockBiodataUser, nil)
	rg := b.router.Group("/api/v1")
	NewBiodataController(b.buucm, rg).Route()
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/v1/biodata/1", nil)
	b.router.ServeHTTP(recorder, request)
	response := recorder.Body.Bytes()
	var biodataUser model.BiodataUser
	json.Unmarshal(response, &biodataUser)
	assert.Equal(b.T(), http.StatusOK, recorder.Code)
	assert.Equal(b.T(), mockBiodataUser, biodataUser)
}
func (b *BiodataUserControllerTestSuite) TestGetByIdHandler_ServerError() {
	b.buucm.On("FindById", mockBiodataUser.Id).Return(mockBiodataUser, errors.New("failed"))
	rg := b.router.Group("/api/v1")
	NewBiodataController(b.buucm, rg).Route()
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/v1/biodata/1", nil)
	b.router.ServeHTTP(recorder, request)
	assert.Equal(b.T(), http.StatusInternalServerError, recorder.Code)
}
func (b *BiodataUserControllerTestSuite) TestListHandler_Success() {
	mockData := []model.BiodataUser{
		{
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
		},
	}
	b.buucm.On("FindAll").Return(mockData, nil)
	rg := b.router.Group("/api/v1")
	NewBiodataController(b.buucm, rg).Route()
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/v1/biodata", nil)
	b.router.ServeHTTP(recorder, request)
	response := recorder.Body.Bytes()
	var biodatas []model.BiodataUser
	json.Unmarshal(response, &biodatas)
	assert.Equal(b.T(), http.StatusOK, recorder.Code)
	assert.Equal(b.T(), mockData, biodatas)
}
func (b *BiodataUserControllerTestSuite) TestListHandler_ServerError() {
	b.buucm.On("FindAll").Return(nil, errors.New("not found"))
	rg := b.router.Group("/api/v1")
	NewBiodataController(b.buucm, rg).Route()
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/api/v1/biodata", nil)
	b.router.ServeHTTP(recorder, request)
	assert.Equal(b.T(), http.StatusInternalServerError, recorder.Code)
}
func (b *BiodataUserControllerTestSuite) TestUpdateHandler_Success() {
	b.buucm.On("Update", mockBiodataUser).Return(nil)
	rg := b.router.Group("/api/v1")
	NewBiodataController(b.buucm, rg).Route()
	recorder := httptest.NewRecorder()
	payloadMarshal, err := json.Marshal(mockBiodataUser)
	assert.NoError(b.T(), err)
	request := httptest.NewRequest(http.MethodPut, "/api/v1/biodata", bytes.NewBuffer(payloadMarshal))
	b.router.ServeHTTP(recorder, request)
	response := recorder.Body.Bytes()
	var updateResponse struct {
		Message string
	}
	json.Unmarshal(response, &updateResponse)
	assert.Equal(b.T(), http.StatusOK, recorder.Code)
	assert.Equal(b.T(), "successfully update biodata", updateResponse.Message)
}
func (b *BiodataUserControllerTestSuite) TestUpdateHandler_ServerError() {
	b.buucm.On("Update", mockBiodataUser).Return(errors.New("failed"))
	rg := b.router.Group("/api/v1")
	NewBiodataController(b.buucm, rg).Route()
	recorder := httptest.NewRecorder()
	payloadMarshal, err := json.Marshal(mockBiodataUser)
	assert.NoError(b.T(), err)
	request := httptest.NewRequest(http.MethodPut, "/api/v1/biodata", bytes.NewBuffer(payloadMarshal))
	b.router.ServeHTTP(recorder, request)
	response := recorder.Body.Bytes()
	var updateResponse struct {
		Message string
	}
	json.Unmarshal(response, &updateResponse)
	assert.Equal(b.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(b.T(), "failed", updateResponse.Message)
}
func (b *BiodataUserControllerTestSuite) TestUpdateHandler_BindingError() {
	rg := b.router.Group("/api/v1")
	NewBiodataController(b.buucm, rg).Route()
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPut, "/api/v1/biodata", nil)
	b.router.ServeHTTP(recorder, request)
	assert.Equal(b.T(), http.StatusBadRequest, recorder.Code)
}
func (b *BiodataUserControllerTestSuite) TestCreateHandler_ServerError() {
	b.buucm.On("CreateNew", mockBiodataUser).Return(errors.New("failed"))
	rg := b.router.Group("/api/v1")
	NewBiodataController(b.buucm, rg).Route()
	recorder := httptest.NewRecorder()
	payloadMarshal, err := json.Marshal(mockBiodataUser)
	assert.NoError(b.T(), err)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/biodata", bytes.NewBuffer(payloadMarshal))
	b.router.ServeHTTP(recorder, request)
	response := recorder.Body.Bytes()
	var biodata model.BiodataUser
	json.Unmarshal(response, &biodata)
	assert.Equal(b.T(), http.StatusInternalServerError, recorder.Code)
	assert.NotEqual(b.T(), mockBiodataUser.Id, biodata.Id)
}
func (b *BiodataUserControllerTestSuite) TestCreateHandler_BindingError() {
	rg := b.router.Group("/api/v1")
	NewBiodataController(b.buucm, rg).Route()
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/api/v1/biodata", nil)
	b.router.ServeHTTP(recorder, request)
	assert.Equal(b.T(), http.StatusBadRequest, recorder.Code)
}
