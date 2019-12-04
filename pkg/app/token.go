package app

import (
	"github.com/dgrijalva/jwt-go"
	"go-echo-example/pkg/setting"
	"time"
)

type Claims struct {
	Id uint
	jwt.StandardClaims
}

func GenToken(id uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(setting.JWTExpire * time.Second)

	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "echo-example",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(setting.G_AppConf.JwtSecret))
	return token, err
}
