package util_jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

const (
	// SecretKey JWT密钥
	SecretKey = "zxk$&=*ek$t0u(jfzn^0dfzm7w1vau*r*$%50n@&cq@d7wmoa$"
)

// GenerateJWT 生成JWT
func GenerateJWT(claims map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
	jwtString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", errors.New(err.Error())
	}
	return jwtString, nil
}

// ValidateJWT 校验JWT
func ValidateJWT(t string) (*jwt.Token, string) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return token, err.Error()
	}
	return token, ""
}
