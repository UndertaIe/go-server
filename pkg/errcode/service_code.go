package errcode

var (
	ErrorService                = NewError(20000, "服务异常，请稍后重试")
	ErrorUserRecordNotFound     = NewError(20001, "查询用户不存在")
	ErrorPlatformRecordNotFound = NewError(20002, "查询平台不存在")
)
