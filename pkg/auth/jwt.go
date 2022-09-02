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

type ClaimsInterface interface { // 外部调用时需要实现的接口（用于控制token失效时间、JWT密钥和颁发者等参数）
	GetJWTSecret() []byte
	GetJWTIssuer() string
	GetJWTExpireTime() int64
}

func GenerateToken(appKey, appSecret string, i ClaimsInterface) (string, error) {
	md5 := utils.NewHasher(crypto.MD5)
	claims := Claims{
		AppKey:    md5.Hash(appKey),
		AppSecret: md5.Hash(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: i.GetJWTExpireTime(),
			Issuer:    i.GetJWTIssuer(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(i.GetJWTSecret())
	return token, err
}

func ParseToken(token string, i ClaimsInterface) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return i.GetJWTSecret(), nil
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

// TODO: 待优化
type UserClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

func GenerateUserToken(user_id int, i ClaimsInterface) (string, error) {
	claims := UserClaims{
		UserId: user_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: i.GetJWTExpireTime(),
			Issuer:    i.GetJWTIssuer(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(i.GetJWTSecret())
	return token, err
}

func ParseUserToken(token string, i ClaimsInterface) (*UserClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return i.GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*UserClaims)
		if ok && tokenClaims.Valid {
			if claims.UserId == 0 {
				return nil, &jwt.ValidationError{}
			}
			return claims, nil
		}
	}

	return nil, err
}
