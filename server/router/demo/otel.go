package demo

import (
	"fmt"

	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/cache"
	"github.com/gin-gonic/gin"
)

func DeleteUserWithTracing(c *gin.Context) {
	resp := app.Response{Ctx: c}
	user_id := c.Param("id")
	cache.DeleteCacheWithTracing(c)
	resp.To(fmt.Sprintf("已清空用户(%s)缓存", user_id))
}

func UpdateUserWithTracing(c *gin.Context) {
	resp := app.Response{Ctx: c}
	user_id := c.Param("id")
	cache.DeleteCacheWithTracing(c)
	resp.To(fmt.Sprintf("已清空用户(%s)缓存", user_id))
}
