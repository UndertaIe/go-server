package demo

import "github.com/gin-gonic/gin"

func Sentry(c *gin.Context) {
	panic("oops...")
}
