package app

const (
	ERR_UNKNOWN = 0 //未知错误 0   defalut
	SUCCESS     = 1 // 业务成功 1
	FAIL        = 2 // 业务失败 2

	CONNECT_STATUS_OK = 1000

	ERR_PARAMS_ILLEGAL = 2000
)

var Message = map[int]string{
	SUCCESS:     "OK",
	FAIL:        "FAIL",
	ERR_UNKNOWN: "Unknown error",

	CONNECT_STATUS_OK:  "server running",
	ERR_PARAMS_ILLEGAL: "params error",
}

// 获取错误信息
func GetMessage(errCode int) string {
	if msg, ok := Message[errCode]; ok {
		return msg
	}
	return Message[ERR_UNKNOWN]
}
