package app

import (
	"github.com/labstack/echo"
	"net/http"
)

// 封装echo
type App struct {
	Ctx echo.Context
}

func NewApp(ctx echo.Context) *App {
	return &App{
		Ctx: ctx,
	}
}

type Response struct {
	Code int         `json:"code"` // 业务状态码
	Msg  string      `json:"msg"`  // 业务状态信息
	Data interface{} `json:"data"` // 数据
}

func (self *App) Response(httpCode, errCode int, data interface{}) error {

	return self.Ctx.JSON(httpCode, &Response{
		Code: errCode,
		Msg:  GetMessage(errCode),
		Data: data,
	})
}

func (self *App) SuccessResponse(data interface{}) error {
	return self.Response(http.StatusOK, SUCCESS, data)

}
