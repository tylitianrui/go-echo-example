package v1Account

import (
	"crypto/md5"
	"fmt"
	"github.com/labstack/echo"
	"go-echo-example/helper"
	"go-echo-example/models"
	"go-echo-example/pkg/app"
	"go-echo-example/pkg/utils"
	"go-echo-example/service"
	"hash"
	"net/http"
)

type RegisterAccountData struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Mobile   int    `json:"mobile"`
	Code     string `json:"code"`
}

func Register(ctx echo.Context) error {
	var (
		err  error
		code []byte
	)
	c := app.NewApp(ctx)

	registerAccountData := new(RegisterAccountData)
	if err = ctx.Bind(registerAccountData); err != nil {
		return c.Response(http.StatusOK, app.ErrParamsInvalid, nil)
	}

	mobile_verify_key := utils.GetVerifyCodeKeyFromRds(registerAccountData.Mobile)

	if code, err = helper.RedisGet(mobile_verify_key); err != nil {
		return c.Response(http.StatusOK, app.ErrMobileVerifyCideNotFound, nil)
	}

	if string(code) == registerAccountData.Code {
		helper.RedisDel(mobile_verify_key)
	}
	accountSrv := service.NewAccountService()

	m := app.Crypt{
		Alg: func() hash.Hash {
			return md5.New()
		},
	}

	account := &models.Account{
		User:     registerAccountData.User,
		Password: m.Proc(registerAccountData.Password),
		Mobile:   registerAccountData.Mobile,
		Status:   app.RegisterActivated,
		Created:  utils.TimeStamp(),
	}

	if err = accountSrv.Create(account); err != nil {
		fmt.Println(err)
		return c.Response(http.StatusOK, app.ErrRegisterFail, nil)
	}

	return c.SuccessResponse("created")
}
