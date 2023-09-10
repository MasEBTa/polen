package manager

import "polen/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
}

type repoManager struct {
	infraManager InfraManager
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
