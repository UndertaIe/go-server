package demo

import (
	"fmt"
	"strconv"
	"time"

	"github.com/UndertaIe/passwd/internal/service"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/cache"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Now(c *gin.Context) {
	resp := app.Response{Ctx: c}
	now := time.Now().UTC().GoString()
	resp.To(now)
}

func CacheNow(c *gin.Context) {
	resp := app.Response{Ctx: c}
	now := time.Now().UTC().GoString()
	resp.To(now)
}

func GetUser(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("id"))
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	param := service.UserGetRequest{UserId: user_id}
	srv := service.NewService(c.Request.Context())
	log.Info("缓存未命中，访问数据库...")
	user, err := srv.GetUser(&param)

	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.To(user)
}

func DeleteUser(c *gin.Context) {
	resp := app.Response{Ctx: c}
	user_id := c.Param("id")
	cache.DeleteCache(c)
	resp.To(fmt.Sprintf("已清空用户(%s)缓存", user_id))
}

func UpdateUser(c *gin.Context) {
	resp := app.Response{Ctx: c}
	user_id := c.Param("id")
	cache.DeleteCache(c)
	resp.To(fmt.Sprintf("已清空用户(%s)缓存", user_id))
}
