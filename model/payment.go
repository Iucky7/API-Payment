package model

import "time"

type Payment struct {
	Id          string `json:"id"`
	MerchantId  string `json:"merchantId"`
	BankAccount int `json:"bankAccount"`
	Amount      int64 `json:"amount"`
	PaymentDate time.Time `json:"paymentDate"`
}
