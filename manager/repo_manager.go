package manager

import "polen/repository"

type RepoManager interface {
	UserRepo() repository.UserRepository
	BiodataRepo() repository.BiodataUser
	TopUpRepo() repository.TopUp
	DepositeInterestRepo() repository.DepositeInterest
	LoanInterestRepo() repository.LoanInterest
	SaldoRepo() repository.SaldoRepository
}

type repoManager struct {
	infraManager InfraManager
}

// LoanInterestRepo implements RepoManager.
func (r *repoManager) LoanInterestRepo() repository.LoanInterest {
	return repository.NewLoanInterestRepository(r.infraManager.Conn())
}

// SaldoRepo implements RepoManager.
func (r *repoManager) SaldoRepo() repository.SaldoRepository {
	return repository.NewSaldoRepository(r.infraManager.Conn())
}

// DepositeInterestRepo implements RepoManager.
func (r *repoManager) DepositeInterestRepo() repository.DepositeInterest {
	return repository.NewDepositeInterestRepository(r.infraManager.Conn())
}

// TopUpRepo implements RepoManager.
func (r *repoManager) TopUpRepo() repository.TopUp {
	return repository.NewTopUpRepository(r.infraManager.Conn())
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
