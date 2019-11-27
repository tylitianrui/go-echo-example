package app

const (
	ErrUnknown = 0 //未知错误 0   defalut
	SUCCESS    = 1 // 业务成功 1
	FAIL       = 2 // 业务失败 2

	ErrServer = 500 // 服务器错误

	ConnectStatusOk = 1000

	ErrParamsInvalid = 2000
	ErrUserMobile    = 2001

	ErrMobileVerifyCideNotFound = 2002 // 未找到验证码
	ErrRegisterFail             = 2003 // 账号注册失败
)

var Message = map[int]string{
	SUCCESS:                     "OK",
	FAIL:                        "FAIL",
	ErrUnknown:                  "Unknown error",
	ErrServer:                   "server internal  error",
	ConnectStatusOk:             "server running",
	ErrParamsInvalid:            "params invalid",
	ErrUserMobile:               "user mobile invalid",
	ErrMobileVerifyCideNotFound: "regain mobile verify  code",
	ErrRegisterFail:             "Register Fail",
}

// 获取错误信息
func GetMessage(errCode int) string {
	if msg, ok := Message[errCode]; ok {
		return msg
	}
	return Message[ErrUnknown]
}
