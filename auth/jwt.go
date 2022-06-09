package auth

import (
	"douying/setting"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateToken(username string) (string,error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp" : time.Now().Unix() + 5000,
	}).SignedString([]byte(setting.Conf.SigningKey))
	return token,err
}

func CheckToken(tokenString string) (string,error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(setting.Conf.SigningKey), nil
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s",token.Claims.(jwt.MapClaims)["username"]), err
}
