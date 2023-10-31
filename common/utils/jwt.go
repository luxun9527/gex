package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtContent struct {
	UserID   int64
	Username string
	NickName string
}

const (
	TokenRedisKey = "jwt_black_list_"
)

type CustomClaims struct {
	JwtContent
	jwt.RegisteredClaims
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

const singingKey = "874cb20cd4a3e0db979c8135300c12fe"
const ExpiresTime = time.Second * 60 * 60 * 24 * 10

func NewJWT() *JWT {
	return &JWT{

		[]byte(singingKey),
	}
}
func GenerateKey(token string) string {
	return TokenRedisKey + MD5(token)
}
func (j *JWT) CreateClaims(baseClaims JwtContent) CustomClaims {
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"GVA"},                         // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),       // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ExpiresTime)), // 过期时间 7天  配置文件
			Issuer:    "github.com/luxun9527/gex",                      // 签名的发行者
		},
		JwtContent: baseClaims,
	}
	return claims
}

// 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
