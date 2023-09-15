package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"polen/mock/usecasemock"
	"polen/model/dto"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthControllerTestSuite struct {
	suite.Suite
	uuc    *usecasemock.UserUseCaseMock
	auc    *usecasemock.AuthUseCaseMock
	router *gin.Engine
}

func (a *AuthControllerTestSuite) SetupTest() {
	a.uuc = new(usecasemock.UserUseCaseMock)
	a.auc = new(usecasemock.AuthUseCaseMock)
	a.router = gin.Default()
}

var mockAuthRequest = dto.AuthRequest{
	Username: "akbar",
	Email:    "akbar@gmail.com",
	Password: "123",
	Role:     "peminjam",
}
var mockAuthResponse = dto.AuthResponse{
	Username: mockAuthRequest.Username,
	Token:    "",
}

func TestAuthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthControllerTestSuite))
}

func (a *AuthControllerTestSuite) TestLoginHandler_Success() {
	a.auc.On("Login", mockAuthRequest).Return(mockAuthResponse, nil)
	mockRg := a.router.Group("/api/v1")
	NewAuthController(a.uuc, a.auc, mockRg).Route()
	recorder := httptest.NewRecorder()
	payloadMarshal, err := json.Marshal(mockAuthRequest)
	assert.NoError(a.T(), err)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(payloadMarshal))

	a.router.ServeHTTP(recorder, request)
	response := recorder.Body.Bytes()

	var loginResponseSuccess struct {
		Message string
		Data    dto.AuthResponse
	}
	json.Unmarshal(response, &loginResponseSuccess)
	assert.Equal(a.T(), http.StatusOK, recorder.Code)
	assert.Equal(a.T(), "akbar", loginResponseSuccess.Data.Username)
	assert.Equal(a.T(), "successfully login", loginResponseSuccess.Message)
}
func (a *AuthControllerTestSuite) TestLoginHandler_ServerError() {
	a.auc.On("Login", mockAuthRequest).Return(mockAuthResponse, errors.New("failed"))
	mockRg := a.router.Group("/api/v1")
	NewAuthController(a.uuc, a.auc, mockRg).Route()
	recorder := httptest.NewRecorder()
	payloadMarshal, err := json.Marshal(mockAuthRequest)
	assert.NoError(a.T(), err)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(payloadMarshal))

	a.router.ServeHTTP(recorder, request)
	response := recorder.Body.Bytes()

	var loginResponseErr struct {
		Message string
	}
	json.Unmarshal(response, &loginResponseErr)
	assert.Equal(a.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(a.T(), "failed", loginResponseErr.Message)
}
func (a *AuthControllerTestSuite) TestLoginHandler_BindingError() {
	mockRg := a.router.Group("/api/v1")
	NewAuthController(a.uuc, a.auc, mockRg).Route()
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", nil)
	a.router.ServeHTTP(recorder, request)
	assert.Equal(a.T(), http.StatusBadRequest, recorder.Code)
}
func (a *AuthControllerTestSuite) TestRegisterHandler_Success() {
	a.uuc.On("Register", mockAuthRequest).Return(nil)
	mockRg := a.router.Group("/api/v1")
	NewAuthController(a.uuc, a.auc, mockRg).Route()
	recorder := httptest.NewRecorder()
	payloadMarshal, err := json.Marshal(mockAuthRequest)
	assert.NoError(a.T(), err)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(payloadMarshal))

	a.router.ServeHTTP(recorder, request)
	response := recorder.Body.Bytes()

	var registerResSuccess struct {
		Message string
	}
	json.Unmarshal(response, &registerResSuccess)
	assert.Equal(a.T(), http.StatusOK, recorder.Code)
	assert.Equal(a.T(), "successfully register", registerResSuccess.Message)
}
func (a *AuthControllerTestSuite) TestRegisterHandler_ServerError() {
	a.uuc.On("Register", mockAuthRequest).Return(errors.New("failed"))
	mockRg := a.router.Group("/api/v1")
	NewAuthController(a.uuc, a.auc, mockRg).Route()
	recorder := httptest.NewRecorder()
	payloadMarshal, err := json.Marshal(mockAuthRequest)
	assert.NoError(a.T(), err)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(payloadMarshal))

	a.router.ServeHTTP(recorder, request)
	response := recorder.Body.Bytes()

	var registerResponseErr struct {
		Message string
	}
	json.Unmarshal(response, &registerResponseErr)
	assert.Equal(a.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(a.T(), "failed", registerResponseErr.Message)
}
func (a *AuthControllerTestSuite) TestRegisterHandler_BindingError() {
	mockRg := a.router.Group("/api/v1")
	NewAuthController(a.uuc, a.auc, mockRg).Route()
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", nil)
	a.router.ServeHTTP(recorder, request)
	assert.Equal(a.T(), http.StatusBadRequest, recorder.Code)
}
