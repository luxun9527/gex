package utils

import (
	"log"
	"testing"
)

func TestGetJwt(t *testing.T) {
	jwt := NewJWT()
	claims := jwt.CreateClaims(JwtContent{
		UserID:   1,
		Username: "zhangsan",
		NickName: "",
	})
	token, err := jwt.CreateToken(claims)
	if err != nil {
		log.Fatal(err)
	}
	parseToken, err := jwt.ParseToken(token)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(parseToken)
	log.Println(token)
}
