package demo

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RateLimit(c *gin.Context) {
	w := c.Writer
	fmt.Fprint(w, "IP: ")
	fmt.Fprint(w, c.RemoteIP())
	fmt.Fprintln(w)

	fmt.Fprint(w, "Route: ")
	fmt.Fprint(w, c.FullPath())
	fmt.Fprintln(w)

	fmt.Fprintln(w, "<<< 通过了限流策略 >>>")
}
