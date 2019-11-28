package v1Account

import (
	"github.com/labstack/echo"
	"go-echo-example/helper"
	"go-echo-example/pkg/app"
	"go-echo-example/pkg/setting"
	"go-echo-example/pkg/utils"
	"go-echo-example/pkg/verify"
	"net/http"
)

type VerifyCodeData struct {
	Mobile int `json:"mobile"`
}

// 获取手机验证码
func VerifyCode(ctx echo.Context) error {
	var (
		err error
	)
	c := app.NewApp(ctx)

	reqData := new(VerifyCodeData)

	if err = ctx.Bind(reqData); err != nil {

		return c.Response(http.StatusOK, app.ErrParamsInvalid, nil)
	}

	if !verify.VerfifyMobile(reqData.Mobile) {
		return c.Response(http.StatusOK, app.ErrUserMobile, nil)
	}
	verifyCode := verify.GenVerifyCodeString(setting.G_AppConf.MobileVerifyNum)
	mobile_verify_key := utils.GetVerifyCodeKeyFromRds(reqData.Mobile)
	helper.RedisSet(mobile_verify_key, verifyCode, setting.MobileVerfycCodeExpire)

	return c.SuccessResponse(verifyCode)

}
