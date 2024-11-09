package errcode

var (
	Success         = NewError(0, "成功")
	ServerError     = NewError(10000000, "服务器内部错误")
	InvaildParams   = NewError(10000001, "入参错误")
	NotFound        = NewError(10000002, "找不到")
	NoAuth          = NewError(10000003, "鉴权失败")
	TooManyRequests = NewError(10000004, "请求频繁")
)
