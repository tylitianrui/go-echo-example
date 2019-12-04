package v1Account

import (
	"crypto/md5"
	"fmt"
	"github.com/labstack/echo"
	"go-echo-example/models"
	"go-echo-example/pkg/app"
	"go-echo-example/pkg/verify"
	"go-echo-example/service"
	"hash"
	"net/http"
)

type LoginReqData struct {
	Password string `json:"password"`
	Mobile   int    `json:"mobile"`
}

func Login(ctx echo.Context) error {
	var (
		//ctx context.Context
		err     error
		account *models.Account
	)
	req := new(LoginReqData)
	c := app.NewApp(ctx)

	if err = c.Ctx.Bind(req); err != nil {
		return c.Response(http.StatusOK, app.ErrParamsInvalid, nil)
	}
	// 校验手机格式
	if !verify.VerfifyMobile(req.Mobile) {
		return c.Response(http.StatusOK, app.ErrUserMobile, nil)
	}

	accountSrv := service.NewAccountService()

	if account, err = accountSrv.GetAccountByMobile(req.Mobile); err != nil {
		fmt.Println(err)
		return c.Response(http.StatusOK, app.ErrAccountOrPwdWrong, nil)
	}
	m := app.Crypt{
		Alg: func() hash.Hash {
			return md5.New()
		},
	}
	if !m.Check(req.Password, account.Password) {
		return c.Response(http.StatusOK, app.ErrAccountOrPwdWrong, nil)
	}

	token, _ := app.GenToken(account.ID)
	//if token,_:= app.GenToken(account.ID);err != nil {
	//    fmt.Println(err)
	//    return
	//}

	resp := struct {
		Token string
	}{token}

	return c.Response(http.StatusOK, app.SUCCESS, resp)

}
