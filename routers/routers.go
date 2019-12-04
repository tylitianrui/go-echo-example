package routers

import (
	"github.com/labstack/echo"
	"go-echo-example/api/opt/status"
	"go-echo-example/api/v1/account"
)

func InitRouters(e *echo.Echo) {
	Opt(e)

	v1 := e.Group("/v1")
	ApiV1(v1)

}

// 服务器状态查询
func Opt(e *echo.Echo) {
	opt := e.Group("/opt")
	opt.GET("/status", status.Live)
}

func ApiV1(g *echo.Group) {
	account := g.Group("/account")
	account.POST("/verify", v1Account.VerifyCode)
	account.POST("/register", v1Account.Register)
	account.POST("/login", v1Account.Login)
}
