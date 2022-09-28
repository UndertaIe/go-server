package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtOptions interface { // 外部调用时需要实现的接口（用于控制token失效时间、JWT密钥和颁发者等参数）
	GetJWTSecret() []byte
	GetJWTIssuer() string
	GetJWTExpireTime() int64
}

type RoleClaims struct {
	Role
	jwt.StandardClaims
}

type Role struct {
	UserId int `json:"user_id"`
	RoleId int `json:"role_id"`
}

type RoleLevel int

const (
	Public RoleLevel = iota - 1
	User
	Admin
)

func (r RoleLevel) Pass(level int) bool {
	return r <= RoleLevel(level)
}

func (r RoleLevel) IsPublic() bool {
	return r == Public
}

func GenerateJwtToken(r Role, opt JwtOptions) (string, error) {
	claims := RoleClaims{
		Role: r,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: opt.GetJWTExpireTime(),
			Issuer:    opt.GetJWTIssuer(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(opt.GetJWTSecret())
	return token, err
}

func ParseJwtToken(token string, opt JwtOptions) (*RoleClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &RoleClaims{}, func(token *jwt.Token) (interface{}, error) {
		return opt.GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*RoleClaims)
		if ok && tokenClaims.Valid {
			if claims.UserId == 0 {
				return nil, &jwt.ValidationError{}
			}
			return claims, nil
		}
	}

	return nil, err
}
