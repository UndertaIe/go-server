package demo

import (
	"fmt"
	"time"

	"github.com/UndertaIe/go-server-env/internal/service"
	"github.com/UndertaIe/go-server-env/pkg/app"
	"github.com/UndertaIe/go-server-env/pkg/cache"
	"github.com/UndertaIe/go-server-env/pkg/errcode"
	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
)

func Now(c *gin.Context) {
	resp := app.NewResponse(c)
	now := time.Now().UTC().GoString()
	resp.To(now)
}

func CacheNow(c *gin.Context) {
	resp := app.NewResponse(c)
	now := time.Now().UTC().GoString()
	resp.To(now)
}

func GetUser(c *gin.Context) {
	user_id, err := conv.Int(c.Param("id"))
	resp := app.NewResponse(c)
	if err != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	param := service.UserGetParam{UserId: user_id}
	srv := service.NewService(c.Request.Context())
	log.Info("缓存未命中，访问数据库...")
	user, err2 := srv.GetUser(&param)

	if err2 != nil {
		newErr := errcode.ErrorService.WithDetails(err2.Error())
		resp.ToError(newErr)
		return
	}
	resp.To(user)
}

func DeleteUser(c *gin.Context) {
	resp := app.NewResponse(c)
	user_id := c.Param("id")
	cache.DeleteCache(c)
	resp.To(fmt.Sprintf("已清空用户(%s)缓存", user_id))
}

func UpdateUser(c *gin.Context) {
	resp := app.NewResponse(c)
	user_id := c.Param("id")
	cache.DeleteCache(c)
	resp.To(fmt.Sprintf("已清空用户(%s)缓存", user_id))
}
