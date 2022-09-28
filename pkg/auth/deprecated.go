package auth

import (
	"crypto"

	"github.com/UndertaIe/passwd/pkg/utils"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GenerateToken(appKey, appSecret string, opt JwtOptions) (string, error) {
	md5 := utils.NewHasher(crypto.MD5)
	claims := Claims{
		AppKey:    md5.Hash(appKey),
		AppSecret: md5.Hash(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: opt.GetJWTExpireTime(),
			Issuer:    opt.GetJWTIssuer(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(opt.GetJWTSecret())
	return token, err
}

func ParseToken(token string, opt JwtOptions) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return opt.GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
