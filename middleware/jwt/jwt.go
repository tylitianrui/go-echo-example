package jwt

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go-echo-example/pkg/setting"
)

func LoginRequired() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		ContextKey:  "id",
		SigningKey:  []byte(setting.G_AppConf.JwtSecret),
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  "Bearer",
	}
	return middleware.JWTWithConfig(config)
}
