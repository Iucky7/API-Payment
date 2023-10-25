package repository

import (
	"api-payment/model"
	"api-payment/utils/constant"
	"database/sql"
	"fmt"
)

type MerchantRepository interface {
	BaseRepository[model.Merchant]
}

type merchantRepository struct {
	db *sql.DB
}

func (m *merchantRepository) Create(payload model.Merchant) error{
	_, err := m.db.Exec(constant.CREATE_MERCHANT,payload.Id,payload.Name)
	if err != nil{
		return nil
	}
	return nil
}

func (m *merchantRepository) List() ([]model.Merchant, error){
	var merchants []model.Merchant
	rows, err := m.db.Query(constant.LIST_MERCHANT)
	if err != nil{
		return nil,err
	}
	for rows.Next(){
		var merchant model.Merchant
		err := rows.Scan(&merchant.Id,&merchant.Name)
		if err != nil {
			return nil,fmt.Errorf("error scan merchant : %s", err.Error())
		}
		merchants = append(merchants,merchant)
	}
	return merchants,nil
}

func NewMerchantRepository(db *sql.DB) MerchantRepository{
	return &merchantRepository{
		db: db,
	}
}