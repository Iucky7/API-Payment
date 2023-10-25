package usecase

import (
	"api-payment/model"
	"api-payment/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	RegisterNewUser(payload model.UserCredential) error
	FindAllUser() ([]model.UserCredential,error)
	FindByUsername(username string) (model.UserCredential,error)
	FindByUsernamePassword(username string, password string) (model.UserCredential,error)
}

type userUseCase struct{
	repo repository.UserRepository
}

func (u *userUseCase) RegisterNewUser(payload model.UserCredential) error{
	bytes, _ := bcrypt.GenerateFromPassword([]byte(payload.Password),bcrypt.DefaultCost)
	payload.Password = string(bytes)
	err := u.repo.Create(payload)
	if err != nil{
		return fmt.Errorf("failed to create new user : %s",err.Error())
	}
	return nil	
}

func (u *userUseCase) FindAllUser() ([]model.UserCredential,error){
	return u.repo.List()
}

func (u *userUseCase) FindByUsername(username string) (model.UserCredential,error){
	return u.repo.GetByUsername(username)
}

func (u *userUseCase) FindByUsernamePassword(username string, password string) (model.UserCredential,error){
	return u.repo.GetByUsernamePassword(username,password)
}

func NewUserUseCase (repository repository.UserRepository) UserUseCase{
	return &userUseCase{
		repo: repository,
	}
}