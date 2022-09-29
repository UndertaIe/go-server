package v1

import (
	"strconv"

	"github.com/UndertaIe/passwd/internal/service"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/UndertaIe/passwd/pkg/page"
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
	user_id, err := strconv.Atoi(c.Param("id"))
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	param := service.UserGetRequest{UserId: user_id}
	srv := service.NewService(c.Request.Context())

	user, err := srv.GetUser(&param)

	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.To(user)
}

// List godoc
// @Summary     获取多个用户
// @Description  获取用户分页
// @Tags         User
// @Produce      json
// @Success      200  {object}  []service.User  "成功"
// @Failure      400  {object}  errcode.Error "请求错误"
// @Failure      500  {object}  errcode.Error "内部错误"
// @Router       /api/v1/user [get]
func (u User) List(c *gin.Context) {
	srv := service.NewService(c.Request.Context())

	param := service.UserGetRequest{}
	pager := page.NewPager(c)
	user, err := srv.GetUserList(&param, pager)

	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
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
	param := new(service.UserCreateRequest)
	err := c.ShouldBind(param)
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	srv := service.NewService(c.Request.Context())

	err = srv.CreateUser(param)

	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
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
	params := new(service.UserUpdateRequest)
	user_id, err0 := strconv.Atoi(c.Param("id"))
	params.UserId = user_id
	err1 := c.ShouldBind(params)
	resp := app.Response{Ctx: c}
	if err0 != nil || err1 != nil {
		newErr := errcode.InvalidParams.WithDetails(err0.Error()).WithDetails(err1.Error())
		resp.ToError(newErr)
		return
	}
	srv := service.NewService(c.Request.Context())
	err := srv.UpdateUser(params)

	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
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
	user_id, err := strconv.Atoi(c.Param("id"))
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	param := service.UserDeleteRequest{UserId: user_id}
	srv := service.NewService(c.Request.Context())

	err = srv.DeleteUser(&param)
	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
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
// @Router       /api/v1/user/phone [post]
func (User) PhoneExists(c *gin.Context) {
	srv := service.NewService(c)
	param := &service.UserPhoneExistsParam{}
	err := c.Bind(param)
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	has, err := srv.IsExistsUserPhone(param)
	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	if has {
		resp.ToError(errcode.UserPhoneExists) // 手机号已存在
		return
	}
	resp.Ok()
}

// EmailExists godoc
// @Summary     判断email是否已经被注册
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
	param := &service.UserEmailExistsParam{}
	err := c.Bind(param)
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	has, err := srv.IsExistsUserEmail(param)
	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	if has {
		resp.ToError(errcode.UserEmailExists) // 邮箱已存在
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
	param := &service.UserNameExistsParam{}
	err := c.Bind(param)
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	has, err := srv.IsExistsUserName(param)
	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	if has {
		resp.ToError(errcode.UserNameExists) // 用户名已存在
		return
	}
	resp.Ok()
}
