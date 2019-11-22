package app

const (
	ERR_UNKNOWN = 0 //未知错误 0   defalut
	SUCCESS     = 1 // 业务成功 1
	FAIL        = 2 // 业务失败 2
)

var Message = map[int]string{
	SUCCESS:     "OK",
	FAIL:        "FAIL",
	ERR_UNKNOWN: "Unknown error",
}

// 获取错误信息
func GetMessage(errCode int) string {
	if msg, ok := Message[errCode]; ok {
		return msg
	}
	return Message[ERR_UNKNOWN]
}
