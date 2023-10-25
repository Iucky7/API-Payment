package usecase

import (
	"api-payment/model"
	"api-payment/repository"
	"fmt"
)

type MerchantUseCase interface {
	RegisterNewMerchant(payload model.Merchant) error
	FindAllMerchantList() ([]model.Merchant, error)
}

type merchantUseCase struct {
	repo repository.MerchantRepository
}

func (m *merchantUseCase) RegisterNewMerchant(payload model.Merchant) error {
	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}
	err := m.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create Merchant : %s", err.Error())
	}
	return nil
}

func (m *merchantUseCase) FindAllMerchantList() ([]model.Merchant, error) {
	return m.repo.List()
}

func NewMerchantUseCase(merchantRepo repository.MerchantRepository) MerchantUseCase {
	return &merchantUseCase{
		repo: merchantRepo,
	}
}
