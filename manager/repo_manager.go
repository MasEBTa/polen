package manager

import "polen/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	BiodataRepo() repository.BiodataUser
}

type repoManager struct {
	infraManager InfraManager
}

// BiodataRepo implements RepoManager.
func (r *repoManager) BiodataRepo() repository.BiodataUser {
	return repository.NewBiodataUserRepository(r.infraManager.Conn())

}

// UserRepo implements RepoManager.
func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infraManager.Conn())
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
