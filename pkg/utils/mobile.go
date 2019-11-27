package utils

import (
	"fmt"
	"go-echo-example/pkg/app"
)

func GetVerifyCodeKeyFromRds(mobile string) string {
	return fmt.Sprintf("%s%s", app.MobileVerifyCodeKeyPrefix, mobile)
}
