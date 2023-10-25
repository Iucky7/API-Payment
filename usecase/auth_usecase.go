package usecase

import (
	"api-payment/utils/security"
	"fmt"
)

type AuthUseCase interface {
	Login(username string, password string) (string, error)
	Logout(token string) error
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

func (a *authUseCase) Logout(token string) error {
	err := security.Logout(token)
	if err != nil {
		return fmt.Errorf("failed to log out")
	}

	return nil
}

func NewAuthUseCase(userUseCase UserUseCase) AuthUseCase {
	return &authUseCase{
		userUc: userUseCase,
	}
}