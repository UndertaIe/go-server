package demo

import (
	"fmt"

	"github.com/UndertaIe/go-eden/app"
	"github.com/UndertaIe/go-eden/cache"
	"github.com/gin-gonic/gin"
)

func DeleteUserWithTracing(c *gin.Context) {
	resp := app.NewResponse(c)
	user_id := c.Param("id")
	cache.DeleteCacheWithTracing(c)
	resp.To(fmt.Sprintf("已清空用户(%s)缓存", user_id))
}

func UpdateUserWithTracing(c *gin.Context) {
	resp := app.NewResponse(c)
	user_id := c.Param("id")
	cache.DeleteCacheWithTracing(c)
	resp.To(fmt.Sprintf("已清空用户(%s)缓存", user_id))
}
