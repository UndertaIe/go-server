package errcode

var (
	Success                           = NewError(1000, "成功")
	OK                                = NewError(1001, "成功")
	ServerError                       = NewError(1002, "服务内部错误")
	InvalidParams                     = NewError(1003, "入参错误")
	NotFound                          = NewError(1004, "查询资源不存在")
	UnauthorizedAuthNotExist          = NewError(1005, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError            = NewError(1006, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout          = NewError(1007, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate         = NewError(1008, "鉴权失败，Token生成失败")
	UnauthorizedTokenSignatureInvalid = NewError(1009, "鉴权失败，Token签名异常")
	UnauthorizedUserError             = NewError(1010, "鉴权失败，当前用户无访问权限")
	TooManyRequests                   = NewError(1011, "请求过多，请稍后重试")
	RequestTimeout                    = NewError(1012, "请求超时，请稍后重试")
	UnKnownAuthType                   = NewError(1013, "未知的认证方式")
)
