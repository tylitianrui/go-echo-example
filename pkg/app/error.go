package app

const (
	ErrUnknown = 0 //未知错误 0   defalut
	SUCCESS    = 1 // 业务成功 1
	FAIL       = 2 // 业务失败 2

	ConnectStatusOk = 1000

	ParamsInvalid = 2000
	UserMobile    = 2001
)

var Message = map[int]string{
	SUCCESS:    "OK",
	FAIL:       "FAIL",
	ErrUnknown: "Unknown error",

	ConnectStatusOk: "server running",
	ParamsInvalid:   "params invalid",
	UserMobile:      "user mobile invalid",
}

// 获取错误信息
func GetMessage(errCode int) string {
	if msg, ok := Message[errCode]; ok {
		return msg
	}
	return Message[ErrUnknown]
}
