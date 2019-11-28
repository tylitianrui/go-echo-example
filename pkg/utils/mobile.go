package utils

import (
	"fmt"
	"go-echo-example/pkg/app"
)

func GetVerifyCodeKeyFromRds(mobile int) string {
	return fmt.Sprintf("%s%d", app.MobileVerifyCodeKeyPrefix, mobile)
}

//
//func ModifyMobile(mobile interface{}) int {
//	switch mobile.(type) {
//	case string:
//
//	}
//
//}
