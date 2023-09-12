package manager

import "polen/usecase"

type UseCaseManager interface {
	AuthUseCase() usecase.AuthUseCase
	UserUseCase() usecase.UserUseCase
	BiodataUserUseCase() usecase.BiodataUserUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

// BiodataUserUseCase implements UseCaseManager.
func (u *useCaseManager) BiodataUserUseCase() usecase.BiodataUserUseCase {
	return usecase.NewBiodataUserUseCase(u.repoManager.BiodataRepo(), u.UserUseCase())
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
