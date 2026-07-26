package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// Header.Payload.Signature
// jwt密钥
var jwtSecret = []byte("memo-secret-key")

// 生成token（令牌）
func GenerateToken(userId uint) (string, error) {
	//创建claims
	claims := jwt.MapClaims{
		"userId": userId,
		// 过期时间
		"exp": time.Now().
			Add(24 * time.Hour).
			Unix(),
	}

	//生成token
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	//签名
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (uint, error) {
	// 解析token
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}
	userID, ok := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("invalid userId")
	}
	return uint(userID), nil
}
