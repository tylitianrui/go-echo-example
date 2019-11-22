package status

import (
	"github.com/labstack/echo"
	"go-echo-example/pkg/app"
	"net/http"
)

type LiveStatus struct {
	TimeStamp int64 `json:"time_stamp"`
}

func Live(ctx echo.Context) error {
	c := app.NewApp(ctx)
	return c.Response(http.StatusOK, app.CONNECT_STATUS_OK, nil)

}
