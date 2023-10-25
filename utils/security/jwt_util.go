package security

import (
	"api-payment/config"
	"api-payment/model"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var blacklist = make(map[string]bool)

func CreateAccessToken(user model.UserCredential) (string,error){
	cfg, _ := config.NewConfig()

	now := time.Now().UTC()
	end := now.Add(cfg.AccessTokenLifeTime)

	claims := &TokenMyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: cfg.ApplicationName,
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: user.Username,
	}

	token := jwt.NewWithClaims(cfg.JwtSigningMethod,claims)
	ss, err := token.SignedString(cfg.JwtSignatureKey)
	if err!=nil{
		return "",fmt.Errorf("failed to create access token : %s",err.Error())
	}
	return ss,nil
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims,error){
	cfg, _ := config.NewConfig()
	token, err := jwt.Parse(tokenString,func(t *jwt.Token)(interface{},error)  {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != cfg.JwtSigningMethod{
			return nil,fmt.Errorf("invalid token signing method")
		}
		return cfg.JwtSignatureKey,nil
	})
	if err!= nil{
		return nil,err
	}
	
	if _, ok := blacklist[tokenString]; ok {
		return nil, fmt.Errorf("Please Login Again")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != cfg.ApplicationName{
		return nil,fmt.Errorf("Invalid token MapClaims")
	}
	return claims,nil

}

func Logout(tokenString string) error {
	blacklist[tokenString] = true
	return nil
}