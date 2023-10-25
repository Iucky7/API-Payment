package manager

import "api-payment/repository"

type RepoManager interface {
	MerchantRepo() repository.MerchantRepository
	PaymentRepo() repository.PaymentRepository
	UserRepo() repository.UserRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) MerchantRepo() repository.MerchantRepository {
	return repository.NewMerchantRepository(r.infra.Conn())
}

func (r *repoManager) PaymentRepo() repository.PaymentRepository {
	return repository.NewPaymentRepository(r.infra.Conn())
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func NewRepoManager(infraParam InfraManager) RepoManager {
	return &repoManager{
		infra: infraParam,
	}
}