package manager

import (
	"polen/usecase"

	"github.com/gin-gonic/gin"
)

type UseCaseManager interface {
	AuthUseCase() usecase.AuthUseCase
	UserUseCase() usecase.UserUseCase
	BiodataUserUseCase() usecase.BiodataUserUseCase
	TopUpUsecase() usecase.TopUpUseCase
	DepositerInterestUseCase() usecase.DepositeInterestUseCase
}

type useCaseManager struct {
	repoManager RepoManager
	ctx         *gin.Context
}

// DepositerInterestUseCase implements UseCaseManager.
func (u *useCaseManager) DepositerInterestUseCase() usecase.DepositeInterestUseCase {
	return usecase.NewDepositeInterestUseCase(u.repoManager.DepositeInterestRepo())
}

// TopUpUsecase implements UseCaseManager.
func (u *useCaseManager) TopUpUsecase() usecase.TopUpUseCase {
	return usecase.NewTopUpUseCase(u.repoManager.TopUpRepo(), u.UserUseCase())
}

// BiodataUserUseCase implements UseCaseManager.
func (u *useCaseManager) BiodataUserUseCase() usecase.BiodataUserUseCase {
	return usecase.NewBiodataUserUseCase(u.repoManager.BiodataRepo(), u.UserUseCase(), u.ctx)
}

// AuthUseCase implements UseCaseManager.
func (u *useCaseManager) AuthUseCase() usecase.AuthUseCase {
	return usecase.NewAuthUseCase(u.repoManager.UserRepo())
}

// UserUseCase implements UseCaseManager.
func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.UserRepo())
}

func NewUsecaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
