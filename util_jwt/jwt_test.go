package util_jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"reflect"
	"testing"
	"time"
)

func TestGenerateJWT(t *testing.T) {
	fmt.Println(GenerateJWT(map[string]interface{}{
		"user_id": int(1),
		"version": 1,
		"exp":     time.Now().Add(2 * time.Hour).Unix(),
	}))
}

func TestValidateJWT(t *testing.T) {
	token, msg := ValidateJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTAyMDQ5OTMsInVzZXJfaWQiOjEsInZlcnNpb24iOjF9.3Cjs0eJTRloYrUboIndMrPsjiF2gw56fug0Bu11nYU8")

	if !token.Valid {
		fmt.Println("valid fail", msg)
		return
	}
	fmt.Println(token.Claims.(jwt.MapClaims)["user_id"])
	fmt.Println(reflect.TypeOf(token.Claims.(jwt.MapClaims)["user_id"]))

	var r map[string]interface{}
	r = token.Claims.(jwt.MapClaims)

	fmt.Println(r)
	fmt.Println(r["user_id"])
	fmt.Println(reflect.TypeOf(r["user_id"]))
}
