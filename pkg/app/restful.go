package app

import "github.com/labstack/echo"

// 封装echo
type App struct {
	Ctx echo.Context
}

type Response struct {
	Code int         `json:"code"` // 业务状态码
	Msg  string      `json:"msg"`  // 业务状态信息
	Data interface{} `json:"data"` // 数据
}

func (self *App) Response(httpCode, errCode int, data interface{}) error {

	return self.Ctx.JSON(httpCode, &Response{
		Code: errCode,
		Msg:  "",
		Data: data,
	})
}
