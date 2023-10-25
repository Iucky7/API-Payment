package repository

import (
	"api-payment/model"
	"api-payment/utils/constant"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	BaseRepository[model.UserCredential]
	GetByUsername(username string) (model.UserCredential, error)
	GetByUsernamePassword(username string, password string) (model.UserCredential, error)
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Create(payload model.UserCredential) error{
	_, err := u.db.Exec(constant.CREATE_USER_CREDENTIAL,payload.Id,payload.Username,payload.Password)
	if err != nil{
		return nil
	}
	return nil
}

func (u *userRepository) List() ([]model.UserCredential, error){
	var users []model.UserCredential
	rows, err := u.db.Query(constant.LIST_USER_CREDENTIAL)
	if err != nil{
		return nil,err
	}
	for rows.Next(){
		var user model.UserCredential
		err := rows.Scan(&user.Id,&user.Username,&user.Password)
		if err != nil {
			return nil,fmt.Errorf("error scan user : %s", err.Error())
		}
		users = append(users,user)
	}
	return users,nil
}

func (u *userRepository) GetByUsername(username string) (model.UserCredential, error){
	var user model.UserCredential
	err := u.db.QueryRow(constant.GET_USER_CREDENTIAL_BY_USERNAME,username).Scan(&user.Id,&user.Username,&user.Password)
	if err != nil{
		return model.UserCredential{},err
	}
	return user,nil
}

func (u *userRepository)GetByUsernamePassword(username string, password string) (model.UserCredential, error){
	user, err := u.GetByUsername(username)
	if err != nil{
		return model.UserCredential{},nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil{
		return model.UserCredential{},nil
	}
	return user,nil
}

func NewUserRepository(db *sql.DB) UserRepository{
	return &userRepository{
		db: db,
	}
}