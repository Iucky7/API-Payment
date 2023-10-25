package repository

import (
	"api-payment/model"
	"api-payment/utils/constant"
	"database/sql"
	"fmt"
)

type PaymentRepository interface {
	BaseRepository[model.Payment]
}

type paymentRepository struct {
	db *sql.DB
}

func (p *paymentRepository) Create(payload model.Payment) error{
	_, err := p.db.Exec(constant.CREATE_PAYMENT,payload.Id,payload.MerchantId,payload.BankAccount,payload.Amount)
	if err != nil{
		return nil
	}
	return nil
}

func (p *paymentRepository) List() ([]model.Payment, error){
	var payments []model.Payment
	rows, err := p.db.Query(constant.LIST_PAYMENT)
	if err != nil{
		return nil,err
	}
	for rows.Next(){
		var payment model.Payment
		err := rows.Scan(&payment.Id,&payment.MerchantId,&payment.BankAccount,&payment.Amount)
		if err != nil {
			return nil,fmt.Errorf("error scan payment : %s", err.Error())
		}
		payments = append(payments,payment)
	}
	return payments,nil
}

func NewPaymentRepository(db *sql.DB) PaymentRepository{
	return &paymentRepository{
		db: db,
	}
}