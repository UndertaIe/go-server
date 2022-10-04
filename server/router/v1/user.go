package v1

import (
	"github.com/UndertaIe/go-server-env/internal/service"
	"github.com/UndertaIe/go-server-env/pkg/app"
	"github.com/UndertaIe/go-server-env/pkg/errcode"
	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

// Get godoc
// @Summary     获取单个用户
// @Description  通过id获取单个用户
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "用户 ID"
// @Success      200  {object}  service.User  "成功"
// @Failure      400  {object}  errcode.Error "请求错误"
// @Failure      500  {object}  errcode.Error "内部错误"
// @Router       /api/v1/user/{id} [get]
func (u User) Get(c *gin.Context) {
	resp := app.NewResponse(c)
	user_id, binderr := conv.Int(c.Param("id"))
	if binderr != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	param := service.UserGetParam{UserId: user_id}
	srv := service.NewService(c.Request.Context())

	user, err := srv.GetUser(&param)

	if err != nil {
		resp.ToError(err)
		return
	}
	resp.To(user)
}

// List godoc
// @Summary     获取用户分页
// @Description  获取用户分页
// @Tags         User
// @Produce      json
// @Success      200  {object}  []service.User  "成功"
// @Failure      400  {object}  errcode.Error "请求错误"
// @Failure      500  {object}  errcode.Error "内部错误"
// @Router       /api/v1/user [get]
func (u User) List(c *gin.Context) {
	resp := app.NewResponse(c)
	pager := app.NewPager(c)
	srv := service.NewService(c.Request.Context())

	user, err := srv.GetUserList(pager)
	if err != nil {
		resp.ToError(err)
		return
	}
	resp.ToList(user, pager)
}

// Create godoc
// @Summary     创建用户
// @Description  通过一些字段创建用户
// @Tags         User
// @Accept       json
// @Param user_name 	body string true  "用户名"
// @Param password 		body string true  "用户密码"
// @Param phone_number  body string true  "手机号码"
// @Param email 		body string false "电子邮件"
// @Param sex 			body int 	false "性别"
// @Param description   body string false "用户简介"
// @Success      200  {string}  string        "成功"
// @Failure      400  {object}  errcode.Error "请求错误"
// @Failure      500  {object}  errcode.Error "内部错误"
// @Router       /api/v1/user [post]
func (u User) Create(c *gin.Context) {
	resp := app.NewResponse(c)
	param := new(service.UserCreateParam)
	if e := c.ShouldBind(param); e != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}

	srv := service.NewService(c.Request.Context())
	err := srv.CreateUser(param)
	if err != nil {
		resp.ToError(err)
		return
	}
	resp.Ok()
}

// Update godoc
// @Summary     更新单个用户
// @Description  通过id和一些model字段更新单个用户
// @Tags         User
// @Accept       json
// @Param        id   path      int  true  "用户 ID"
// @Param user_name 	body string true  "用户名"
// @Param password 		body string true  "用户密码"
// @Param phone_number  body string true  "手机号码"
// @Param email 		body string false "电子邮件"
// @Param sex 			body int 	false "性别"
// @Param description   body string false "用户简介"
// @Success      200  {string}  string   "成功"
// @Failure      400  {object}  errcode.Error "请求错误"
// @Failure      500  {object}  errcode.Error "内部错误"
// @Router       /api/v1/user/{id} [put]
func (u User) Update(c *gin.Context) {
	resp := app.NewResponse(c)
	param := new(service.UserUpdateParam)
	var binderr error
	param.UserId, binderr = conv.Int(c.Param("id"))
	if binderr != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	binderr2 := c.ShouldBind(param)
	if binderr2 != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	srv := service.NewService(c.Request.Context())
	err := srv.UpdateUser(param)

	if err != nil {
		resp.ToError(err)
		return
	}
	resp.Ok()
}

// Delete godoc
// @Summary     删除单个用户
// @Description  通过id删除单个用户
// @Tags         User
// @Param        id   path      int  true  "用户 ID"
// @Success      200  {string}  string  "成功"
// @Failure      400  {object}  errcode.Error "请求错误"
// @Failure      500  {object}  errcode.Error "内部错误"
// @Router       /api/v1/user/{id} [delete]
func (u User) Delete(c *gin.Context) {
	user_id, binderr := conv.Int(c.Param("id"))
	resp := app.NewResponse(c)
	if binderr != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	param := service.UserDeleteParam{UserId: user_id}
	srv := service.NewService(c.Request.Context())

	err := srv.DeleteUser(&param)
	if err != nil {
		resp.ToError(err)
		return
	}
	resp.Ok()
}

// PhoneExists godoc
// @Summary     判断手机号是否已经被注册
// @Tags         UserSignUp
// @Accept       json
// @Produce      json
// @Param        phone_number   body	string		true	"手机号"
// @Success      200  {string}  string 			"成功"
// @Failure      400  {object}  errcode.Error 	"请求错误"
// @Failure      500  {object}  errcode.Error 	"内部错误"
// @Router       /api/v1/user/exists/phone [post]
func (User) PhoneExists(c *gin.Context) {
	srv := service.NewService(c)
	param := service.UserPhoneExistsParam{}
	binderr := c.Bind(&param)
	resp := app.NewResponse(c)
	if binderr != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	has, err := srv.IsExistsUserPhone(param)
	if err != nil {
		resp.ToError(err)
		return
	}
	if has { // 手机号已存在
		resp.ToError(errcode.UserPhoneExists)
		return
	}
	resp.Ok()
}

// EmailExists godoc
// @Summary     判断邮箱是否已经被注册
// @Tags         UserSignUp
// @Accept       json
// @Produce      json
// @Param        email   body	string		true	"用户email"
// @Success      200  {string}  string 			"成功"
// @Failure      400  {object}  errcode.Error 	"请求错误"
// @Failure      500  {object}  errcode.Error 	"内部错误"
// @Router       /api/v1//user/email [post]
func (User) EmailExists(c *gin.Context) {
	srv := service.NewService(c)
	param := service.UserEmailExistsParam{}
	binderr := c.Bind(&param)
	resp := app.NewResponse(c)
	if binderr != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	has, err := srv.IsExistsUserEmail(param)
	if err != nil {
		resp.ToError(errcode.ErrorService)
		return
	}
	if has { // 邮箱已存在
		resp.ToError(errcode.UserEmailExists)
		return
	}
	resp.Ok()
}

// UserNameExists godoc
// @Summary     判断用户名是否已经被注册
// @Tags         UserSignUp
// @Accept       json
// @Produce      json
// @Param        user_name   body	string		true	"用户名"
// @Success      200  {string}  string 			"成功"
// @Failure      400  {object}  errcode.Error 	"请求错误"
// @Failure      500  {object}  errcode.Error 	"内部错误"
// @Router       /api/v1//user/name [post]
func (User) UserNameExists(c *gin.Context) {
	srv := service.NewService(c)
	param := service.UserNameExistsParam{}
	binderr := c.Bind(&param)
	resp := app.NewResponse(c)
	if binderr != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	has, err := srv.IsExistsUserName(param)
	if err != nil {
		resp.ToError(err)
		return
	}
	if has { // 用户名已存在
		resp.ToError(errcode.UserNameExists)
		return
	}
	resp.Ok()
}
