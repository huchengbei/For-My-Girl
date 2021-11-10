package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/huchengbei/for-my-girl/pkg/setting"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-for-my-girl",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokeClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokeClaims != nil {
		if claims, ok := tokeClaims.Claims.(*Claims); ok && tokeClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
