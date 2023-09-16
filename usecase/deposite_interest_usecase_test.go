package usecase

import (
	"errors"
	"polen/mock"
	"polen/mock/repomock"
	"polen/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DepositeInteresetUseCaseTestSuite struct {
	suite.Suite
	dirm *repomock.DepositeInterestRepoMock
	diuc DepositeInterestUseCase
}

func (d *DepositeInteresetUseCaseTestSuite) SetupTest() {
	d.dirm = new(repomock.DepositeInterestRepoMock)
	d.diuc = NewDepositeInterestUseCase(d.dirm)
}
func TestDepositeInteresetUseCaseSuite(t *testing.T) {
	suite.Run(t, new(DepositeInteresetUseCaseTestSuite))
}
func (d *DepositeInteresetUseCaseTestSuite) TestDeleteById_Success() {
	d.dirm.On("FindById", mock.MockDepositeInterest.Id).Return(mock.MockDepositeInterest, nil)
	d.dirm.On("DeleteById", mock.MockDepositeInterest.Id).Return(nil)
	err := d.diuc.DeleteById(mock.MockDepositeInterest.Id)
	assert.Nil(d.T(), err)
}
func (d *DepositeInteresetUseCaseTestSuite) TestDeleteById_Failed() {
	d.dirm.On("FindById", mock.MockDepositeInterest.Id).Return(mock.MockDepositeInterest, nil)
	d.dirm.On("DeleteById", mock.MockDepositeInterest.Id).Return(errors.New("failed to delete deposite"))
	err := d.diuc.DeleteById(mock.MockDepositeInterest.Id)
	assert.Error(d.T(), err)
}
func (d *DepositeInteresetUseCaseTestSuite) TestDeleteById_IdDepositeInvalid() {
	d.dirm.On("FindById", mock.MockDepositeInterest.Id).Return(mock.MockDepositeInterest, errors.New("error"))
	d.dirm.On("DeleteById", mock.MockDepositeInterest.Id).Return(nil)
	err := d.diuc.DeleteById(mock.MockDepositeInterest.Id)
	assert.Error(d.T(), err)
}
func (d *DepositeInteresetUseCaseTestSuite) TestFindById_Success() {
	d.dirm.On("FindById", mock.MockDepositeInterest.Id).Return(mock.MockDepositeInterest, nil)
	di, err := d.diuc.FindById(mock.MockDepositeInterest.Id)
	assert.Nil(d.T(), err)
	assert.Equal(d.T(), mock.MockDepositeInterest.Id, di.Id)
}
func (d *DepositeInteresetUseCaseTestSuite) TestFindById_Failed() {
	d.dirm.On("FindById", mock.MockDepositeInterest.Id).Return(mock.MockDepositeInterest, errors.New("error"))
	di, err := d.diuc.FindById(mock.MockDepositeInterest.Id)
	assert.Error(d.T(), err)
	assert.Equal(d.T(), model.DepositeInterest{}, di)
}
func (d *DepositeInteresetUseCaseTestSuite) TestUpdate_Success() {
	d.dirm.On("FindById", mock.MockDepositeInterest.Id).Return(mock.MockDepositeInterest, nil)
	d.dirm.On("Update", mock.MockDepositeInterest).Return(nil)
	err := d.diuc.Update(mock.MockDepositeInterest)
	assert.Nil(d.T(), err)
}
func (d *DepositeInteresetUseCaseTestSuite) TestUpdate_Failed() {
	d.dirm.On("FindById", mock.MockDepositeInterest.Id).Return(mock.MockDepositeInterest, nil)
	d.dirm.On("Update", mock.MockDepositeInterest).Return(errors.New("failed to update deposite"))
	err := d.diuc.Update(mock.MockDepositeInterest)
	assert.Error(d.T(), err)
}
func (d *DepositeInteresetUseCaseTestSuite) TestUpdate_EmptyInvalid() {
	// Id required
	d.dirm.On("FindById", mock.MockDepositeInterest.Id).Return(mock.MockDepositeInterest, nil)
	d.dirm.On("Update", model.DepositeInterest{}).Return(errors.New("id is required"))
	err := d.diuc.Update(model.DepositeInterest{})
	assert.Error(d.T(), err)
}
func (d *DepositeInteresetUseCaseTestSuite) TestUpdate_IdDepositeInvalid() {
	d.dirm.On("FindById", mock.MockDepositeInterest.Id).Return(mock.MockDepositeInterest, errors.New("error"))
	d.dirm.On("Update", mock.MockDepositeInterest).Return(nil)
	err := d.diuc.Update(mock.MockDepositeInterest)
	assert.Error(d.T(), err)
}
func (d *DepositeInteresetUseCaseTestSuite) TestCreateNew_Success() {
	d.dirm.On("Save", mock.MockDepositeInterest).Return(nil)
	err := d.diuc.CreateNew(mock.MockDepositeInterest)
	assert.Nil(d.T(), err)
}
func (d *DepositeInteresetUseCaseTestSuite) TestCreateNew_Failed() {
	d.dirm.On("Save", mock.MockDepositeInterest).Return(errors.New("failed save Deposite Interest"))
	err := d.diuc.CreateNew(mock.MockDepositeInterest)
	assert.Error(d.T(), err)
}
func (d *DepositeInteresetUseCaseTestSuite) TestCreateNew_EmptyInvalid() {
	d.dirm.On("Save", model.DepositeInterest{}).Return(errors.New("id is required"))
	err := d.diuc.CreateNew(model.DepositeInterest{})
	assert.Error(d.T(), err)
}
