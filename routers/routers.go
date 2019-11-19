package routers

import (
	"github.com/labstack/echo"
	"go-echo-example/api/opt/status"
)

func InitRouters(e *echo.Echo) {
	Opt(e)

}

// 服务器状态查询
func Opt(e *echo.Echo) {
	opt := e.Group("/opt")
	opt.GET("/status", status.Live)
}
