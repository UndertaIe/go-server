package middleware

import (
	"github.com/UndertaIe/passwd/global"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/auth"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := errcode.Success
		token := c.GetHeader("Bearer")
		if token != "" {
			_, parseErr := auth.ParseToken(token, global.NewGlobal())
			if parseErr != nil {
				switch parseErr.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = errcode.UnauthorizedTokenTimeout
				case jwt.ValidationErrorSignatureInvalid:
					err = errcode.UnauthorizedTokenSignatureInvalid
				default:
					err = errcode.UnauthorizedTokenError
				}
			}
		} else {
			err = errcode.InvalidParams
		}
		if err != errcode.Success {
			app.NewResponse(c).ToError(err)
			c.Abort()
			return
		}
		c.Next()
	}
}

// 优化：将GetHeader ParseUserToken等方法抽象为接口作为参数传入
func UserJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := errcode.Success
		token := c.GetHeader("Bearer")
		var uClaims *auth.UserClaims
		var parseErr error
		if token != "" {
			uClaims, parseErr = auth.ParseUserToken(token, global.NewGlobal())
			if parseErr != nil {
				switch parseErr.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = errcode.UnauthorizedTokenTimeout
				case jwt.ValidationErrorSignatureInvalid:
					err = errcode.UnauthorizedTokenSignatureInvalid
				default:
					err = errcode.UnauthorizedTokenError
				}
			}
		} else {
			err = errcode.InvalidParams
		}
		if err != errcode.Success {
			app.NewResponse(c).ToError(err)
			c.Abort()
			return
		}
		c.Set("user_id", uClaims.UserId)
		c.Next()
	}
}
