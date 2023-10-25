package manager

import "api-payment/usecase"

type UseCaseManager interface {
	MerchantUseCase() usecase.MerchantUseCase
	PaymentUseCase() usecase.PaymentUseCase
	UserUseCase() usecase.UserUseCase
	AuthUseCase() usecase.AuthUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) MerchantUseCase() usecase.MerchantUseCase {
	return usecase.NewMerchantUseCase(u.repoManager.MerchantRepo())
}

func (u *useCaseManager) PaymentUseCase() usecase.PaymentUseCase {
	return usecase.NewPaymentUseCase(u.repoManager.PaymentRepo())
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.UserRepo())
}

func (u *useCaseManager) AuthUseCase() usecase.AuthUseCase {
	return usecase.NewAuthUseCase(u.UserUseCase())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repo,
	}
}