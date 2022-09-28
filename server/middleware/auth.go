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

var g = global.NewGlobal()

// 优化：将GetHeader ParseUserToken等方法抽象为接口作为参数传入
func RoleAuth(level auth.RoleLevel) gin.HandlerFunc {
	if level.IsPublic() {
		return func(c *gin.Context) { c.Next() }
	}
	return func(c *gin.Context) {
		err := errcode.Success
		token := c.GetHeader("Bearer")
		var rClaims *auth.RoleClaims
		var parseErr error
		if token != "" {
			rClaims, parseErr = auth.ParseJwtToken(token, g)
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
			err = errcode.UnauthorizedTokenError
		}
		if err != errcode.Success {
			app.NewResponse(c).ToError(err)
			c.Abort()
			return
		}
		if !level.Pass(rClaims.RoleId) {
			app.NewResponse(c).ToError(errcode.UnauthorizedUserError)
			c.Abort()
			return
		}
		c.Set("user_id", rClaims.UserId)
		c.Set("role_id", rClaims.RoleId)
		c.Next()
	}
}
