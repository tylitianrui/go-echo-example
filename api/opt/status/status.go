package status

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type LiveStatus struct {
	TimeStamp int64 `json:"time_stamp"`
}

func Live(cxt echo.Context) error {
	return cxt.JSON(http.StatusOK, &LiveStatus{TimeStamp: time.Now().UnixNano()})
}
