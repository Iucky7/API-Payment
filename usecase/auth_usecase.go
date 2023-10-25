package usecase

import (
	"api-payment/utils/security"
	"fmt"
)

type AuthUseCase interface {
	Login(username string, password string) (string, error)
}

type authUseCase struct {
	userUc UserUseCase
}

func (a *authUseCase) Login(username string, password string) (string, error) {
	user, err := a.userUc.FindByUsernamePassword(username, password)
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}
	
	token, err := security.CreateAccessToken(user)
	if err != nil {
		return "", fmt.Errorf("failed to generate token : %s ", err.Error())
	}
	return token, nil
}

func NewAuthUseCase(userUseCase UserUseCase) AuthUseCase {
	return &authUseCase{
		userUc: userUseCase,
	}
}