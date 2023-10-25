package usecase

import (
	"api-payment/model"
	"api-payment/repository"
	"fmt"
	"time"
)

type PaymentUseCase interface {
	RegisterNewPayment(payload model.Payment) error
	FindAllPaymentList() ([]model.Payment, error)
}

type paymentUseCase struct {
	repo repository.PaymentRepository
}

func (p *paymentUseCase) RegisterNewPayment(payload model.Payment) error {
	if payload.MerchantId == "" {
		return fmt.Errorf("merchantId is required")
	}
	
	payload.PaymentDate = time.Now()
	err := p.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create Payment : %s", err.Error())
	}
	return nil
}

func (p *paymentUseCase) FindAllPaymentList() ([]model.Payment, error) {
	return p.repo.List()
}

func NewPaymentUseCase(paymentRepo repository.PaymentRepository) PaymentUseCase {
	return &paymentUseCase{
		repo: paymentRepo,
	}
}
